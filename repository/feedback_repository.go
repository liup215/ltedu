package repository

import (
	"edu/model"

	"gorm.io/gorm"
)

type IFeedbackRepository interface {
	Create(feedback *model.UserFeedback) error
	FindByID(id uint) (*model.UserFeedback, error)
	FindAll(req model.FeedbackListRequest) ([]*model.UserFeedback, int64, error)
	FindByUserID(userID uint, page model.Page) ([]*model.UserFeedback, int64, error)
	UpdateStatus(id uint, status, adminNote string) error
	GetStats() (*model.FeedbackStats, error)
}

type feedbackRepository struct {
	db *gorm.DB
}

func NewFeedbackRepository(db *gorm.DB) IFeedbackRepository {
	return &feedbackRepository{db: db}
}

func (r *feedbackRepository) Create(feedback *model.UserFeedback) error {
	return r.db.Create(feedback).Error
}

func (r *feedbackRepository) FindByID(id uint) (*model.UserFeedback, error) {
	var feedback model.UserFeedback
	err := r.db.First(&feedback, id).Error
	if err != nil {
		return nil, err
	}
	return &feedback, nil
}

func (r *feedbackRepository) FindAll(req model.FeedbackListRequest) ([]*model.UserFeedback, int64, error) {
	page := req.Page.CheckPage()
	query := r.db.Model(&model.UserFeedback{})

	if req.Status != "" {
		query = query.Where("`status` = ?", req.Status)
	}
	if req.Type != "" {
		query = query.Where("`type` = ?", req.Type)
	}
	if req.UserID > 0 {
		query = query.Where("user_id = ?", req.UserID)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var items []*model.UserFeedback
	offset := (page.PageIndex - 1) * page.PageSize
	err := query.Order("id DESC").Offset(offset).Limit(page.PageSize).Find(&items).Error
	return items, total, err
}

func (r *feedbackRepository) FindByUserID(userID uint, page model.Page) ([]*model.UserFeedback, int64, error) {
	p := page.CheckPage()
	var total int64
	if err := r.db.Model(&model.UserFeedback{}).Where("user_id = ?", userID).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var items []*model.UserFeedback
	offset := (p.PageIndex - 1) * p.PageSize
	err := r.db.Where("user_id = ?", userID).Order("id DESC").Offset(offset).Limit(p.PageSize).Find(&items).Error
	return items, total, err
}

func (r *feedbackRepository) UpdateStatus(id uint, status, adminNote string) error {
	updates := map[string]interface{}{"status": status}
	if adminNote != "" {
		updates["admin_note"] = adminNote
	}
	return r.db.Model(&model.UserFeedback{}).Where("id = ?", id).Updates(updates).Error
}

func (r *feedbackRepository) GetStats() (*model.FeedbackStats, error) {
	stats := &model.FeedbackStats{
		ByType:      make(map[string]int),
		BySentiment: make(map[string]int),
		ByStatus:    make(map[string]int),
	}

	if err := r.db.Model(&model.UserFeedback{}).Count(&stats.Total).Error; err != nil {
		return nil, err
	}

	type countResult struct {
		Key   string
		Count int
	}

	var typeResults []countResult
	if err := r.db.Model(&model.UserFeedback{}).Select("`type` as key, count(*) as count").Group("`type`").Scan(&typeResults).Error; err != nil {
		return nil, err
	}
	for _, r := range typeResults {
		stats.ByType[r.Key] = r.Count
	}

	var sentimentResults []countResult
	if err := r.db.Model(&model.UserFeedback{}).Select("sentiment as key, count(*) as count").Group("sentiment").Scan(&sentimentResults).Error; err != nil {
		return nil, err
	}
	for _, r := range sentimentResults {
		stats.BySentiment[r.Key] = r.Count
	}

	var statusResults []countResult
	if err := r.db.Model(&model.UserFeedback{}).Select("`status` as key, count(*) as count").Group("`status`").Scan(&statusResults).Error; err != nil {
		return nil, err
	}
	for _, r := range statusResults {
		stats.ByStatus[r.Key] = r.Count
	}

	type avgResult struct {
		Avg float64
	}
	var avgRes avgResult
	r.db.Model(&model.UserFeedback{}).Where("rating > 0").Select("AVG(rating) as avg").Scan(&avgRes)
	stats.AvgRating = avgRes.Avg

	return stats, nil
}
