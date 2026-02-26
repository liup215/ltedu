package mcp

import (
	"edu/model"
	"edu/service"
	"encoding/json"
	"errors"
	"fmt"
)

// getChapterTools returns the list of chapter-related tools
func (s *MCPServer) getChapterTools() []map[string]interface{} {
	return []map[string]interface{}{
		{
			"name":        "chapter_list",
			"description": "List chapters (章节列表). Supports lazy loading by parentId. " + FieldDescChapter,
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"syllabusId": map[string]interface{}{"type": "number", "description": "Syllabus ID"},
					"parentId":   map[string]interface{}{"type": "number", "description": "Parent chapter ID (optional, 0 or omitted for root chapters)"},
					"pageIndex":  map[string]interface{}{"type": "number", "description": "Page index (default: 1)"},
					"pageSize":   map[string]interface{}{"type": "number", "description": "Page size (default: 20)"},
					"fields":     map[string]interface{}{"type": "array", "items": map[string]interface{}{"type": "string"}, "description": "Fields to return (default: id, name)"},
				},
				"required": []string{"syllabusId"},
			},
		},
		{
			"name":        "chapter_get",
			"description": "Get a chapter by ID (根据ID获取章节). " + FieldDescChapter,
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"id":     map[string]interface{}{"type": "number", "description": "Chapter ID"},
					"fields": map[string]interface{}{"type": "array", "items": map[string]interface{}{"type": "string"}, "description": "Fields to return (default: id, name)"},
				},
				"required": []string{"id"},
			},
		},
		{
			"name":        "chapter_create",
			"description": "Create a new chapter (创建新的章节)",
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"name":       map[string]interface{}{"type": "string", "description": "Chapter name"},
					"syllabusId": map[string]interface{}{"type": "number", "description": "Syllabus ID"},
					"parentId":   map[string]interface{}{"type": "number", "description": "Parent chapter ID (0 for root)"},
					"level":      map[string]interface{}{"type": "string", "description": "Syllabus level: 'AS', 'A2', or '' for non-A-Level"},
				},
				"required": []string{"name", "syllabusId"},
			},
		},
		{
			"name":        "chapter_edit",
			"description": "Edit a chapter (编辑章节)",
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"id":         map[string]interface{}{"type": "number", "description": "Chapter ID"},
					"name":       map[string]interface{}{"type": "string", "description": "New chapter name"},
					"syllabusId": map[string]interface{}{"type": "number", "description": "Syllabus ID"},
					"parentId":   map[string]interface{}{"type": "number", "description": "Parent chapter ID"},
					"level":      map[string]interface{}{"type": "string", "description": "Syllabus level: 'AS', 'A2', or '' for non-A-Level"},
				},
				"required": []string{"id"},
			},
		},
		{
			"name":        "chapter_delete",
			"description": "Delete a chapter (删除章节)",
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"id": map[string]interface{}{"type": "number", "description": "Chapter ID"},
				},
				"required": []string{"id"},
			},
		},
	}
}

// Chapter tool implementations
func (s *MCPServer) toolChapterList(args map[string]interface{}) (string, error) {
	syllabusId := getUint(args, "syllabusId", 0)
	if syllabusId == 0 {
		return "", errors.New("syllabusId is required")
	}
	parentId := getUint(args, "parentId", 0)
	pageIndex := getInt(args, "pageIndex", 1)
	pageSize := getInt(args, "pageSize", 20)
	fields := parseFields(args)

	// Determine FilterRoot based on parentId.
	// If parentId is explicitly passed (even 0), we want to filter by it.
	// In the tool logic, we'll map "no parentId implies children of root" or "root chapters"?
	// The requirement is: "If parentId is missing/0, return root chapters."
	// So if parentId is 0, we set FilterRoot = true.
	// If parentId > 0, we query by ParentId normally (FilterRoot = false).

	filterRoot := false
	if parentId == 0 {
		filterRoot = true
	}

	query := model.ChapterQuery{
		SyllabusId: syllabusId,
		ParentId:   parentId,
		FilterRoot: filterRoot,
		Page:       model.Page{PageIndex: pageIndex, PageSize: pageSize},
	}

	records, total, err := service.QualificationSvr.ChapterList(query)
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

func (s *MCPServer) toolChapterGet(args map[string]interface{}) (string, error) {
	id := getUint(args, "id", 0)
	if id == 0 {
		return "", errors.New("id is required")
	}
	fields := parseFields(args)

	record, err := service.QualificationSvr.SelectChapterById(id)
	if err != nil {
		return "", err
	}

	jsonData, _ := json.MarshalIndent(filterFields(record, fields), "", "  ")
	return string(jsonData), nil
}

func (s *MCPServer) toolChapterCreate(args map[string]interface{}) (string, error) {
	if !s.currentUser.IsAdmin {
		return "", errors.New("permission denied: admin access required")
	}

	name := getString(args, "name", "")
	syllabusId := getUint(args, "syllabusId", 0)
	parentId := getUint(args, "parentId", 0)

	if name == "" {
		return "", errors.New("name is required")
	}
	if syllabusId == 0 {
		return "", errors.New("syllabusId is required")
	}

	chapter, err := service.QualificationSvr.CreateChapter(model.Chapter{
		Name:       name,
		SyllabusId: syllabusId,
		ParentId:   parentId,
		Level:      getString(args, "level", ""),
	})
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Chapter created successfully with ID: %d", chapter.ID), nil
}

func (s *MCPServer) toolChapterEdit(args map[string]interface{}) (string, error) {
	if !s.currentUser.IsAdmin {
		return "", errors.New("permission denied: admin access required")
	}

	id := getUint(args, "id", 0)
	if id == 0 {
		return "", errors.New("id is required")
	}

	// Get existing chapter
	existing, err := service.QualificationSvr.SelectChapterById(id)
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
	// Allow parentId to be updated even to 0
	if _, ok := args["parentId"]; ok {
		existing.ParentId = getUint(args, "parentId", 0)
	}
	// Allow level to be updated (including clearing it)
	if _, ok := args["level"]; ok {
		existing.Level = getString(args, "level", "")
	}

	_, err = service.QualificationSvr.EditChapter(*existing)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Chapter %d updated successfully", id), nil
}

func (s *MCPServer) toolChapterDelete(args map[string]interface{}) (string, error) {
	if !s.currentUser.IsAdmin {
		return "", errors.New("permission denied: admin access required")
	}

	id := getUint(args, "id", 0)
	if id == 0 {
		return "", errors.New("id is required")
	}

	err := service.QualificationSvr.DeleteChapter(id)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Chapter %d deleted successfully", id), nil
}
