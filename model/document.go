package model

const (
	// 封面，按照A4纸的尺寸比例
	DocumentCoverWidth  = 210
	DocumentCoverHeight = 297
)

const (
	DocumentStatusPending    = iota // 待转换
	DocumentStatusConverting        // 转换中
	DocumentStatusConverted         // 已转换
	DocumentStatusFailed            // 转换失败
	DocumentStatusDisabled          // 已禁用
	DocumentStatusRePending         // 重新等待转换
)

var DocumentStatusMap = map[int]struct{}{
	DocumentStatusPending:    {},
	DocumentStatusConverting: {},
	DocumentStatusConverted:  {},
	DocumentStatusFailed:     {},
	DocumentStatusDisabled:   {},
}

type Document struct {
	Model
	Name               string           `json:"name"`
	Keywords           string           `json:"keywords"`
	Description        string           `json:"description"`
	UserId             int64            `json:"userId"`
	Width              int              `json:"width"`
	Height             int              `json:"height"`
	Preview            int              `json:"preview"`
	Pages              int              `json:"pages"`
	DownloadCount      int              `json:"downloadCount"`
	ViewCount          int              `json:"viewCount"`
	FavoriteCount      int              `json:"favoriteCount"`
	CommentCount       int              `json:"commentCount"`
	Score              int              `json:"score"`
	ScoreCount         int              `json:"scoreCount"`
	Price              int              `json:"price"`
	Size               int64            `json:"size"`
	Ext                string           `json:"ext"`
	Status             int              `json:"status"`
	DeletedUserId      int64            `json:"deletedUserId"`
	EnableGZIP         bool             `json:"enableGZIP"`
	PreviewExt         string           `json:"previewExt"`
	DocumentCategoryId uint             `json:"documentCategoryId"`
	DocumentCategory   DocumentCategory `json:"documentCategory"`
	SyllabusId         uint             `json:"syllabusId"`
	Syllabus           Syllabus         `json:"syllabus"`
	AttachmentId       uint             `json:"attachmentId"`
	Attachment         Attachment       `json:"attachment"`
	Cover              string           `json:"cover"`
}

func (d *Document) GetResponse() *DocumentQueryResponse {
	return &DocumentQueryResponse{
		ID:                   d.ID,
		Name:                 d.Name,
		Keywords:             d.Keywords,
		Description:          d.Description,
		Width:                d.Width,
		Height:               d.Height,
		Preview:              d.Preview,
		Pages:                d.Pages,
		DownloadCount:        d.DownloadCount,
		ViewCount:            d.ViewCount,
		FavoriteCount:        d.FavoriteCount,
		CommentCount:         d.CommentCount,
		Score:                d.Score,
		ScoreCount:           d.ScoreCount,
		Price:                d.Price,
		Size:                 d.Size,
		Ext:                  d.Ext,
		Status:               d.Status,
		DeletedUserId:        d.DeletedUserId,
		EnableGZIP:           d.EnableGZIP,
		PreviewExt:           d.PreviewExt,
		DocumentCategoryId:   d.DocumentCategoryId,
		DocumentCategoryName: d.DocumentCategory.Name,
		OrganisationId:       d.Syllabus.Qualification.OrganisationId,
		OrganisationName:     d.Syllabus.Qualification.Organisation.Name,
		QualificationId:      d.Syllabus.QualificationId,
		QualificationName:    d.Syllabus.Qualification.Name,
		SyllabusId:           d.SyllabusId,
		SyllabusName:         d.Syllabus.Name,
		AttachmentId:         d.AttachmentId,
		AttachmentHash:       d.Attachment.Hash,
	}
}

type DocumentCreateEditRequest struct {
	ID                 uint   `json:"id"`
	Name               string `json:"name"`
	Keywords           string `json:"keywords"`
	Description        string `json:"description"`
	DocumentCategoryId uint   `json:"documentCategoryId"`
	Price              int    `json:"price"`
	SyllabusId         uint   `json:"syllabusId"`
	AttachmentId       uint   `json:"attachmentId"`
}

func (c *DocumentCreateEditRequest) GetDocument() Document {
	return Document{
		Model:              Model{ID: c.ID},
		Name:               c.Name,
		Keywords:           c.Keywords,
		Description:        c.Description,
		DocumentCategoryId: c.DocumentCategoryId,
		Price:              c.Price,
		SyllabusId:         c.SyllabusId,
		AttachmentId:       c.AttachmentId,
	}
}

type DocumentQueryRequest struct {
	ID                 uint `json:"id"`
	DocumentCategoryId uint `json:"documentCategoryId"`
	SyllabusId         uint `json:"syllabusId"`
	Page
}

type DocumentQueryResponse struct {
	ID                   uint   `json:"id"`
	Name                 string `json:"name"`
	Keywords             string `json:"keywords"`
	Description          string `json:"description"`
	Width                int    `json:"width"`
	Height               int    `json:"height"`
	Preview              int    `json:"preview"`
	Pages                int    `json:"pages"`
	DownloadCount        int    `json:"downloadCount"`
	ViewCount            int    `json:"viewCount"`
	FavoriteCount        int    `json:"favoriteCount"`
	CommentCount         int    `json:"commentCount"`
	Score                int    `json:"score"`
	ScoreCount           int    `json:"scoreCount"`
	Price                int    `json:"price"`
	Size                 int64  `json:"size"`
	Ext                  string `json:"ext"`
	Status               int    `json:"status"`
	DeletedUserId        int64  `json:"deletedUserId"`
	EnableGZIP           bool   `json:"enableGZIP"`
	PreviewExt           string `json:"previewExt"`
	DocumentCategoryId   uint   `json:"documentCategoryId"`
	DocumentCategoryName string `json:"documentCategoryName"`
	OrganisationId       uint   `json:"organisationId"`
	OrganisationName     string `json:"organisationName"`
	QualificationId      uint   `json:"qualificationId"`
	QualificationName    string `json:"qualificationName"`
	SyllabusId           uint   `json:"syllabusId"`
	SyllabusName         string `json:"syllabusName"`
	SyllabusCode         string `json:"syllabusCode"`
	AttachmentId         uint   `json:"attachmentId"`
	AttachmentHash       string `json:"attachmentHash"`
}
