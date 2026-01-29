package repository

import (
"edu/model"
"gorm.io/gorm"
)

// IOrganisationRepository 考试局数据访问接口
type IOrganisationRepository interface {
Create(org *model.Organisation) error
Update(org *model.Organisation) error
Delete(id uint) error
FindByID(id uint) (*model.Organisation, error)
FindByName(name string) ([]*model.Organisation, error)
FindPage(query *model.OrganisationQuery, offset, limit int) ([]*model.Organisation, int64, error)
FindAll(query *model.OrganisationQuery) ([]*model.Organisation, error)
}

type organisationRepository struct {
db *gorm.DB
}

func NewOrganisationRepository(db *gorm.DB) IOrganisationRepository {
return &organisationRepository{db: db}
}

func (r *organisationRepository) Create(org *model.Organisation) error {
return r.db.Create(org).Error
}

func (r *organisationRepository) Update(org *model.Organisation) error {
return r.db.Model(org).Updates(org).Error
}

func (r *organisationRepository) Delete(id uint) error {
return r.db.Delete(&model.Organisation{}, id).Error
}

func (r *organisationRepository) FindByID(id uint) (*model.Organisation, error) {
var org model.Organisation
err := r.db.Where("id = ?", id).First(&org).Error
if gorm.ErrRecordNotFound == err {
return nil, nil
}
return &org, err
}

func (r *organisationRepository) FindByName(name string) ([]*model.Organisation, error) {
var orgs []*model.Organisation
err := r.db.Where("name LIKE ?", "%"+name+"%").Find(&orgs).Error
return orgs, err
}

func (r *organisationRepository) FindPage(query *model.OrganisationQuery, offset, limit int) ([]*model.Organisation, int64, error) {
var orgs []*model.Organisation
var total int64

q := r.db.Model(&model.Organisation{})

if query.ID != 0 {
q = q.Where("id = ?", query.ID)
}
if query.Name != "" {
q = q.Where("name LIKE ?", "%"+query.Name+"%")
}

q.Count(&total)
err := q.Order("id DESC").Offset(offset).Limit(limit).Find(&orgs).Error
return orgs, total, err
}

func (r *organisationRepository) FindAll(query *model.OrganisationQuery) ([]*model.Organisation, error) {
var orgs []*model.Organisation
q := r.db.Model(&model.Organisation{})

if query.ID != 0 {
q = q.Where("id = ?", query.ID)
}
if query.Name != "" {
q = q.Where("name LIKE ?", "%"+query.Name+"%")
}

err := q.Order("id DESC").Find(&orgs).Error
return orgs, err
}
