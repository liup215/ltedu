package repository

import (
	"edu/model"
	"errors"
	"time"

	"gorm.io/gorm"
)

// IMCPTokenRepository MCP Token数据访问接口
type IMCPTokenRepository interface {
	Create(token *model.MCPToken) error
	GetByToken(tokenStr string) (*model.MCPToken, error)
	GetByID(id uint) (*model.MCPToken, error)
	ListByUserID(userID uint, page model.Page) ([]model.MCPToken, int64, error)
	ListAll(query model.MCPTokenQuery) ([]model.MCPToken, int64, error)
	UpdateLastUsed(id uint) error
	Delete(id uint) error
	ValidateToken(tokenStr string) (*model.MCPToken, error)
	Deactivate(id uint) error
	Activate(id uint) error
}

type mcpTokenRepository struct {
	db *gorm.DB
}

// NewMCPTokenRepository creates a new MCP token repository
func NewMCPTokenRepository(db *gorm.DB) IMCPTokenRepository {
	return &mcpTokenRepository{db: db}
}

// Create creates a new MCP token
func (r *mcpTokenRepository) Create(token *model.MCPToken) error {
	return r.db.Create(token).Error
}

// GetByToken finds an MCP token by its token string
func (r *mcpTokenRepository) GetByToken(tokenStr string) (*model.MCPToken, error) {
	var token model.MCPToken
	err := r.db.Preload("User").Where("token = ?", tokenStr).First(&token).Error
	if err != nil {
		return nil, err
	}
	return &token, nil
}

// GetByID finds an MCP token by its ID
func (r *mcpTokenRepository) GetByID(id uint) (*model.MCPToken, error) {
	var token model.MCPToken
	err := r.db.Preload("User").First(&token, id).Error
	if err != nil {
		return nil, err
	}
	return &token, nil
}

// ListByUserID lists all MCP tokens for a given user
func (r *mcpTokenRepository) ListByUserID(userID uint, page model.Page) ([]model.MCPToken, int64, error) {
	var tokens []model.MCPToken
	var total int64

	page = page.CheckPage()
	offset := (page.PageIndex - 1) * page.PageSize

	query := r.db.Model(&model.MCPToken{}).Where("user_id = ?", userID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(page.PageSize).Order("created_at DESC").Find(&tokens).Error; err != nil {
		return nil, 0, err
	}

	return tokens, total, nil
}

// UpdateLastUsed updates the last used timestamp for a token
func (r *mcpTokenRepository) UpdateLastUsed(id uint) error {
	return r.db.Model(&model.MCPToken{}).Where("id = ?", id).Update("last_used", time.Now()).Error
}

// Delete deletes an MCP token
func (r *mcpTokenRepository) Delete(id uint) error {
	return r.db.Delete(&model.MCPToken{}, id).Error
}

// ValidateToken checks if a token is valid (exists, active, and not expired)
func (r *mcpTokenRepository) ValidateToken(tokenStr string) (*model.MCPToken, error) {
	token, err := r.GetByToken(tokenStr)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	if !token.IsActive {
		return nil, errors.New("token is inactive")
	}

	if time.Now().After(token.ExpiresAt) {
		return nil, errors.New("token has expired")
	}

	return token, nil
}

// Deactivate deactivates an MCP token
func (r *mcpTokenRepository) Deactivate(id uint) error {
	return r.db.Model(&model.MCPToken{}).Where("id = ?", id).Update("is_active", false).Error
}

// Activate activates an MCP token
func (r *mcpTokenRepository) Activate(id uint) error {
	return r.db.Model(&model.MCPToken{}).Where("id = ?", id).Update("is_active", true).Error
}

// ListAll lists all MCP tokens with optional filtering (for admin)
func (r *mcpTokenRepository) ListAll(query model.MCPTokenQuery) ([]model.MCPToken, int64, error) {
	var tokens []model.MCPToken
	var total int64

	query.Page = query.Page.CheckPage()
	offset := (query.PageIndex - 1) * query.PageSize

	dbQuery := r.db.Model(&model.MCPToken{}).Preload("User")

	// Filter by user ID if provided
	if query.UserID > 0 {
		dbQuery = dbQuery.Where("user_id = ?", query.UserID)
	}

	if err := dbQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := dbQuery.Offset(offset).Limit(query.PageSize).Order("created_at DESC").Find(&tokens).Error; err != nil {
		return nil, 0, err
	}

	return tokens, total, nil
}
