package service

import (
"edu/conf"
"edu/lib/ai"
"edu/lib/logger"
"edu/model"
"edu/repository"
"fmt"
"time"
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

// ConversationService manages AI conversation sessions.
type ConversationService struct {
baseService
ai ai.Model
}

// StartSession creates a new conversation session for a user.
func (s *ConversationService) StartSession(userID uint, req model.ConversationStartRequest) (*model.ConversationSession, error) {
title := req.Title
if title == "" {
title = "New Conversation"
}

session := &model.ConversationSession{
UserID:    userID,
Title:     title,
Subject:   req.Subject,
Active:    true,
ExpiresAt: time.Now().Add(24 * time.Hour),
}

if err := repository.ConversationSessionRepo.Create(session); err != nil {
return nil, fmt.Errorf("failed to create session: %w", err)
}
return session, nil
}

// SendMessage appends the user message to a session, calls the AI, stores and returns the reply.
func (s *ConversationService) SendMessage(userID uint, req model.ConversationMessageRequest) (*model.ConversationMessage, error) {
session, err := repository.ConversationSessionRepo.GetByID(req.SessionID)
if err != nil {
return nil, fmt.Errorf("session not found: %w", err)
}
if session.UserID != userID {
return nil, fmt.Errorf("access denied")
}
if !session.Active {
return nil, fmt.Errorf("session is closed")
}

// Persist the user message
userMsg := &model.ConversationMessage{
SessionID: session.ID,
Role:      model.ConvRoleUser,
Content:   req.Message,
}
if err := repository.ConversationMessageRepo.Create(userMsg); err != nil {
return nil, fmt.Errorf("failed to store message: %w", err)
}

// Build message history for the AI call
messages, err := s.buildMessageHistory(session)
if err != nil {
return nil, err
}

// Call the AI
replyContent, err := s.ai.CreateCompletionWithMessages(messages)
if err != nil {
return nil, fmt.Errorf("AI response failed: %w", err)
}

// Persist the assistant reply
assistantMsg := &model.ConversationMessage{
SessionID: session.ID,
Role:      model.ConvRoleAssistant,
Content:   replyContent,
}
if err := repository.ConversationMessageRepo.Create(assistantMsg); err != nil {
logger.Logger.Error("failed to store assistant message")
}

return assistantMsg, nil
}

// GetHistory returns all messages for a session.
func (s *ConversationService) GetHistory(userID, sessionID uint) ([]model.ConversationMessage, error) {
session, err := repository.ConversationSessionRepo.GetByID(sessionID)
if err != nil {
return nil, fmt.Errorf("session not found: %w", err)
}
if session.UserID != userID {
return nil, fmt.Errorf("access denied")
}
return repository.ConversationMessageRepo.GetBySessionID(sessionID)
}

// GetSessions returns all active sessions for a user.
func (s *ConversationService) GetSessions(userID uint) ([]model.ConversationSession, error) {
return repository.ConversationSessionRepo.GetByUserID(userID)
}

// ResetSession clears all messages in a session.
func (s *ConversationService) ResetSession(userID, sessionID uint) error {
session, err := repository.ConversationSessionRepo.GetByID(sessionID)
if err != nil {
return fmt.Errorf("session not found: %w", err)
}
if session.UserID != userID {
return fmt.Errorf("access denied")
}
return repository.ConversationMessageRepo.DeleteBySessionID(sessionID)
}

// CloseSession marks a session as inactive.
func (s *ConversationService) CloseSession(userID, sessionID uint) error {
session, err := repository.ConversationSessionRepo.GetByID(sessionID)
if err != nil {
return fmt.Errorf("session not found: %w", err)
}
if session.UserID != userID {
return fmt.Errorf("access denied")
}
session.Active = false
return repository.ConversationSessionRepo.Update(session)
}

// CleanupExpiredSessions removes sessions past their expiry time.
// This can be called from a background task.
func (s *ConversationService) CleanupExpiredSessions() (int64, error) {
	return repository.ConversationRepo.DeleteExpiredSessions()
}

func (s *ConversationService) buildMessageHistory(session *model.ConversationSession) ([]ai.Message, error) {
sysPrompt := conversationSystemPrompt
if session.Subject != "" {
sysPrompt += fmt.Sprintf(subjectSystemPromptSuffix, session.Subject)
}
messages := []ai.Message{
{Role: model.ConvRoleSystem, Content: sysPrompt},
}

history, err := repository.ConversationMessageRepo.GetBySessionID(session.ID)
if err != nil {
return nil, fmt.Errorf("failed to load history: %w", err)
}
for _, m := range history {
messages = append(messages, ai.Message{Role: m.Role, Content: m.Content})
}
return messages, nil
}
