package mcp

import (
	"edu/model"
	"edu/service"
	"encoding/json"
	"errors"
	"fmt"
)

// getPaperSeriesTools returns the list of paper series-related tools
func (s *MCPServer) getPaperSeriesTools() []map[string]interface{} {
	return []map[string]interface{}{
		{
			"name":        "paper_series_list",
			"description": "List all paper series (考试季列表). " + FieldDescPaperSeries,
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"syllabusId": map[string]interface{}{"type": "number", "description": "Filter by syllabus ID"},
					"pageIndex":  map[string]interface{}{"type": "number", "description": "Page index (default: 1)"},
					"pageSize":   map[string]interface{}{"type": "number", "description": "Page size (default: 20)"},
					"fields":     map[string]interface{}{"type": "array", "items": map[string]interface{}{"type": "string"}, "description": "Fields to return (default: id, name)"},
				},
			},
		},
		{
			"name":        "paper_series_get",
			"description": "Get a paper series by ID (根据ID获取考试季). " + FieldDescPaperSeries,
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"id":     map[string]interface{}{"type": "number", "description": "Paper series ID"},
					"fields": map[string]interface{}{"type": "array", "items": map[string]interface{}{"type": "string"}, "description": "Fields to return (default: id, name)"},
				},
				"required": []string{"id"},
			},
		},
		{
			"name":        "paper_series_create",
			"description": "Create a new paper series (创建新的考试季)",
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"name":       map[string]interface{}{"type": "string", "description": "Paper series name"},
					"syllabusId": map[string]interface{}{"type": "number", "description": "Syllabus ID"},
				},
				"required": []string{"name", "syllabusId"},
			},
		},
		{
			"name":        "paper_series_edit",
			"description": "Edit a paper series (编辑考试季)",
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"id":         map[string]interface{}{"type": "number", "description": "Paper series ID"},
					"name":       map[string]interface{}{"type": "string", "description": "New paper series name"},
					"syllabusId": map[string]interface{}{"type": "number", "description": "Syllabus ID"},
				},
				"required": []string{"id"},
			},
		},
		{
			"name":        "paper_series_delete",
			"description": "Delete a paper series (删除考试季)",
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"id": map[string]interface{}{"type": "number", "description": "Paper series ID"},
				},
				"required": []string{"id"},
			},
		},
	}
}

// PaperSeries tool implementations
func (s *MCPServer) toolPaperSeriesList(args map[string]interface{}) (string, error) {
	pageIndex := getInt(args, "pageIndex", 1)
	pageSize := getInt(args, "pageSize", 20)
	syllabusId := getUint(args, "syllabusId", 0)
	fields := parseFields(args)

	query := model.PaperSeriesQuery{
		SyllabusId: syllabusId,
		Page:       model.Page{PageIndex: pageIndex, PageSize: pageSize},
	}

	records, total, err := service.QuestionPaperSvr.SelectSeriesList(query)
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

func (s *MCPServer) toolPaperSeriesGet(args map[string]interface{}) (string, error) {
	id := getUint(args, "id", 0)
	if id == 0 {
		return "", errors.New("id is required")
	}
	fields := parseFields(args)

	record, err := service.QuestionPaperSvr.SelectSeriesById(id)
	if err != nil {
		return "", err
	}

	jsonData, _ := json.MarshalIndent(filterFields(record, fields), "", "  ")
	return string(jsonData), nil
}

func (s *MCPServer) toolPaperSeriesCreate(args map[string]interface{}) (string, error) {
	if !s.currentUser.IsAdmin {
		return "", errors.New("permission denied: admin access required")
	}

	name := getString(args, "name", "")
	syllabusId := getUint(args, "syllabusId", 0)

	if name == "" {
		return "", errors.New("name is required")
	}
	if syllabusId == 0 {
		return "", errors.New("syllabusId is required")
	}

	series := model.PaperSeries{
		Name:       name,
		SyllabusId: syllabusId,
	}
	err := service.QuestionPaperSvr.CreateSeries(series)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Paper series created successfully"), nil
}

func (s *MCPServer) toolPaperSeriesEdit(args map[string]interface{}) (string, error) {
	if !s.currentUser.IsAdmin {
		return "", errors.New("permission denied: admin access required")
	}

	id := getUint(args, "id", 0)
	if id == 0 {
		return "", errors.New("id is required")
	}

	// Get existing series
	existing, err := service.QuestionPaperSvr.SelectSeriesById(id)
	if err != nil {
		return "", err
	}

	// Update fields if provided
	if name := getString(args, "name", ""); name != "" {
		existing.Name = name
	}
	if sylId := getUint(args, "syllabusId", 0); sylId != 0 {
		existing.SyllabusId = sylId
	}

	err = service.QuestionPaperSvr.EditSeries(*existing)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Paper series %d updated successfully", id), nil
}

func (s *MCPServer) toolPaperSeriesDelete(args map[string]interface{}) (string, error) {
	if !s.currentUser.IsAdmin {
		return "", errors.New("permission denied: admin access required")
	}

	id := getUint(args, "id", 0)
	if id == 0 {
		return "", errors.New("id is required")
	}

	err := service.QuestionPaperSvr.DeleteSeries(id)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Paper series %d deleted successfully", id), nil
}
