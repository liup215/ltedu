package mcp

import (
	"edu/model"
	"edu/service"
	"encoding/json"
	"errors"
	"fmt"
)

// getAvailableTools returns the list of all available MCP tools
func (s *MCPServer) getAvailableTools() []map[string]interface{} {
	return []map[string]interface{}{
		// Organisation (考试局) tools
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

		// Qualification (考试) tools
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

		// Syllabus (考纲) tools
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

		// Chapter (章节) tools
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

		// PaperSeries (考试季) tools
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

		// PaperCode (试卷代码) tools
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

		// PastPaper (真题试卷) tools
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
					"name":           map[string]interface{}{"type": "string", "description": "Past paper name"},
					"syllabusId":     map[string]interface{}{"type": "number", "description": "Syllabus ID"},
					"year":           map[string]interface{}{"type": "number", "description": "Year"},
					"paperCodeId":    map[string]interface{}{"type": "number", "description": "Paper code ID"},
					"paperSeriesId":  map[string]interface{}{"type": "number", "description": "Paper series ID"},
					"questionNumber": map[string]interface{}{"type": "number", "description": "Number of questions"},
				},
				"required": []string{"name", "syllabusId", "year"},
			},
		},
		{
			"name":        "past_paper_edit",
			"description": "Edit a past paper (编辑真题试卷)",
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"id":             map[string]interface{}{"type": "number", "description": "Past paper ID"},
					"name":           map[string]interface{}{"type": "string", "description": "New past paper name"},
					"syllabusId":     map[string]interface{}{"type": "number", "description": "Syllabus ID"},
					"year":           map[string]interface{}{"type": "number", "description": "Year"},
					"paperCodeId":    map[string]interface{}{"type": "number", "description": "Paper code ID"},
					"paperSeriesId":  map[string]interface{}{"type": "number", "description": "Paper series ID"},
					"questionNumber": map[string]interface{}{"type": "number", "description": "Number of questions"},
				},
				"required": []string{"id"},
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

		// Question (试题) tools
		{
			"name":        "question_list",
			"description": "List all questions (试题列表). " + FieldDescQuestion,
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"id":          map[string]interface{}{"type": "number", "description": "Filter by question ID"},
					"stem":        map[string]interface{}{"type": "string", "description": "Filter by question stem (text search)"},
					"syllabusId":  map[string]interface{}{"type": "number", "description": "Filter by syllabus ID"},
					"pastPaperId": map[string]interface{}{"type": "number", "description": "Filter by past paper ID"},
					"difficult":   map[string]interface{}{"type": "number", "description": "Filter by difficulty (1-5)"},
					"status":      map[string]interface{}{"type": "number", "description": "Filter by status (1=normal, 2=forbidden, 3=deleted)"},
					"paperName":   map[string]interface{}{"type": "string", "description": "Filter by paper name"},
					"chapters":    map[string]interface{}{"type": "array", "items": map[string]interface{}{"type": "number"}, "description": "Filter by chapter IDs"},
					"pageIndex":   map[string]interface{}{"type": "number", "description": "Page index (default: 1)"},
					"pageSize":    map[string]interface{}{"type": "number", "description": "Page size (default: 20)"},
					"fields":      map[string]interface{}{"type": "array", "items": map[string]interface{}{"type": "string"}, "description": "Fields to return (default: id, stem)"},
				},
			},
		},
		{
			"name":        "question_get",
			"description": "Get a question by ID (根据ID获取试题). " + FieldDescQuestion,
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"id":     map[string]interface{}{"type": "number", "description": "Question ID"},
					"fields": map[string]interface{}{"type": "array", "items": map[string]interface{}{"type": "string"}, "description": "Fields to return (default: id, stem)"},
				},
				"required": []string{"id"},
			},
		},
		{
			"name":        "question_create",
			"description": "Create a new question (创建新的试题)",
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"syllabusId":       map[string]interface{}{"type": "number", "description": "Syllabus ID"},
					"stem":             map[string]interface{}{"type": "string", "description": "Question stem/content"},
					"difficult":        map[string]interface{}{"type": "number", "description": "Difficulty level (1-5)"},
					"pastPaperId":      map[string]interface{}{"type": "number", "description": "Past paper ID (optional)"},
					"indexInPastPaper": map[string]interface{}{"type": "number", "description": "Index in past paper (optional)"},
				},
				"required": []string{"syllabusId", "stem"},
			},
		},
		{
			"name":        "question_edit",
			"description": "Edit a question (编辑试题)",
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"id":               map[string]interface{}{"type": "number", "description": "Question ID"},
					"syllabusId":       map[string]interface{}{"type": "number", "description": "Syllabus ID"},
					"stem":             map[string]interface{}{"type": "string", "description": "Question stem/content"},
					"difficult":        map[string]interface{}{"type": "number", "description": "Difficulty level (1-5)"},
					"pastPaperId":      map[string]interface{}{"type": "number", "description": "Past paper ID"},
					"indexInPastPaper": map[string]interface{}{"type": "number", "description": "Index in past paper"},
				},
				"required": []string{"id"},
			},
		},
		{
			"name":        "question_delete",
			"description": "Delete a question (删除试题)",
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"id": map[string]interface{}{"type": "number", "description": "Question ID"},
				},
				"required": []string{"id"},
			},
		},
	}
}

