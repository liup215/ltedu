package service

import (
	"edu/conf"
	"edu/lib/ai"
	"edu/lib/logger"
	"edu/model"
	"edu/repository"
	"errors"
)

const conversationSystemPrompt = `You are an intelligent educational assistant for the LTEdu platform. 
You help students and teachers with:
- Explaining syllabus topics and knowledge points
- Answering questions about study materials
- Providing learning advice and study strategies
- Helping with exam preparation

Be helpful, accurate, and encouraging. Keep responses concise and educational.`

var ConversationSvr = &ConversationService{
	baseService: newBaseService(),
	aiModel:     ai.NewModel(conf.Conf.AiConfig, logger.Logger),
}

type ConversationService struct {
	baseService
	aiModel ai.Model
}

// StartSession creates a new conversation session for a user
func (s *ConversationService) StartSession(userId uint, title string) (*model.ConversationSession, error) {
	if title == "" {
		title = "New Conversation"
	}
	session := &model.ConversationSession{
		UserId:   userId,
		Title:    title,
		IsActive: true,
	}
	if err := repository.ConversationSessionRepo.Create(session); err != nil {
		return nil, err
	}
	return session, nil
}

// SendMessage sends a user message and returns the AI reply
func (s *ConversationService) SendMessage(userId uint, req model.ConversationMessageRequest) (*model.ConversationMessage, error) {
	// Verify session ownership
	session, err := repository.ConversationSessionRepo.GetByID(req.SessionId)
	if err != nil {
		return nil, errors.New("session not found")
	}
	if session.UserId != userId {
		return nil, errors.New("unauthorized to access this session")
	}
	if !session.IsActive {
		return nil, errors.New("session is closed")
	}

	// Save user message
	userMsg := &model.ConversationMessage{
		SessionId: req.SessionId,
		Role:      "user",
		Content:   req.Content,
	}
	if err := repository.ConversationMessageRepo.Create(userMsg); err != nil {
		return nil, err
	}

	// Build conversation history for AI
	history, err := repository.ConversationMessageRepo.AllBySession(req.SessionId)
	if err != nil {
		return nil, err
	}

	messages := make([]ai.ChatMessage, 0, len(history)+1)
	messages = append(messages, ai.ChatMessage{Role: "system", Content: conversationSystemPrompt})
	for _, m := range history {
		messages = append(messages, ai.ChatMessage{Role: m.Role, Content: m.Content})
	}

	// Get AI response
	aiResponse, err := s.aiModel.CreateChatCompletion(messages)
	if err != nil {
		return nil, err
	}

	// Save assistant message
	assistantMsg := &model.ConversationMessage{
		SessionId: req.SessionId,
		Role:      "assistant",
		Content:   aiResponse,
	}
	if err := repository.ConversationMessageRepo.Create(assistantMsg); err != nil {
		return nil, err
	}

	// Update session message count and title (use first user message as title if default)
	session.MessageCount += 2
	if session.Title == "New Conversation" && len(history) == 1 {
		title := req.Content
		if len(title) > 50 {
			title = title[:50] + "..."
		}
		session.Title = title
	}
	_ = repository.ConversationSessionRepo.Update(session)

	return assistantMsg, nil
}

// GetHistory returns paginated message history for a session
func (s *ConversationService) GetHistory(userId, sessionId uint, page, pageSize int) ([]model.ConversationMessage, int64, error) {
	session, err := repository.ConversationSessionRepo.GetByID(sessionId)
	if err != nil {
		return nil, 0, errors.New("session not found")
	}
	if session.UserId != userId {
		return nil, 0, errors.New("unauthorized to access this session")
	}
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 50
	}
	return repository.ConversationMessageRepo.ListBySession(sessionId, page, pageSize)
}

// ListSessions returns paginated sessions for a user
func (s *ConversationService) ListSessions(userId uint, page, pageSize int) ([]model.ConversationSession, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 20
	}
	return repository.ConversationSessionRepo.ListByUser(userId, page, pageSize)
}

// ResetSession clears all messages in a session
func (s *ConversationService) ResetSession(userId, sessionId uint) error {
	session, err := repository.ConversationSessionRepo.GetByID(sessionId)
	if err != nil {
		return errors.New("session not found")
	}
	if session.UserId != userId {
		return errors.New("unauthorized to access this session")
	}
	if err := repository.ConversationMessageRepo.DeleteBySession(sessionId); err != nil {
		return err
	}
	session.MessageCount = 0
	return repository.ConversationSessionRepo.Update(session)
}

// CloseSession marks a session as inactive
func (s *ConversationService) CloseSession(userId, sessionId uint) error {
	session, err := repository.ConversationSessionRepo.GetByID(sessionId)
	if err != nil {
		return errors.New("session not found")
	}
	if session.UserId != userId {
		return errors.New("unauthorized to access this session")
	}
	session.IsActive = false
	return repository.ConversationSessionRepo.Update(session)
}
