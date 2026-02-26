package mcp

import (
	"edu/model"
	"edu/service"
	"encoding/json"
	"errors"
	"fmt"
)

// getPaperCodeTools returns the list of paper code-related tools
func (s *MCPServer) getPaperCodeTools() []map[string]interface{} {
	return []map[string]interface{}{
		{
			"name":        "paper_code_list",
			"description": "List all paper codes (试卷代码列表). " + FieldDescPaperCode,
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
			"name":        "paper_code_get",
			"description": "Get a paper code by ID (根据ID获取试卷代码). " + FieldDescPaperCode,
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"id":     map[string]interface{}{"type": "number", "description": "Paper code ID"},
					"fields": map[string]interface{}{"type": "array", "items": map[string]interface{}{"type": "string"}, "description": "Fields to return (default: id, name)"},
				},
				"required": []string{"id"},
			},
		},
		{
			"name":        "paper_code_create",
			"description": "Create a new paper code (创建新的试卷代码)",
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"name":       map[string]interface{}{"type": "string", "description": "Paper code name"},
					"syllabusId": map[string]interface{}{"type": "number", "description": "Syllabus ID"},
					"level":      map[string]interface{}{"type": "string", "description": "Syllabus level: 'AS', 'A2', or '' for non-A-Level"},
				},
				"required": []string{"name", "syllabusId"},
			},
		},
		{
			"name":        "paper_code_edit",
			"description": "Edit a paper code (编辑试卷代码)",
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"id":         map[string]interface{}{"type": "number", "description": "Paper code ID"},
					"name":       map[string]interface{}{"type": "string", "description": "New paper code name"},
					"syllabusId": map[string]interface{}{"type": "number", "description": "Syllabus ID"},
					"level":      map[string]interface{}{"type": "string", "description": "Syllabus level: 'AS', 'A2', or '' for non-A-Level"},
				},
				"required": []string{"id"},
			},
		},
		{
			"name":        "paper_code_delete",
			"description": "Delete a paper code (删除试卷代码)",
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"id": map[string]interface{}{"type": "number", "description": "Paper code ID"},
				},
				"required": []string{"id"},
			},
		},
	}
}

// PaperCode tool implementations
func (s *MCPServer) toolPaperCodeList(args map[string]interface{}) (string, error) {
	pageIndex := getInt(args, "pageIndex", 1)
	pageSize := getInt(args, "pageSize", 20)
	syllabusId := getUint(args, "syllabusId", 0)
	fields := parseFields(args)

	query := model.PaperCodeQuery{
		SyllabusId: syllabusId,
		Page:       model.Page{PageIndex: pageIndex, PageSize: pageSize},
	}

	records, total, err := service.QuestionPaperSvr.SelectCodeList(query)
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

func (s *MCPServer) toolPaperCodeGet(args map[string]interface{}) (string, error) {
	id := getUint(args, "id", 0)
	if id == 0 {
		return "", errors.New("id is required")
	}
	fields := parseFields(args)

	record, err := service.QuestionPaperSvr.SelectCodeById(id)
	if err != nil {
		return "", err
	}

	jsonData, _ := json.MarshalIndent(filterFields(record, fields), "", "  ")
	return string(jsonData), nil
}

func (s *MCPServer) toolPaperCodeCreate(args map[string]interface{}) (string, error) {
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

	code := model.PaperCode{
		Name:       name,
		SyllabusId: syllabusId,
		Level:      getString(args, "level", ""),
	}
	err := service.QuestionPaperSvr.CreateCode(code)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Paper code created successfully"), nil
}

func (s *MCPServer) toolPaperCodeEdit(args map[string]interface{}) (string, error) {
	if !s.currentUser.IsAdmin {
		return "", errors.New("permission denied: admin access required")
	}

	id := getUint(args, "id", 0)
	if id == 0 {
		return "", errors.New("id is required")
	}

	// Get existing code
	existing, err := service.QuestionPaperSvr.SelectCodeById(id)
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
	// Allow level to be updated (including clearing it)
	if _, ok := args["level"]; ok {
		existing.Level = getString(args, "level", "")
	}

	err = service.QuestionPaperSvr.EditCode(*existing)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Paper code %d updated successfully", id), nil
}

func (s *MCPServer) toolPaperCodeDelete(args map[string]interface{}) (string, error) {
	if !s.currentUser.IsAdmin {
		return "", errors.New("permission denied: admin access required")
	}

	id := getUint(args, "id", 0)
	if id == 0 {
		return "", errors.New("id is required")
	}

	err := service.QuestionPaperSvr.DeleteCode(id)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Paper code %d deleted successfully", id), nil
}
