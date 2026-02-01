package mcp

import (
	"edu/model"
	"edu/service"
	"encoding/json"
	"errors"
	"fmt"
)

// getPastPaperTools returns the tools for PastPaper management
func (s *MCPServer) getPastPaperTools() []map[string]interface{} {
	return []map[string]interface{}{
		{
			"name":        "past_paper_list",
			"description": "List all past papers (真题试卷列表). " + FieldDescPastPaper,
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"id":            map[string]interface{}{"type": "number", "description": "Filter by past paper ID"},
					"name":          map[string]interface{}{"type": "string", "description": "Filter by past paper name (text search)"},
					"syllabusId":    map[string]interface{}{"type": "number", "description": "Filter by syllabus ID"},
					"year":          map[string]interface{}{"type": "number", "description": "Filter by year"},
					"paperCodeId":   map[string]interface{}{"type": "number", "description": "Filter by paper code ID"},
					"paperSeriesId": map[string]interface{}{"type": "number", "description": "Filter by paper series ID"},
					"pageIndex":     map[string]interface{}{"type": "number", "description": "Page index (default: 1)"},
					"pageSize":      map[string]interface{}{"type": "number", "description": "Page size (default: 20)"},
					"fields":        map[string]interface{}{"type": "array", "items": map[string]interface{}{"type": "string"}, "description": "Fields to return (default: id, name)"},
				},
			},
		},
		{
			"name":        "past_paper_get",
			"description": "Get a past paper by ID (根据ID获取真题试卷). " + FieldDescPastPaper,
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"id":     map[string]interface{}{"type": "number", "description": "Past paper ID"},
					"fields": map[string]interface{}{"type": "array", "items": map[string]interface{}{"type": "string"}, "description": "Fields to return (default: id, name)"},
				},
				"required": []string{"id"},
			},
		},
		{
			"name":        "past_paper_create",
			"description": "Create a new past paper (创建新的真题试卷)",
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"name":           map[string]interface{}{"type": "string", "description": "Past paper name (Required)"},
					"syllabusId":     map[string]interface{}{"type": "number", "description": "Syllabus ID (Required)"},
					"year":           map[string]interface{}{"type": "number", "description": "Year (Required)"},
					"paperCodeId":    map[string]interface{}{"type": "number", "description": "Paper code ID (Required)"},
					"paperSeriesId":  map[string]interface{}{"type": "number", "description": "Paper series ID (Required)"},
					"questionNumber": map[string]interface{}{"type": "number", "description": "Number of questions (Required)"},
				},
				"required": []string{"name", "syllabusId", "year", "paperCodeId", "paperSeriesId", "questionNumber"},
			},
		},
		{
			"name":        "past_paper_edit",
			"description": "Edit a past paper (编辑真题试卷)",
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"id":             map[string]interface{}{"type": "number", "description": "Past paper ID (Required)"},
					"name":           map[string]interface{}{"type": "string", "description": "New past paper name (Required)"},
					"syllabusId":     map[string]interface{}{"type": "number", "description": "Syllabus ID (Required)"},
					"year":           map[string]interface{}{"type": "number", "description": "Year (Required)"},
					"paperCodeId":    map[string]interface{}{"type": "number", "description": "Paper code ID (Required)"},
					"paperSeriesId":  map[string]interface{}{"type": "number", "description": "Paper series ID (Required)"},
					"questionNumber": map[string]interface{}{"type": "number", "description": "Number of questions (Required)"},
				},
				"required": []string{"id", "name", "syllabusId", "year", "paperCodeId", "paperSeriesId", "questionNumber"},
			},
		},
		{
			"name":        "past_paper_delete",
			"description": "Delete a past paper (删除真题试卷)",
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"id": map[string]interface{}{"type": "number", "description": "Past paper ID"},
				},
				"required": []string{"id"},
			},
		},
	}
}

// PastPaper tool implementations
func (s *MCPServer) toolPastPaperList(args map[string]interface{}) (string, error) {
	pageIndex := getInt(args, "pageIndex", 1)
	pageSize := getInt(args, "pageSize", 20)
	id := getUint(args, "id", 0)
	name := getString(args, "name", "")
	syllabusId := getUint(args, "syllabusId", 0)
	year := getInt(args, "year", 0)
	paperCodeId := getUint(args, "paperCodeId", 0)
	paperSeriesId := getUint(args, "paperSeriesId", 0)
	fields := parseFields(args)

	query := model.PastPaperQuery{
		ID:            id,
		Name:          name,
		SyllabusId:    syllabusId,
		Year:          year,
		PaperCodeId:   paperCodeId,
		PaperSeriesId: paperSeriesId,
		Page:          model.Page{PageIndex: pageIndex, PageSize: pageSize},
	}

	records, total, err := service.QuestionPaperSvr.SelectPastPaperList(query)
	if err != nil {
		return "", err
	}

	result := map[string]interface{}{
		"total":   total,
		"records": filterFields(records, fields),
	}
	jsonData, _ := json.MarshalIndent(result, "", "  ")
	return string(jsonData), nil
}

