package repository

import (
	"edu/model"

	"gorm.io/gorm"
)

// IAuditLogRepository defines persistence operations for AdminLog audit records.
type IAuditLogRepository interface {
	Create(log *model.AdminLog) error
	FindByAdminID(adminID uint, limit int) ([]*model.AdminLog, error)
}

type auditLogRepository struct {
	db *gorm.DB
}

// NewAuditLogRepository creates an audit log repository instance.
func NewAuditLogRepository(db *gorm.DB) IAuditLogRepository {
	return &auditLogRepository{db: db}
}

// Create persists a new audit log record.
func (r *auditLogRepository) Create(log *model.AdminLog) error {
	return r.db.Create(log).Error
}

// FindByAdminID returns the most recent audit log entries for a given admin user.
func (r *auditLogRepository) FindByAdminID(adminID uint, limit int) ([]*model.AdminLog, error) {
	var logs []*model.AdminLog
	err := r.db.Where("admin_id = ?", adminID).
		Order("created_at DESC").
		Limit(limit).
		Find(&logs).Error
	return logs, err
}