// executeTool executes a tool and returns the result as a string
func (s *MCPServer) executeTool(name string, args map[string]interface{}) (string, error) {
	switch name {
	// Organisation tools
	case "organisation_list":
		return s.toolOrganisationList(args)
	case "organisation_get":
		return s.toolOrganisationGet(args)
	case "organisation_create":
		return s.toolOrganisationCreate(args)
	case "organisation_edit":
		return s.toolOrganisationEdit(args)
	case "organisation_delete":
		return s.toolOrganisationDelete(args)

	// Qualification tools
	case "qualification_list":
		return s.toolQualificationList(args)
	case "qualification_get":
		return s.toolQualificationGet(args)
	case "qualification_create":
		return s.toolQualificationCreate(args)
	case "qualification_edit":
		return s.toolQualificationEdit(args)
	case "qualification_delete":
		return s.toolQualificationDelete(args)

	// Syllabus tools
	case "syllabus_list":
		return s.toolSyllabusList(args)
	case "syllabus_get":
		return s.toolSyllabusGet(args)
	case "syllabus_create":
		return s.toolSyllabusCreate(args)
	case "syllabus_edit":
		return s.toolSyllabusEdit(args)
	case "syllabus_delete":
		return s.toolSyllabusDelete(args)

	// Chapter tools
	case "chapter_list":
		return s.toolChapterList(args)
	case "chapter_get":
		return s.toolChapterGet(args)
	case "chapter_create":
		return s.toolChapterCreate(args)
	case "chapter_edit":
		return s.toolChapterEdit(args)
	case "chapter_delete":
		return s.toolChapterDelete(args)

	// PaperSeries tools
	case "paper_series_list":
		return s.toolPaperSeriesList(args)
	case "paper_series_get":
		return s.toolPaperSeriesGet(args)
	case "paper_series_create":
		return s.toolPaperSeriesCreate(args)
	case "paper_series_edit":
		return s.toolPaperSeriesEdit(args)
	case "paper_series_delete":
		return s.toolPaperSeriesDelete(args)

	// PaperCode tools
	case "paper_code_list":
		return s.toolPaperCodeList(args)
	case "paper_code_get":
		return s.toolPaperCodeGet(args)
	case "paper_code_create":
		return s.toolPaperCodeCreate(args)
	case "paper_code_edit":
		return s.toolPaperCodeEdit(args)
	case "paper_code_delete":
		return s.toolPaperCodeDelete(args)

	// PastPaper tools
	case "past_paper_list":
		return s.toolPastPaperList(args)
	case "past_paper_get":
		return s.toolPastPaperGet(args)
	case "past_paper_create":
		return s.toolPastPaperCreate(args)
	case "past_paper_edit":
		return s.toolPastPaperEdit(args)
	case "past_paper_delete":
		return s.toolPastPaperDelete(args)

	// Question tools
	case "question_list":
		return s.toolQuestionList(args)
	case "question_get":
		return s.toolQuestionGet(args)
	case "question_create":
		return s.toolQuestionCreate(args)
	case "question_edit":
		return s.toolQuestionEdit(args)
	case "question_delete":
		return s.toolQuestionDelete(args)

	default:
		return "", errors.New("unknown tool: " + name)
	}
}

// Helper functions to parse arguments
func getFloat(args map[string]interface{}, key string, defaultVal float64) float64 {
	if val, ok := args[key]; ok {
		if floatVal, ok := val.(float64); ok {
			return floatVal
		}
	}
	return defaultVal
}

func getInt(args map[string]interface{}, key string, defaultVal int) int {
	return int(getFloat(args, key, float64(defaultVal)))
}

func getUint(args map[string]interface{}, key string, defaultVal uint) uint {
	return uint(getFloat(args, key, float64(defaultVal)))
}

func getString(args map[string]interface{}, key string, defaultVal string) string {
	if val, ok := args[key]; ok {
		if strVal, ok := val.(string); ok {
			return strVal
		}
	}
	return defaultVal
}

// Organisation tool implementations
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

