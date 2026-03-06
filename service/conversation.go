package service

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"edu/conf"
	"edu/lib/ai"
	"edu/lib/logger"
	"edu/model"
	"edu/repository"
	"fmt"
	"time"

	"go.uber.org/zap"
)

const (
	// conversationSessionTTL is the default time-to-live for a conversation session.
	conversationSessionTTL = 24 * time.Hour

	// maxHistoryMessages is the maximum number of messages kept in full before
	// summarisation is triggered.
	maxHistoryMessages = 20

	// summaryTriggerCount is the message count threshold at which a rolling
	// summary is generated to replace the oldest messages.
	summaryTriggerCount = 20

	// maxMessagesAfterSummary is the number of recent messages retained alongside
	// the summary after trimming.
	maxMessagesAfterSummary = 10

	// systemPromptTemplate is the base system prompt injected at the beginning of
	// every conversation to give the AI context about its role.
	systemPromptTemplate = `You are an intelligent educational assistant for an online learning platform.
You are helping a %s named "%s".
Current date/time: %s

User context:
%s

Guidelines:
- Maintain context from previous messages in this conversation.
- If the user asks a follow-up question, use the conversation history to understand what they are referring to.
- Tailor your responses to the user's role (%s): %s
- Be concise but thorough.
- If the user says "reset" or "start over", acknowledge that the context has been reset.`
)

// conversationSystemPrompt is the base educational domain system prompt for the AI tutor.
const conversationSystemPrompt = `You are an intelligent educational assistant for LTEdu, a learning platform focused on academic subjects including Biology, Physics, Chemistry, Mathematics, and more.

Your capabilities:
- Answer subject-specific questions (Biology, Physics, Chemistry, Mathematics, etc.)
- Explain concepts clearly in both Chinese and English
- Help students understand exam question patterns and mark schemes
- Provide study guidance and learning strategies
- Analyse student performance and suggest improvements
- Support Cambridge IGCSE, A-Level, and other qualification formats

Guidelines:
- Always be encouraging and supportive
- Use clear, age-appropriate language
- When explaining scientific concepts, use examples
- For exam questions, follow the relevant mark scheme conventions
- Respond in the same language the student uses (Chinese or English)
- For multi-part questions, address each part systematically`

// subjectSystemPromptSuffix adds subject-specific context to the system prompt.
const subjectSystemPromptSuffix = "\n\nCurrent subject context: %s. Focus your responses on this subject."

// ConversationSvr is the singleton conversation service.
var ConversationSvr = &ConversationService{
	baseService: newBaseService(),
	ai:          ai.NewModel(conf.Conf.AiConfig, logger.Logger),
}

// ConversationService manages multi-turn AI conversation sessions.
type ConversationService struct {
	baseService
	ai ai.Model
}

// generateSessionKey creates a cryptographically random 32-byte hex session key.
func generateSessionKey() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// userRoleGuidance returns role-specific guidance text for the system prompt.
func userRoleGuidance(role string) string {
	switch role {
	case "teacher":
		return "focus on pedagogy, curriculum design, exam preparation strategies, and assessment techniques"
	case "admin":
		return "focus on platform management, analytics, user management, and system operations"
	default: // "student"
		return "focus on learning, practice, exam revision, and understanding concepts clearly"
	}
}

// buildSystemPrompt constructs a personalised system prompt for a given session.
func buildSystemPrompt(session *model.ConversationSession, userName string) string {
	ctx := &model.ConversationContext{}
	if session.ContextData != "" {
		_ = json.Unmarshal([]byte(session.ContextData), ctx)
	}

	contextSummary := "No specific context provided."
	if len(ctx.Preferences) > 0 || len(ctx.RecentActions) > 0 || len(ctx.CurrentSelection) > 0 {
		parts, _ := json.Marshal(ctx)
		contextSummary = string(parts)
	}

	role := session.UserRole
	if role == "" {
		role = "student"
	}

	return fmt.Sprintf(systemPromptTemplate,
		role, userName,
		time.Now().Format("2006-01-02 15:04"),
		contextSummary,
		role, userRoleGuidance(role),
	)
}

// StartSession creates a new conversation session for the given user.
func (s *ConversationService) StartSession(userID uint, userName string, userRole string, ctx *model.ConversationContext) (*model.ConversationSession, error) {
	key, err := generateSessionKey()
	if err != nil {
		return nil, fmt.Errorf("failed to generate session key: %w", err)
	}

	var contextJSON string
	if ctx != nil {
		b, _ := json.Marshal(ctx)
		contextJSON = string(b)
	}

	if userRole == "" {
		userRole = "student"
	}

	session := &model.ConversationSession{
		UserID:       userID,
		SessionKey:   key,
		UserRole:     userRole,
		ContextData:  contextJSON,
		LastActiveAt: time.Now(),
		ExpiresAt:    time.Now().Add(conversationSessionTTL),
		IsActive:     true,
	}

	if err := repository.ConversationRepo.CreateSession(session); err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}

	return session, nil
}

