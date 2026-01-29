package service

import (
	"crypto/rand"
	"edu/model"
	"edu/repository"
	"encoding/hex"
	"errors"
	"time"
)

var MCPTokenSvr *MCPTokenService

func init() {
	MCPTokenSvr = &MCPTokenService{}
}

type MCPTokenService struct{}

// GenerateTokenString generates a secure random token string
func (s *MCPTokenService) GenerateTokenString() (string, error) {
	bytes := make([]byte, 32) // 32 bytes = 64 hex characters
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// CreateToken creates a new MCP token for a user
func (s *MCPTokenService) CreateToken(userID uint, name string, expiresAt time.Time) (*model.MCPToken, error) {
	tokenStr, err := s.GenerateTokenString()
	if err != nil {
		return nil, err
	}

	// Default expiration: 1 year from now
	if expiresAt.IsZero() {
		expiresAt = time.Now().AddDate(1, 0, 0)
	}

	token := &model.MCPToken{
		UserID:    userID,
		Token:     tokenStr,
		Name:      name,
		ExpiresAt: expiresAt,
		LastUsed:  time.Now(),
		IsActive:  true,
	}

	if err := repository.MCPTokenRepo.Create(token); err != nil {
		return nil, err
	}

	return token, nil
}

// ValidateToken validates an MCP token and returns the associated user
func (s *MCPTokenService) ValidateToken(tokenStr string) (*model.User, error) {
	token, err := repository.MCPTokenRepo.ValidateToken(tokenStr)
	if err != nil {
		return nil, err
	}

	// Update last used timestamp
	_ = repository.MCPTokenRepo.UpdateLastUsed(token.ID)

	// Return the user associated with this token
	if token.User.ID == 0 {
		return nil, errors.New("user not found")
	}

	return &token.User, nil
}

// ListUserTokens lists all MCP tokens for a user
func (s *MCPTokenService) ListUserTokens(userID uint, page model.Page) ([]model.MCPToken, int64, error) {
	return repository.MCPTokenRepo.ListByUserID(userID, page)
}

// DeleteToken deletes an MCP token
func (s *MCPTokenService) DeleteToken(tokenID uint, userID uint) error {
	token, err := repository.MCPTokenRepo.GetByID(tokenID)
	if err != nil {
		return err
	}

	// Ensure the token belongs to the user
	if token.UserID != userID {
		return errors.New("unauthorized")
	}

	return repository.MCPTokenRepo.Delete(tokenID)
}

// DeactivateToken deactivates an MCP token
func (s *MCPTokenService) DeactivateToken(tokenID uint, userID uint) error {
	token, err := repository.MCPTokenRepo.GetByID(tokenID)
	if err != nil {
		return err
	}

	// Ensure the token belongs to the user
	if token.UserID != userID {
		return errors.New("unauthorized")
	}

	return repository.MCPTokenRepo.Deactivate(tokenID)
}

// ActivateToken activates an MCP token
func (s *MCPTokenService) ActivateToken(tokenID uint, userID uint) error {
	token, err := repository.MCPTokenRepo.GetByID(tokenID)
	if err != nil {
		return err
	}

	// Ensure the token belongs to the user
	if token.UserID != userID {
		return errors.New("unauthorized")
	}

	return repository.MCPTokenRepo.Activate(tokenID)
}

// ListAllTokens lists all MCP tokens (for admin)
func (s *MCPTokenService) ListAllTokens(query model.MCPTokenQuery) ([]model.MCPToken, int64, error) {
	return repository.MCPTokenRepo.ListAll(query)
}

// AdminDeleteToken deletes any MCP token (for admin)
func (s *MCPTokenService) AdminDeleteToken(tokenID uint) error {
	return repository.MCPTokenRepo.Delete(tokenID)
}

// AdminDeactivateToken deactivates any MCP token (for admin)
func (s *MCPTokenService) AdminDeactivateToken(tokenID uint) error {
	return repository.MCPTokenRepo.Deactivate(tokenID)
}

// AdminActivateToken activates any MCP token (for admin)
func (s *MCPTokenService) AdminActivateToken(tokenID uint) error {
	return repository.MCPTokenRepo.Activate(tokenID)
}
