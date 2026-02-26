package mcp

import (
	"edu/model"
	"edu/service"
	"encoding/json"
	"errors"
	"fmt"
)

// getQuestionTools returns the tools for Question management
func (s *MCPServer) getQuestionTools() []map[string]interface{} {
	return []map[string]interface{}{
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
					"syllabusId":       map[string]interface{}{"type": "number", "description": "Syllabus ID (Required)"},
					"stem":             map[string]interface{}{"type": "string", "description": "Question stem/content (Required)"},
					"difficult":        map[string]interface{}{"type": "number", "description": "Difficulty level (1-5) (Required)"},
					"pastPaperId":      map[string]interface{}{"type": "number", "description": "Past paper ID (Required)"},
					"indexInPastPaper": map[string]interface{}{"type": "number", "description": "Index in past paper (Required)"},
					"questionContents": map[string]interface{}{
						"type": "array",
						"description": `Question contents array (Required). Each item represents a question part.
Common fields: score (int), partLabel (string), subpartLabel (string), analyze (string), questionTypeId (int: 1=SingleChoice, 2=MultipleChoice, 3=TrueFalse, 4=GapFilling, 5=ShortAnswer).

Type-specific fields (include one based on questionTypeId):
1. SingleChoice: "singleChoice": {"options": [{"prefix": "A", "content": "text"}], "answer": "A"}
2. MultipleChoice: "multipleChoice": {"options": [{"prefix": "A", "content": "text"}], "answer": ["A", "B"]}
3. TrueFalse: "trueOrFalse": {"answer": 1} (1=True, 2=False)
4. GapFilling: "gapFilling": {"answer": ["text1", "text2"]}
5. ShortAnswer: "shortAnswer": {"answer": "text"}`,
						"items": map[string]interface{}{"type": "object"},
					},
				},
				"required": []string{"syllabusId", "stem", "difficult", "pastPaperId", "indexInPastPaper", "questionContents"},
			},
		},
		{
			"name":        "question_edit",
			"description": "Edit a question (编辑试题)",
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"id":               map[string]interface{}{"type": "number", "description": "Question ID (Required)"},
					"syllabusId":       map[string]interface{}{"type": "number", "description": "Syllabus ID (Required)"},
					"stem":             map[string]interface{}{"type": "string", "description": "Question stem/content (Required)"},
					"difficult":        map[string]interface{}{"type": "number", "description": "Difficulty level (1-5) (Required)"},
					"pastPaperId":      map[string]interface{}{"type": "number", "description": "Past paper ID (Required)"},
					"indexInPastPaper": map[string]interface{}{"type": "number", "description": "Index in past paper (Required)"},
					"questionContents": map[string]interface{}{
						"type": "array",
						"description": `Question contents array (Required). Each item represents a question part.
Common fields: score (int), partLabel (string), subpartLabel (string), analyze (string), questionTypeId (int: 1=SingleChoice, 2=MultipleChoice, 3=TrueFalse, 4=GapFilling, 5=ShortAnswer).

Type-specific fields (include one based on questionTypeId):
1. SingleChoice: "singleChoice": {"options": [{"prefix": "A", "content": "text"}], "answer": "A"}
2. MultipleChoice: "multipleChoice": {"options": [{"prefix": "A", "content": "text"}], "answer": ["A", "B"]}
3. TrueFalse: "trueOrFalse": {"answer": 1} (1=True, 2=False)
4. GapFilling: "gapFilling": {"answer": ["text1", "text2"]}
5. ShortAnswer: "shortAnswer": {"answer": "text"}`,
						"items": map[string]interface{}{"type": "object"},
					},
				},
				"required": []string{"id", "syllabusId", "stem", "difficult", "pastPaperId", "indexInPastPaper", "questionContents"},
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

	query := model.QuestionQueryRequest{
		ID:          id,
		Stem:        stem,
		SyllabusId:  syllabusId,
		PastPaperId: pastPaperId,
		Difficult:   difficult,
		Status:      status,
		PaperName:   paperName,
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

// Helper to parse questionContents from interface{} to []model.QuestionContent
func parseQuestionContents(input interface{}) ([]model.QuestionContent, error) {
	if input == nil {
		return nil, nil
	}
	data, err := json.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal questionContents: %v", err)
	}
	var contents []model.QuestionContent
	if err := json.Unmarshal(data, &contents); err != nil {
		return nil, fmt.Errorf("failed to unmarshal questionContents to struct: %v", err)
	}
	return contents, nil
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
	questionContentsVal, ok := args["questionContents"]

	if syllabusId == 0 {
		return "", errors.New("syllabusId is required")
	}
	if stem == "" {
		return "", errors.New("stem is required")
	}
	if pastPaperId == 0 {
		return "", errors.New("pastPaperId is required")
	}
	if indexInPastPaper == 0 {
		return "", errors.New("indexInPastPaper is required")
	}
	if !ok || questionContentsVal == nil {
		return "", errors.New("questionContents is required")
	}

	contents, err := parseQuestionContents(questionContentsVal)
	if err != nil {
		return "", err
	}

	questionID, err := service.QuestionSvr.CreateQuestion(model.Question{
		SyllabusId:       syllabusId,
		Stem:             stem,
		Difficult:        difficult,
		PastPaperId:      pastPaperId,
		IndexInPastPaper: indexInPastPaper,
		Status:           model.QUESTION_STATE_NORMAL,
		QuestionContents: contents,
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

	// Debug logging
	fmt.Printf("DEBUG toolQuestionEdit args: %+v\n", args)
	if val, ok := args["pastPaperId"]; ok {
		fmt.Printf("DEBUG pastPaperId type: %T, value: %v\n", val, val)
	}

	id := getUint(args, "id", 0)
	if id == 0 {
		return "", errors.New("id is required")
	}

	// All parameters are required
	syllabusId := getUint(args, "syllabusId", 0)
	stem := getString(args, "stem", "")
	difficult := getInt(args, "difficult", 1)
	pastPaperId := getUint(args, "pastPaperId", 0)
	indexInPastPaper := getInt(args, "indexInPastPaper", 0)
	questionContentsVal, ok := args["questionContents"]

	if syllabusId == 0 {
		return "", errors.New("syllabusId is required")
	}
	if stem == "" {
		return "", errors.New("stem is required")
	}
	if pastPaperId == 0 {
		return "", errors.New("pastPaperId is required")
	}
	if indexInPastPaper == 0 {
		return "", errors.New("indexInPastPaper is required")
	}
	if !ok || questionContentsVal == nil {
		return "", errors.New("questionContents is required")
	}

	contents, err := parseQuestionContents(questionContentsVal)
	if err != nil {
		return "", err
	}

	// Get existing question
	existing, err := service.QuestionSvr.SelectQuestionById(id)
	if err != nil {
		return "", err
	}

	// Update all fields
	existing.SyllabusId = syllabusId
	existing.Stem = stem
	existing.Difficult = difficult
	existing.PastPaperId = pastPaperId
	existing.IndexInPastPaper = indexInPastPaper
	existing.QuestionContents = contents

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
