package model

import "time"

// MCPToken represents a user-specific MCP access token
type MCPToken struct {
	Model
	UserID    uint      `json:"userId" gorm:"index"`
	User      User      `json:"user"`
	Token     string    `json:"token" gorm:"unique;size:64"`
	Name      string    `json:"name" gorm:"size:100"`        // Friendly name for the token
	ExpiresAt time.Time `json:"expiresAt" gorm:"index"`      // Token expiration date
	LastUsed  time.Time `json:"lastUsed"`                    // Last time the token was used
	IsActive  bool      `json:"isActive" gorm:"default:true"` // Whether the token is active
}

// MCPTokenQuery represents query parameters for MCP tokens
type MCPTokenQuery struct {
	ID     uint   `json:"id"`
	UserID uint   `json:"userId"`
	Token  string `json:"token"`
	Page
}

// MCPTokenCreateRequest represents a request to create a new MCP token
type MCPTokenCreateRequest struct {
	Name      string `json:"name" binding:"required"`
	ExpiresAt string `json:"expiresAt"` // Optional, ISO 8601 format
}

// MCPTokenResponse represents an MCP token response
type MCPTokenResponse struct {
	ID        uint      `json:"id"`
	Token     string    `json:"token"`
	Name      string    `json:"name"`
	ExpiresAt time.Time `json:"expiresAt"`
	IsActive  bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
	LastUsed  time.Time `json:"lastUsed"`
}
