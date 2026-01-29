package repository

import (
"edu/model"
"gorm.io/gorm"
)

// IVocabularySetRepository 单词集数据访问接口
type IVocabularySetRepository interface {
Create(set *model.VocabularySet) error
Update(set *model.VocabularySet) error
Delete(id uint) error
FindByID(id uint) (*model.VocabularySet, error)
FindPage(query *model.VocabularySetQuery, offset, limit int) ([]*model.VocabularySet, int64, error)
FindAll(query *model.VocabularySetQuery) ([]*model.VocabularySet, error)
FindByName(name string) (*model.VocabularySet, error)
}

type vocabularySetRepository struct {
db *gorm.DB
}

func NewVocabularySetRepository(db *gorm.DB) IVocabularySetRepository {
return &vocabularySetRepository{db: db}
}

func (r *vocabularySetRepository) Create(set *model.VocabularySet) error {
return r.db.Create(set).Error
}

func (r *vocabularySetRepository) Update(set *model.VocabularySet) error {
return r.db.Model(set).Updates(set).Error
}

func (r *vocabularySetRepository) Delete(id uint) error {
return r.db.Delete(&model.VocabularySet{}, id).Error
}

func (r *vocabularySetRepository) FindByID(id uint) (*model.VocabularySet, error) {
var set model.VocabularySet
err := r.db.Where("id = ?", id).First(&set).Error
if gorm.ErrRecordNotFound == err {
return nil, nil
}
return &set, err
}

func (r *vocabularySetRepository) FindPage(query *model.VocabularySetQuery, offset, limit int) ([]*model.VocabularySet, int64, error) {
var sets []*model.VocabularySet
var total int64

q := r.db.Model(&model.VocabularySet{})

if query.ID != 0 {
q = q.Where("id = ?", query.ID)
}
if query.SyllabusId != 0 {
q = q.Where("syllabus_id = ?", query.SyllabusId)
}
if query.Name != "" {
q = q.Where("name LIKE ?", "%"+query.Name+"%")
}

q.Count(&total)
err := q.Order("id DESC").Offset(offset).Limit(limit).Find(&sets).Error
return sets, total, err
}

func (r *vocabularySetRepository) FindAll(query *model.VocabularySetQuery) ([]*model.VocabularySet, error) {
var sets []*model.VocabularySet
q := r.db.Model(&model.VocabularySet{})

if query.ID != 0 {
q = q.Where("id = ?", query.ID)
}
if query.SyllabusId != 0 {
q = q.Where("syllabus_id = ?", query.SyllabusId)
}
if query.Name != "" {
q = q.Where("name LIKE ?", "%"+query.Name+"%")
}

err := q.Order("id DESC").Find(&sets).Error
return sets, err
}

func (r *vocabularySetRepository) FindByName(name string) (*model.VocabularySet, error) {
var set model.VocabularySet
err := r.db.Where("name = ?", name).First(&set).Error
if gorm.ErrRecordNotFound == err {
return nil, nil
}
return &set, err
}

// IVocabularyItemRepository 单词项数据访问接口
type IVocabularyItemRepository interface {
Create(item *model.VocabularyItem) error
Update(item *model.VocabularyItem) error
Delete(id uint) error
FindByID(id uint) (*model.VocabularyItem, error)
FindPage(query *model.VocabularyItemQuery, offset, limit int) ([]*model.VocabularyItem, int64, error)
FindAll(query *model.VocabularyItemQuery) ([]*model.VocabularyItem, error)
Count(query *model.VocabularyItemQuery) (int64, error)
}

type vocabularyItemRepository struct {
db *gorm.DB
}

func NewVocabularyItemRepository(db *gorm.DB) IVocabularyItemRepository {
return &vocabularyItemRepository{db: db}
}

func (r *vocabularyItemRepository) Create(item *model.VocabularyItem) error {
return r.db.Create(item).Error
}

func (r *vocabularyItemRepository) Update(item *model.VocabularyItem) error {
return r.db.Model(item).Updates(item).Error
}

func (r *vocabularyItemRepository) Delete(id uint) error {
return r.db.Delete(&model.VocabularyItem{}, id).Error
}

func (r *vocabularyItemRepository) FindByID(id uint) (*model.VocabularyItem, error) {
var item model.VocabularyItem
err := r.db.Where("id = ?", id).First(&item).Error
if gorm.ErrRecordNotFound == err {
return nil, nil
}
return &item, err
}

func (r *vocabularyItemRepository) FindPage(query *model.VocabularyItemQuery, offset, limit int) ([]*model.VocabularyItem, int64, error) {
var items []*model.VocabularyItem
var total int64

q := r.db.Model(&model.VocabularyItem{})

if query.ID != 0 {
q = q.Where("id = ?", query.ID)
}
if query.VocabularySetId != 0 {
q = q.Where("vocabulary_set_id = ?", query.VocabularySetId)
}
if query.Key != "" {
q = q.Where("key LIKE ?", "%"+query.Key+"%")
}

q.Count(&total)
err := q.Order("id DESC").Offset(offset).Limit(limit).Find(&items).Error
return items, total, err
}

func (r *vocabularyItemRepository) FindAll(query *model.VocabularyItemQuery) ([]*model.VocabularyItem, error) {
var items []*model.VocabularyItem
q := r.db.Model(&model.VocabularyItem{})

if query.ID != 0 {
q = q.Where("id = ?", query.ID)
}
if query.VocabularySetId != 0 {
q = q.Where("vocabulary_set_id = ?", query.VocabularySetId)
}
if query.Key != "" {
q = q.Where("key LIKE ?", "%"+query.Key+"%")
}

err := q.Order("id DESC").Find(&items).Error
return items, err
}

func (r *vocabularyItemRepository) Count(query *model.VocabularyItemQuery) (int64, error) {
var total int64
q := r.db.Model(&model.VocabularyItem{})

if query.ID != 0 {
q = q.Where("id = ?", query.ID)
}
if query.VocabularySetId != 0 {
q = q.Where("vocabulary_set_id = ?", query.VocabularySetId)
}
if query.Key != "" {
q = q.Where("key LIKE ?", "%"+query.Key+"%")
}

err := q.Count(&total).Error
return total, err
}
