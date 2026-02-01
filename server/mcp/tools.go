package mcp

import (
	"errors"
	"strconv"
)

// getAvailableTools returns the list of all available MCP tools
func (s *MCPServer) getAvailableTools() []map[string]interface{} {
	var tools []map[string]interface{}

	// Aggregating tools from different domains
	tools = append(tools, s.getOrganisationTools()...)
	tools = append(tools, s.getQualificationTools()...)
	tools = append(tools, s.getSyllabusTools()...)
	tools = append(tools, s.getChapterTools()...)
	tools = append(tools, s.getPaperSeriesTools()...)
	tools = append(tools, s.getPaperCodeTools()...)
	tools = append(tools, s.getPastPaperTools()...)
	tools = append(tools, s.getQuestionTools()...)

	return tools
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
		if strVal, ok := val.(string); ok {
			if floatVal, err := strconv.ParseFloat(strVal, 64); err == nil {
				return floatVal
			}
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
