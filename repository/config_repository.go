package repository

import (
	"edu/model"
	"gorm.io/gorm"
)

// IAppConfigRepository 配置数据访问接口
type IAppConfigRepository interface {
	Create(config *model.AppConfig) error
	Update(config *model.AppConfig) error
	FindByKey(key string) (*model.AppConfig, error)
	FirstOrCreateByKey(key string) (*model.AppConfig, error)
}

type appConfigRepository struct {
	db *gorm.DB
}

// NewAppConfigRepository 创建配置仓储实例
func NewAppConfigRepository(db *gorm.DB) IAppConfigRepository {
	return &appConfigRepository{db: db}
}

// Create 创建配置
func (r *appConfigRepository) Create(config *model.AppConfig) error {
	return r.db.Create(config).Error
}

// Update 更新配置
func (r *appConfigRepository) Update(config *model.AppConfig) error {
	return r.db.Save(config).Error
}

// FindByKey 根据key查询配置
func (r *appConfigRepository) FindByKey(key string) (*model.AppConfig, error) {
	var config model.AppConfig
	err := r.db.Where("key = ?", key).First(&config).Error

	if gorm.ErrRecordNotFound == err {
		return nil, nil
	}
	return &config, err
}

// FirstOrCreateByKey 根据key查询，不存在则创建
func (r *appConfigRepository) FirstOrCreateByKey(key string) (*model.AppConfig, error) {
	var config model.AppConfig
	err := r.db.Where("key = ?", key).First(&config).Error

	if gorm.ErrRecordNotFound == err {
		// 创建新配置
		config = model.AppConfig{Key: key}
		if err = r.db.Create(&config).Error; err != nil {
			return nil, err
		}
		return &config, nil
	}

	if err != nil {
		return nil, err
	}

	return &config, nil
}