// Qualification tool implementations
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

	// Get existing past paper
	existing, err := service.QuestionPaperSvr.SelectPastPaperById(id)
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
	if year := getInt(args, "year", 0); year != 0 {
		existing.Year = year
	}
	if codeId := getUint(args, "paperCodeId", 0); codeId != 0 {
		existing.PaperCodeId = codeId
	}
	if seriesId := getUint(args, "paperSeriesId", 0); seriesId != 0 {
		existing.PaperSeriesId = seriesId
	}
	if qNum := getInt(args, "questionNumber", 0); qNum != 0 {
		existing.QuestionNumber = qNum
	}

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

// Question tool implementations
func (s *MCPServer) toolQuestionList(args map[string]interface{}) (string, error) {
	pageIndex := getInt(args, "pageIndex", 1)
	pageSize := getInt(args, "pageSize", 20)
	id := getUint(args, "id", 0)
	stem := getString(args, "stem", "")
	syllabusId := getUint(args, "syllabusId", 0)
	pastPaperId := getUint(args, "pastPaperId", 0)
	difficult := getInt(args, "difficult", 0)
	status := getInt(args, "status", 0)
	paperName := getString(args, "paperName", "")
	fields := parseFields(args)

	// Parse chapters array
	var chapters []uint
	if chaptersVal, ok := args["chapters"]; ok {
		if chaptersArray, ok := chaptersVal.([]interface{}); ok {
			for _, v := range chaptersArray {
				if num, ok := v.(float64); ok {
					chapters = append(chapters, uint(num))
				}
			}
		}
	}

	query := model.QuestionQueryRequest{
		ID:          id,
		Stem:        stem,
		SyllabusId:  syllabusId,
		PastPaperId: pastPaperId,
		Difficult:   difficult,
		Status:      status,
		PaperName:   paperName,
		Chapters:    chapters,
		Page:        model.Page{PageIndex: pageIndex, PageSize: pageSize},
	}

	records, total, err := service.QuestionSvr.SelectQuestionList(query)
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

func (s *MCPServer) toolQuestionGet(args map[string]interface{}) (string, error) {
	id := getUint(args, "id", 0)
	if id == 0 {
		return "", errors.New("id is required")
	}
	fields := parseFields(args)

	record, err := service.QuestionSvr.SelectQuestionById(id)
	if err != nil {
		return "", err
	}

	jsonData, _ := json.MarshalIndent(filterFields(record, fields), "", "  ")
	return string(jsonData), nil
}

func (s *MCPServer) toolQuestionCreate(args map[string]interface{}) (string, error) {
	if !s.currentUser.IsAdmin {
		return "", errors.New("permission denied: admin access required")
	}

	syllabusId := getUint(args, "syllabusId", 0)
	stem := getString(args, "stem", "")
	difficult := getInt(args, "difficult", 1)
	pastPaperId := getUint(args, "pastPaperId", 0)
	indexInPastPaper := getInt(args, "indexInPastPaper", 0)

	if syllabusId == 0 {
		return "", errors.New("syllabusId is required")
	}
	if stem == "" {
		return "", errors.New("stem is required")
	}

	questionID, err := service.QuestionSvr.CreateQuestion(model.Question{
		SyllabusId:       syllabusId,
		Stem:             stem,
		Difficult:        difficult,
		PastPaperId:      pastPaperId,
		IndexInPastPaper: indexInPastPaper,
		Status:           model.QUESTION_STATE_NORMAL,
	}, s.currentUser.ID)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Question created successfully with ID: %d", questionID), nil
}

func (s *MCPServer) toolQuestionEdit(args map[string]interface{}) (string, error) {
	if !s.currentUser.IsAdmin {
		return "", errors.New("permission denied: admin access required")
	}

	id := getUint(args, "id", 0)
	if id == 0 {
		return "", errors.New("id is required")
	}

	// Get existing question
	existing, err := service.QuestionSvr.SelectQuestionById(id)
	if err != nil {
		return "", err
	}

	// Update fields if provided
	if sylId := getUint(args, "syllabusId", 0); sylId != 0 {
		existing.SyllabusId = sylId
	}
	if stem := getString(args, "stem", ""); stem != "" {
		existing.Stem = stem
	}
	if diff := getInt(args, "difficult", 0); diff != 0 {
		existing.Difficult = diff
	}
	if paperId := getUint(args, "pastPaperId", 0); paperId != 0 {
		existing.PastPaperId = paperId
	}
	if idx := getInt(args, "indexInPastPaper", 0); idx != 0 {
		existing.IndexInPastPaper = idx
	}

	err = service.QuestionSvr.EditQuestion(*existing)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Question %d updated successfully", id), nil
}

func (s *MCPServer) toolQuestionDelete(args map[string]interface{}) (string, error) {
	if !s.currentUser.IsAdmin {
		return "", errors.New("permission denied: admin access required")
	}

	id := getUint(args, "id", 0)
	if id == 0 {
		return "", errors.New("id is required")
	}

	err := service.QuestionSvr.DeleteQuestion(id)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Question %d deleted successfully", id), nil
}
