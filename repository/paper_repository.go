package repository

import (
	"errors"

	"edu/model"
	"gorm.io/gorm"
)

// IPaperSeriesRepository 试卷系列数据访问接口
type IPaperSeriesRepository interface {
	Create(series *model.PaperSeries) error
	Update(series *model.PaperSeries) error
	Delete(id uint) error
	FindByID(id uint) (*model.PaperSeries, error)
	FindList(query model.PaperSeriesQuery) ([]*model.PaperSeries, int64, error)
	FindAll(query model.PaperSeriesQuery) ([]*model.PaperSeries, error)
}

// IPaperCodeRepository 试卷代码数据访问接口
type IPaperCodeRepository interface {
	Create(code *model.PaperCode) error
	Update(code *model.PaperCode) error
	Delete(id uint) error
	FindByID(id uint) (*model.PaperCode, error)
	FindList(query model.PaperCodeQuery) ([]*model.PaperCode, int64, error)
	FindAll(query model.PaperCodeQuery) ([]*model.PaperCode, error)
}

type paperSeriesRepository struct {
	db *gorm.DB
}

type paperCodeRepository struct {
	db *gorm.DB
}

// NewPaperSeriesRepository 创建试卷系列仓储实例
func NewPaperSeriesRepository(db *gorm.DB) IPaperSeriesRepository {
	return &paperSeriesRepository{db: db}
}

// NewPaperCodeRepository 创建试卷代码仓储实例
func NewPaperCodeRepository(db *gorm.DB) IPaperCodeRepository {
	return &paperCodeRepository{db: db}
}

// ============ PaperSeries 实现 ============

// Create 创建试卷系列
func (r *paperSeriesRepository) Create(series *model.PaperSeries) error {
	return r.db.Create(series).Error
}

// Update 更新试卷系列
func (r *paperSeriesRepository) Update(series *model.PaperSeries) error {
	updates := map[string]interface{}{
		"name":        series.Name,
		"syllabus_id": series.SyllabusId,
	}
	return r.db.Model(&model.PaperSeries{}).
		Where("id = ?", series.ID).
		Updates(updates).Error
}

// Delete 删除试卷系列
func (r *paperSeriesRepository) Delete(id uint) error {
	return r.db.Delete(&model.PaperSeries{}, id).Error
}

// FindByID 根据ID查询试卷系列
func (r *paperSeriesRepository) FindByID(id uint) (*model.PaperSeries, error) {
	var series model.PaperSeries
	err := r.buildSeriesQuery(model.PaperSeriesQuery{ID: id}).First(&series).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &series, err
}

// FindList 分页查询试卷系列
func (r *paperSeriesRepository) FindList(query model.PaperSeriesQuery) ([]*model.PaperSeries, int64, error) {
	var series []*model.PaperSeries
	var total int64

	q := r.buildSeriesQuery(query)
	q.Model(&model.PaperSeries{}).Count(&total)

	page := query.CheckPage()
	err := q.
		Offset((page.PageIndex - 1) * page.PageSize).
		Limit(page.PageSize).
		Find(&series).Error

	return series, total, err
}

// FindAll 查询所有试卷系列
func (r *paperSeriesRepository) FindAll(query model.PaperSeriesQuery) ([]*model.PaperSeries, error) {
	var series []*model.PaperSeries
	err := r.buildSeriesQuery(query).Find(&series).Error
	return series, err
}

// buildSeriesQuery 构建试卷系列查询
func (r *paperSeriesRepository) buildSeriesQuery(query model.PaperSeriesQuery) *gorm.DB {
	q := r.db.Model(&model.PaperSeries{})

	if query.ID != 0 {
		q = q.Where("id = ?", query.ID)
	}
	if query.SyllabusId != 0 {
		q = q.Where("syllabus_id = ?", query.SyllabusId)
	}
	if query.Name != "" {
		q = q.Where("name LIKE ?", "%"+query.Name+"%")
	}

	return q.
		Preload("Syllabus").
		Preload("Syllabus.Qualification").
		Preload("Syllabus.Qualification.Organisation").
		Order("id DESC")
}

// ============ PaperCode 实现 ============

// Create 创建试卷代码
func (r *paperCodeRepository) Create(code *model.PaperCode) error {
	return r.db.Create(code).Error
}

// Update 更新试卷代码
func (r *paperCodeRepository) Update(code *model.PaperCode) error {
	updates := map[string]interface{}{
		"name":        code.Name,
		"syllabus_id": code.SyllabusId,
	}
	if code.ExamNodeId != 0 {
		updates["exam_node_id"] = code.ExamNodeId
	}
	return r.db.Model(&model.PaperCode{}).
		Where("id = ?", code.ID).
		Updates(updates).Error
}

// Delete 删除试卷代码
func (r *paperCodeRepository) Delete(id uint) error {
	return r.db.Delete(&model.PaperCode{}, id).Error
}

// FindByID 根据ID查询试卷代码
func (r *paperCodeRepository) FindByID(id uint) (*model.PaperCode, error) {
	var code model.PaperCode
	err := r.buildCodeQuery(model.PaperCodeQuery{ID: id}).First(&code).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &code, err
}

// FindList 分页查询试卷代码
func (r *paperCodeRepository) FindList(query model.PaperCodeQuery) ([]*model.PaperCode, int64, error) {
	var codes []*model.PaperCode
	var total int64

	q := r.buildCodeQuery(query)
	q.Model(&model.PaperCode{}).Count(&total)

	page := query.CheckPage()
	err := q.
		Offset((page.PageIndex - 1) * page.PageSize).
		Limit(page.PageSize).
		Find(&codes).Error

	return codes, total, err
}

// FindAll 查询所有试卷代码
func (r *paperCodeRepository) FindAll(query model.PaperCodeQuery) ([]*model.PaperCode, error) {
	var codes []*model.PaperCode
	err := r.buildCodeQuery(query).Find(&codes).Error
	return codes, err
}

// buildCodeQuery 构建试卷代码查询
func (r *paperCodeRepository) buildCodeQuery(query model.PaperCodeQuery) *gorm.DB {
	q := r.db.Model(&model.PaperCode{})

	if query.ID != 0 {
		q = q.Where("id = ?", query.ID)
	}
	if query.SyllabusId != 0 {
		q = q.Where("syllabus_id = ?", query.SyllabusId)
	}
	if query.Name != "" {
		q = q.Where("name LIKE ?", "%"+query.Name+"%")
	}

	return q.
		Preload("Syllabus").
		Preload("Syllabus.Qualification").
		Preload("Syllabus.Qualification.Organisation").
		Preload("ExamNode").
		Order("id DESC")
}
