package mcp

import (
	"edu/model"
	"edu/service"
	"encoding/json"
	"errors"
	"fmt"
)

// getSyllabusTools returns the list of syllabus-related tools
func (s *MCPServer) getSyllabusTools() []map[string]interface{} {
	return []map[string]interface{}{
		{
			"name":        "syllabus_list",
			"description": "List all syllabuses (考纲列表). " + FieldDescSyllabus,
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"qualificationId": map[string]interface{}{"type": "number", "description": "Filter by qualification ID"},
					"pageIndex":       map[string]interface{}{"type": "number", "description": "Page index (default: 1)"},
					"pageSize":        map[string]interface{}{"type": "number", "description": "Page size (default: 20)"},
					"fields":          map[string]interface{}{"type": "array", "items": map[string]interface{}{"type": "string"}, "description": "Fields to return (default: id, name)"},
				},
			},
		},
		{
			"name":        "syllabus_get",
			"description": "Get a syllabus by ID (根据ID获取考纲). " + FieldDescSyllabus,
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"id":     map[string]interface{}{"type": "number", "description": "Syllabus ID"},
					"fields": map[string]interface{}{"type": "array", "items": map[string]interface{}{"type": "string"}, "description": "Fields to return (default: id, name)"},
				},
				"required": []string{"id"},
			},
		},
		{
			"name":        "syllabus_create",
			"description": "Create a new syllabus (创建新的考纲)",
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"name":            map[string]interface{}{"type": "string", "description": "Syllabus name"},
					"code":            map[string]interface{}{"type": "string", "description": "Syllabus code"},
					"qualificationId": map[string]interface{}{"type": "number", "description": "Qualification ID"},
				},
				"required": []string{"name", "qualificationId"},
			},
		},
		{
			"name":        "syllabus_edit",
			"description": "Edit a syllabus (编辑考纲)",
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"id":              map[string]interface{}{"type": "number", "description": "Syllabus ID"},
					"name":            map[string]interface{}{"type": "string", "description": "New syllabus name"},
					"code":            map[string]interface{}{"type": "string", "description": "Syllabus code"},
					"qualificationId": map[string]interface{}{"type": "number", "description": "Qualification ID"},
				},
				"required": []string{"id"},
			},
		},
		{
			"name":        "syllabus_delete",
			"description": "Delete a syllabus (删除考纲)",
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"id": map[string]interface{}{"type": "number", "description": "Syllabus ID"},
				},
				"required": []string{"id"},
			},
		},
	}
}

// Syllabus tool implementations
func (s *MCPServer) toolSyllabusList(args map[string]interface{}) (string, error) {
	pageIndex := getInt(args, "pageIndex", 1)
	pageSize := getInt(args, "pageSize", 20)
	qualificationId := getUint(args, "qualificationId", 0)
	fields := parseFields(args)

	query := model.SyllabusQuery{
		QualificationId: qualificationId,
		Page:            model.Page{PageIndex: pageIndex, PageSize: pageSize},
	}

	records, total, err := service.QualificationSvr.SelectSyllabusList(query)
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

func (s *MCPServer) toolSyllabusGet(args map[string]interface{}) (string, error) {
	id := getUint(args, "id", 0)
	if id == 0 {
		return "", errors.New("id is required")
	}
	fields := parseFields(args)

	record, err := service.QualificationSvr.SelectSyllabusById(id)
	if err != nil {
		return "", err
	}

	jsonData, _ := json.MarshalIndent(filterFields(record, fields), "", "  ")
	return string(jsonData), nil
}

func (s *MCPServer) toolSyllabusCreate(args map[string]interface{}) (string, error) {
	if !s.currentUser.IsAdmin {
		return "", errors.New("permission denied: admin access required")
	}

	name := getString(args, "name", "")
	qualificationId := getUint(args, "qualificationId", 0)
	code := getString(args, "code", "")

	if name == "" {
		return "", errors.New("name is required")
	}
	if qualificationId == 0 {
		return "", errors.New("qualificationId is required")
	}

	syl, err := service.QualificationSvr.CreateSyllabus(model.Syllabus{
		Name:            name,
		Code:            code,
		QualificationId: qualificationId,
	})
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Syllabus created successfully with ID: %d", syl.ID), nil
}

func (s *MCPServer) toolSyllabusEdit(args map[string]interface{}) (string, error) {
	if !s.currentUser.IsAdmin {
		return "", errors.New("permission denied: admin access required")
	}

	id := getUint(args, "id", 0)
	if id == 0 {
		return "", errors.New("id is required")
	}

	// Get existing syllabus
	existing, err := service.QualificationSvr.SelectSyllabusById(id)
	if err != nil {
		return "", err
	}

	// Update fields if provided
	if name := getString(args, "name", ""); name != "" {
		existing.Name = name
	}
	if code := getString(args, "code", ""); code != "" {
		existing.Code = code
	}
	if qualId := getUint(args, "qualificationId", 0); qualId != 0 {
		existing.QualificationId = qualId
	}

	_, err = service.QualificationSvr.EditSyllabus(*existing)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Syllabus %d updated successfully", id), nil
}

func (s *MCPServer) toolSyllabusDelete(args map[string]interface{}) (string, error) {
	if !s.currentUser.IsAdmin {
		return "", errors.New("permission denied: admin access required")
	}

	id := getUint(args, "id", 0)
	if id == 0 {
		return "", errors.New("id is required")
	}

	err := service.QualificationSvr.DeleteSyllabus(id)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Syllabus %d deleted successfully", id), nil
}