// SendMessage sends a user message within a session and returns the AI reply.
// It maintains full conversation history and generates a rolling summary when
// the history grows beyond maxHistoryMessages.
func (s *ConversationService) SendMessage(userID uint, userName string, sessionKey string, userMessage string) (*model.SendMessageResponse, error) {
	// Load session and enforce ownership / expiry
	session, messages, err := repository.ConversationRepo.GetActiveSessionWithMessages(sessionKey, userID)
	if err != nil {
		return nil, fmt.Errorf("session not found or expired")
	}

	// Build the message list for the AI: system prompt → optional summary → history → new message
	aiMessages := []ai.Message{
		{Role: "system", Content: buildSystemPrompt(session, userName)},
	}

	if session.Summary != "" {
		aiMessages = append(aiMessages, ai.Message{
			Role:    "system",
			Content: "Conversation summary so far: " + session.Summary,
		})
	}

	for _, m := range messages {
		aiMessages = append(aiMessages, ai.Message{
			Role:    m.Role,
			Content: m.Content,
		})
	}
	aiMessages = append(aiMessages, ai.Message{Role: "user", Content: userMessage})

	// Call the AI
	reply, err := s.ai.CreateCompletionWithHistory(aiMessages)
	if err != nil {
		logger.Logger.Error("ConversationService.SendMessage AI error", zap.Error(err))
		return nil, fmt.Errorf("AI service error: %w", err)
	}

	// Persist user message
	userMsg := &model.ConversationMessage{
		SessionID:  session.ID,
		Role:       "user",
		Content:    userMessage,
		OrderIndex: session.MessageCount,
	}
	if err := repository.ConversationRepo.CreateMessage(userMsg); err != nil {
		return nil, fmt.Errorf("failed to save user message: %w", err)
	}

	// Persist assistant message
	assistantMsg := &model.ConversationMessage{
		SessionID:  session.ID,
		Role:       "assistant",
		Content:    reply,
		OrderIndex: session.MessageCount + 1,
	}
	if err := repository.ConversationRepo.CreateMessage(assistantMsg); err != nil {
		return nil, fmt.Errorf("failed to save assistant message: %w", err)
	}

	// Update session metadata
	session.MessageCount += 2
	session.LastActiveAt = time.Now()

	// Trigger summarisation when history is too long
	if session.MessageCount >= summaryTriggerCount {
		if err := s.summariseAndTrim(session); err != nil {
			// Non-fatal: log and continue
			logger.Logger.Warn("ConversationService: summarisation failed", zap.Error(err))
		}
	}

	if err := repository.ConversationRepo.UpdateSession(session); err != nil {
		return nil, fmt.Errorf("failed to update session: %w", err)
	}

	return &model.SendMessageResponse{
		UserMessage:      *userMsg,
		AssistantMessage: *assistantMsg,
		SessionKey:       sessionKey,
		MessageCount:     session.MessageCount,
	}, nil
}

// summariseAndTrim generates an AI summary of the current conversation, stores it
// on the session, and deletes all but the most recent maxMessagesAfterSummary
// messages to keep the context window manageable.
func (s *ConversationService) summariseAndTrim(session *model.ConversationSession) error {
	messages, err := repository.ConversationRepo.GetMessagesBySessionID(session.ID)
	if err != nil || len(messages) <= maxMessagesAfterSummary {
		return err
	}

	// Build a text transcript to summarise
	transcript := ""
	if session.Summary != "" {
		transcript = "Previous summary: " + session.Summary + "\n\n"
	}
	for _, m := range messages {
		transcript += fmt.Sprintf("%s: %s\n", m.Role, m.Content)
	}

	summaryPrompt := fmt.Sprintf(
		"Please summarise the following conversation in 2-3 sentences, preserving key context, decisions, and topics discussed:\n\n%s",
		transcript,
	)

	summary, err := s.ai.CreateCompletion(summaryPrompt)
	if err != nil {
		return err
	}

	// Delete all messages for this session, then re-create only the tail.
	if err := repository.ConversationRepo.DeleteMessagesBySessionID(session.ID); err != nil {
		return err
	}
	startIdx := len(messages) - maxMessagesAfterSummary
	if startIdx < 0 {
		startIdx = 0
	}
	tail := messages[startIdx:]
	for _, m := range tail {
		msg := &model.ConversationMessage{
			SessionID:  m.SessionID,
			Role:       m.Role,
			Content:    m.Content,
			OrderIndex: m.OrderIndex,
		}
		_ = repository.ConversationRepo.CreateMessage(msg)
	}

	session.Summary = summary
	session.MessageCount = maxMessagesAfterSummary
	return nil
}

// GetHistory returns the full message history for a session.
func (s *ConversationService) GetHistory(userID uint, sessionKey string) ([]model.ConversationMessage, error) {
	_, messages, err := repository.ConversationRepo.GetActiveSessionWithMessages(sessionKey, userID)
	return messages, err
}

// ListSessions returns all active sessions for a user.
func (s *ConversationService) ListSessions(userID uint) ([]model.ConversationSession, error) {
	return repository.ConversationRepo.GetSessionsByUserID(userID)
}

// ResetContext clears the message history and optional context data for a session,
// allowing the user to start fresh without creating a new session key.
func (s *ConversationService) ResetContext(userID uint, sessionKey string, newCtx *model.ConversationContext) error {
	session, err := repository.ConversationRepo.GetSessionByKey(sessionKey)
	if err != nil {
		return fmt.Errorf("session not found")
	}
	if session.UserID != userID {
		return fmt.Errorf("unauthorised")
	}

	if err := repository.ConversationRepo.DeleteMessagesBySessionID(session.ID); err != nil {
		return fmt.Errorf("failed to clear history: %w", err)
	}

	session.MessageCount = 0
	session.Summary = ""
	session.LastActiveAt = time.Now()

	if newCtx != nil {
		b, _ := json.Marshal(newCtx)
		session.ContextData = string(b)
	}

	return repository.ConversationRepo.UpdateSession(session)
}

// CloseSession deactivates a conversation session.
func (s *ConversationService) CloseSession(userID uint, sessionKey string) error {
	return repository.ConversationRepo.DeactivateSession(sessionKey, userID)
}

// CleanupExpiredSessions removes sessions past their expiry time.
// This can be called from a background task.
func (s *ConversationService) CleanupExpiredSessions() (int64, error) {
	return repository.ConversationRepo.DeleteExpiredSessions()
}

