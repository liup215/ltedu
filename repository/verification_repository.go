package repository

import (
"edu/model"
"gorm.io/gorm"
)

// IVerificationRepository 验证码数据访问接口
type IVerificationRepository interface {
Create(v *model.Verification) error
FindLastByTarget(target string) (*model.Verification, error)
UpdateStatus(id uint, status int) error
}

type verificationRepository struct {
db *gorm.DB
}

func NewVerificationRepository(db *gorm.DB) IVerificationRepository {
return &verificationRepository{db: db}
}

func (r *verificationRepository) Create(v *model.Verification) error {
return r.db.Create(v).Error
}

func (r *verificationRepository) FindLastByTarget(target string) (*model.Verification, error) {
var ver model.Verification
err := r.db.Where("target = ?", target).Order("created_at desc").First(&ver).Error
if gorm.ErrRecordNotFound == err {
return nil, nil
}
return &ver, err
}

func (r *verificationRepository) UpdateStatus(id uint, status int) error {
return r.db.Model(&model.Verification{}).Where("id = ?", id).Update("status", status).Error
}
