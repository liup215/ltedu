package mcp

import (
	"edu/model"
	"edu/service"
	"encoding/json"
	"errors"
	"fmt"
)

func (s *MCPServer) getOrganisationTools() []map[string]interface{} {
	return []map[string]interface{}{
		{
			"name":        "organisation_list",
			"description": "List all exam organisations (考试局列表). " + FieldDescOrganisation,
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"pageIndex": map[string]interface{}{"type": "number", "description": "Page index (default: 1)"},
					"pageSize":  map[string]interface{}{"type": "number", "description": "Page size (default: 20)"},
					"fields":    map[string]interface{}{"type": "array", "items": map[string]interface{}{"type": "string"}, "description": "Fields to return (default: id, name)"},
				},
			},
		},
		{
			"name":        "organisation_get",
			"description": "Get an exam organisation by ID (根据ID获取考试局). " + FieldDescOrganisation,
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"id":     map[string]interface{}{"type": "number", "description": "Organisation ID"},
					"fields": map[string]interface{}{"type": "array", "items": map[string]interface{}{"type": "string"}, "description": "Fields to return (default: id, name)"},
				},
				"required": []string{"id"},
			},
		},
		{
			"name":        "organisation_create",
			"description": "Create a new exam organisation (创建新的考试局)",
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"name": map[string]interface{}{"type": "string", "description": "Organisation name"},
				},
				"required": []string{"name"},
			},
		},
		{
			"name":        "organisation_edit",
			"description": "Edit an exam organisation (编辑考试局)",
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"id":   map[string]interface{}{"type": "number", "description": "Organisation ID"},
					"name": map[string]interface{}{"type": "string", "description": "New organisation name"},
				},
				"required": []string{"id", "name"},
			},
		},
		{
			"name":        "organisation_delete",
			"description": "Delete an exam organisation (删除考试局)",
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"id": map[string]interface{}{"type": "number", "description": "Organisation ID"},
				},
				"required": []string{"id"},
			},
		},
	}
}

func (s *MCPServer) toolOrganisationList(args map[string]interface{}) (string, error) {
	pageIndex := getInt(args, "pageIndex", 1)
	pageSize := getInt(args, "pageSize", 20)
	fields := parseFields(args)

	query := model.OrganisationQuery{
		Page: model.Page{PageIndex: pageIndex, PageSize: pageSize},
	}

	records, total, err := service.QualificationSvr.SelectOrganisationList(query)
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

func (s *MCPServer) toolOrganisationGet(args map[string]interface{}) (string, error) {
	id := getUint(args, "id", 0)
	if id == 0 {
		return "", errors.New("id is required")
	}
	fields := parseFields(args)

	record, err := service.QualificationSvr.SelectOrganisationById(id)
	if err != nil {
		return "", err
	}

	jsonData, _ := json.MarshalIndent(filterFields(record, fields), "", "  ")
	return string(jsonData), nil
}

func (s *MCPServer) toolOrganisationCreate(args map[string]interface{}) (string, error) {
	if !s.currentUser.IsAdmin {
		return "", errors.New("permission denied: admin access required")
	}

	name := getString(args, "name", "")
	if name == "" {
		return "", errors.New("name is required")
	}

	org, err := service.QualificationSvr.CreateOrganisation(model.Organisation{Name: name})
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Organisation created successfully with ID: %d", org.ID), nil
}

func (s *MCPServer) toolOrganisationEdit(args map[string]interface{}) (string, error) {
	if !s.currentUser.IsAdmin {
		return "", errors.New("permission denied: admin access required")
	}

	id := getUint(args, "id", 0)
	if id == 0 {
		return "", errors.New("id is required")
	}

	name := getString(args, "name", "")
	if name == "" {
		return "", errors.New("name is required")
	}

	_, err := service.QualificationSvr.EditOrganisation(model.Organisation{
		Model: model.Model{ID: id},
		Name:  name,
	})
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Organisation %d updated successfully", id), nil
}

func (s *MCPServer) toolOrganisationDelete(args map[string]interface{}) (string, error) {
	if !s.currentUser.IsAdmin {
		return "", errors.New("permission denied: admin access required")
	}

	id := getUint(args, "id", 0)
	if id == 0 {
		return "", errors.New("id is required")
	}

	err := service.QualificationSvr.DeleteOrganisation(id)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Organisation %d deleted successfully", id), nil
}
