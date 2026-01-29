package repository

import (
	"edu/model"
	"gorm.io/gorm"
)

// IAttachmentRepository 附件数据访问接口
type IAttachmentRepository interface {
	Create(attachment *model.Attachment) error
	Update(attachment *model.Attachment) error
	Delete(id uint) error
	FindByID(id uint) (*model.Attachment, error)
}

type attachmentRepository struct {
	db *gorm.DB
}

// NewAttachmentRepository 创建附件仓储实例
func NewAttachmentRepository(db *gorm.DB) IAttachmentRepository {
	return &attachmentRepository{db: db}
}

// Create 创建附件
func (r *attachmentRepository) Create(attachment *model.Attachment) error {
	return r.db.Create(attachment).Error
}

// Update 更新附件
func (r *attachmentRepository) Update(attachment *model.Attachment) error {
	return r.db.Save(attachment).Error
}

// Delete 删除附件
func (r *attachmentRepository) Delete(id uint) error {
	return r.db.Delete(&model.Attachment{}, id).Error
}

// FindByID 根据ID查询附件
func (r *attachmentRepository) FindByID(id uint) (*model.Attachment, error) {
	var attachment model.Attachment
	err := r.db.Where("id = ?", id).First(&attachment).Error

	if gorm.ErrRecordNotFound == err {
		return nil, nil
	}
	return &attachment, err
}
