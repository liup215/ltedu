package service

import (
	"edu/model"
	"edu/repository"
)

// AuditLogSvr is the singleton audit log service.
var AuditLogSvr *AuditLogService = &AuditLogService{
	baseService: newBaseService(),
}

// AuditLogService records administrative actions for security auditing.
type AuditLogService struct {
	baseService
}

// Record saves a new audit log entry. Errors are silently swallowed so that
// audit logging failures never block the primary operation.
func (svr *AuditLogService) Record(adminID uint, module, opt, remark, ip string) error {
	entry := &model.AdminLog{
		AdminId: adminID,
		Module:  module,
		Opt:     opt,
		Remark:  remark,
		Ip:      ip,
	}
	return repository.AuditLogRepo.Create(entry)
}

// RecentByUser returns the most recent audit log entries for a given user (max 100).
func (svr *AuditLogService) RecentByUser(adminID uint, limit int) ([]*model.AdminLog, error) {
	if limit <= 0 || limit > 100 {
		limit = 100
	}
	return repository.AuditLogRepo.FindByAdminID(adminID, limit)
}
