package mcp

import (
	"edu/model"
	"edu/service"
	"encoding/json"
	"errors"
	"fmt"
)

func (s *MCPServer) getQualificationTools() []map[string]interface{} {
	return []map[string]interface{}{
		{
			"name":        "qualification_list",
			"description": "List all qualifications (考试列表). " + FieldDescQualification,
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"organisationId": map[string]interface{}{"type": "number", "description": "Filter by organisation ID"},
					"pageIndex":      map[string]interface{}{"type": "number", "description": "Page index (default: 1)"},
					"pageSize":       map[string]interface{}{"type": "number", "description": "Page size (default: 20)"},
					"fields":         map[string]interface{}{"type": "array", "items": map[string]interface{}{"type": "string"}, "description": "Fields to return (default: id, name)"},
				},
			},
		},
		{
			"name":        "qualification_get",
			"description": "Get a qualification by ID (根据ID获取考试). " + FieldDescQualification,
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"id":     map[string]interface{}{"type": "number", "description": "Qualification ID"},
					"fields": map[string]interface{}{"type": "array", "items": map[string]interface{}{"type": "string"}, "description": "Fields to return (default: id, name)"},
				},
				"required": []string{"id"},
			},
		},
		{
			"name":        "qualification_create",
			"description": "Create a new qualification (创建新的考试)",
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"name":           map[string]interface{}{"type": "string", "description": "Qualification name"},
					"organisationId": map[string]interface{}{"type": "number", "description": "Organisation ID"},
				},
				"required": []string{"name", "organisationId"},
			},
		},
		{
			"name":        "qualification_edit",
			"description": "Edit a qualification (编辑考试)",
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"id":             map[string]interface{}{"type": "number", "description": "Qualification ID"},
					"name":           map[string]interface{}{"type": "string", "description": "New qualification name"},
					"organisationId": map[string]interface{}{"type": "number", "description": "Organisation ID"},
				},
				"required": []string{"id"},
			},
		},
		{
			"name":        "qualification_delete",
			"description": "Delete a qualification (删除考试)",
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"id": map[string]interface{}{"type": "number", "description": "Qualification ID"},
				},
				"required": []string{"id"},
			},
		},
	}
}

func (s *MCPServer) toolQualificationList(args map[string]interface{}) (string, error) {
	pageIndex := getInt(args, "pageIndex", 1)
	pageSize := getInt(args, "pageSize", 20)
	organisationId := getUint(args, "organisationId", 0)
	fields := parseFields(args)

	query := model.QualificationQuery{
		OrganisationId: organisationId,
		Page:           model.Page{PageIndex: pageIndex, PageSize: pageSize},
	}

	records, total, err := service.QualificationSvr.SelectQualificationList(query)
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

func (s *MCPServer) toolQualificationGet(args map[string]interface{}) (string, error) {
	id := getUint(args, "id", 0)
	if id == 0 {
		return "", errors.New("id is required")
	}
	fields := parseFields(args)

	record, err := service.QualificationSvr.SelectQualificationById(id)
	if err != nil {
		return "", err
	}

	jsonData, _ := json.MarshalIndent(filterFields(record, fields), "", "  ")
	return string(jsonData), nil
}

func (s *MCPServer) toolQualificationCreate(args map[string]interface{}) (string, error) {
	if !s.currentUser.IsAdmin {
		return "", errors.New("permission denied: admin access required")
	}

	name := getString(args, "name", "")
	organisationId := getUint(args, "organisationId", 0)

	if name == "" {
		return "", errors.New("name is required")
	}
	if organisationId == 0 {
		return "", errors.New("organisationId is required")
	}

	qual, err := service.QualificationSvr.CreateQualification(model.Qualification{
		Name:           name,
		OrganisationId: organisationId,
	})
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Qualification created successfully with ID: %d", qual.ID), nil
}

func (s *MCPServer) toolQualificationEdit(args map[string]interface{}) (string, error) {
	if !s.currentUser.IsAdmin {
		return "", errors.New("permission denied: admin access required")
	}

	id := getUint(args, "id", 0)
	if id == 0 {
		return "", errors.New("id is required")
	}

	// Get existing qualification to preserve fields not being updated
	existing, err := service.QualificationSvr.SelectQualificationById(id)
	if err != nil {
		return "", err
	}

	// Update fields if provided
	if name := getString(args, "name", ""); name != "" {
		existing.Name = name
	}
	if orgId := getUint(args, "organisationId", 0); orgId != 0 {
		existing.OrganisationId = orgId
	}

	_, err = service.QualificationSvr.EditQualification(*existing)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Qualification %d updated successfully", id), nil
}

func (s *MCPServer) toolQualificationDelete(args map[string]interface{}) (string, error) {
	if !s.currentUser.IsAdmin {
		return "", errors.New("permission denied: admin access required")
	}

	id := getUint(args, "id", 0)
	if id == 0 {
		return "", errors.New("id is required")
	}

	err := service.QualificationSvr.DeleteQualification(id)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Qualification %d deleted successfully", id), nil
}
