package model

type DocumentCategory struct {
	Model
	Icon        string `json:"icon"`
	Cover       string `json:"cover"`
	ParentId    uint   `json:"parentId"`
	Name        string `json:"name"`
	DocCount    int    `json:"docCount"`
	Sort        int    `json:"sort"`
	Description string `json:"description"`
}

func (c *DocumentCategory) GetResponse() *DocumentCategoryQueryResponse {
	return &DocumentCategoryQueryResponse{
		ID:          c.ID,
		Icon:        c.Icon,
		Cover:       c.Cover,
		ParentId:    c.ParentId,
		Name:        c.Name,
		DocCount:    c.DocCount,
		Sort:        c.Sort,
		Description: c.Description,
	}
}

type DocumentCategoryCreateEditRequest struct {
	ID          uint   `json:"id"`
	Icon        string `json:"icon"`
	Cover       string `json:"cover"`
	ParentId    uint   `json:"parentId"`
	Name        string `json:"name"`
	Sort        int    `json:"sort"`
	Enable      int    `json:"enable"`
	Description string `json:"description"`
}

func (c *DocumentCategoryCreateEditRequest) GetCategory() DocumentCategory {
	return DocumentCategory{
		Model:       Model{ID: c.ID},
		Icon:        c.Icon,
		Cover:       c.Cover,
		ParentId:    c.ParentId,
		Name:        c.Name,
		Sort:        c.Sort,
		Description: c.Description,
	}
}

type DocumentCategoryQueryRequest struct {
	ID       uint `json:"id"`
	ParentId uint `json:"parentId"`
	Page
}

type DocumentCategoryQueryResponse struct {
	ID          uint   `json:"id"`
	Icon        string `json:"icon"`
	Cover       string `json:"cover"`
	ParentId    uint   `json:"parentId"`
	Name        string `json:"name"`
	DocCount    int    `json:"docCount"`
	Sort        int    `json:"sort"`
	Description string `json:"description"`
}