func (s *MCPServer) toolPastPaperGet(args map[string]interface{}) (string, error) {
	id := getUint(args, "id", 0)
	if id == 0 {
		return "", errors.New("id is required")
	}
	fields := parseFields(args)

	record, err := service.QuestionPaperSvr.SelectPastPaperById(id)
	if err != nil {
		return "", err
	}

	jsonData, _ := json.MarshalIndent(filterFields(record, fields), "", "  ")
	return string(jsonData), nil
}

func (s *MCPServer) toolPastPaperCreate(args map[string]interface{}) (string, error) {
	if !s.currentUser.IsAdmin {
		return "", errors.New("permission denied: admin access required")
	}

	name := getString(args, "name", "")
	syllabusId := getUint(args, "syllabusId", 0)
	year := getInt(args, "year", 0)
	paperCodeId := getUint(args, "paperCodeId", 0)
	paperSeriesId := getUint(args, "paperSeriesId", 0)
	questionNumber := getInt(args, "questionNumber", 0)

	if name == "" {
		return "", errors.New("name is required")
	}
	if syllabusId == 0 {
		return "", errors.New("syllabusId is required")
	}
	if year == 0 {
		return "", errors.New("year is required")
	}
	if paperCodeId == 0 {
		return "", errors.New("paperCodeId is required")
	}
	if paperSeriesId == 0 {
		return "", errors.New("paperSeriesId is required")
	}
	if questionNumber == 0 {
		return "", errors.New("questionNumber is required")
	}

	paper, err := service.QuestionPaperSvr.CreatePastPaper(model.PastPaper{
		Name:           name,
		SyllabusId:     syllabusId,
		Year:           year,
		PaperCodeId:    paperCodeId,
		PaperSeriesId:  paperSeriesId,
		QuestionNumber: questionNumber,
	})
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Past paper created successfully with ID: %d", paper.ID), nil
}

func (s *MCPServer) toolPastPaperEdit(args map[string]interface{}) (string, error) {
	if !s.currentUser.IsAdmin {
		return "", errors.New("permission denied: admin access required")
	}

	id := getUint(args, "id", 0)
	if id == 0 {
		return "", errors.New("id is required")
	}

	// All parameters are required, so check them all
	name := getString(args, "name", "")
	syllabusId := getUint(args, "syllabusId", 0)
	year := getInt(args, "year", 0)
	paperCodeId := getUint(args, "paperCodeId", 0)
	paperSeriesId := getUint(args, "paperSeriesId", 0)
	questionNumber := getInt(args, "questionNumber", 0)

	if name == "" {
		return "", errors.New("name is required")
	}
	if syllabusId == 0 {
		return "", errors.New("syllabusId is required")
	}
	if year == 0 {
		return "", errors.New("year is required")
	}
	if paperCodeId == 0 {
		return "", errors.New("paperCodeId is required")
	}
	if paperSeriesId == 0 {
		return "", errors.New("paperSeriesId is required")
	}
	if questionNumber == 0 {
		return "", errors.New("questionNumber is required")
	}

	// Get existing past paper
	existing, err := service.QuestionPaperSvr.SelectPastPaperById(id)
	if err != nil {
		return "", err
	}

	// Update all fields
	existing.Name = name
	existing.SyllabusId = syllabusId
	existing.Year = year
	existing.PaperCodeId = paperCodeId
	existing.PaperSeriesId = paperSeriesId
	existing.QuestionNumber = questionNumber

	err = service.QuestionPaperSvr.EditPastPaper(*existing)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Past paper %d updated successfully", id), nil
}

func (s *MCPServer) toolPastPaperDelete(args map[string]interface{}) (string, error) {
	if !s.currentUser.IsAdmin {
		return "", errors.New("permission denied: admin access required")
	}

	id := getUint(args, "id", 0)
	if id == 0 {
		return "", errors.New("id is required")
	}

	err := service.QuestionPaperSvr.DeletePastPaper(id)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Past paper %d deleted successfully", id), nil
}
