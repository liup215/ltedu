package service

import (
	"edu/conf"
	"edu/lib/ai"
	"edu/lib/logger"
	"edu/model"
	"encoding/json"
	"fmt"

	"go.uber.org/zap"
)

const (
	AI_PROMPT_VOCABULARY = `
	Your are a experienced bilingual teacher. There is a term to define: "%s", the right answer is "%s". Contents in the parenthesis stands for explainations or alternations and so is not necessary. I give an answer that "%s". Please mark my answer (full mark is 10), and comment on it.
	Returns should be in the json string format, without any other characters, like this:
	{"text": "your text", "score": your score}
	`

	AI_PROMPT_QUESTION = `
	Your are a experienced bilingual teacher. There is a question: "%s", the mark scheme is "%s". The Mark scheme abbreviations is : "%s" Contents in the parenthesis stands for explainations or alternations and so is not necessary. I give an answer that "%s". Evaluate all my statement. Please mark my answer (full mark is %v), and comment on it.
	Returns should be in the following json format with no markdown or html tags, , without any other characters, like this:
	{"text": "comment text to the answer", "score": "integer score"}
	`
)

const (
	MarkSchemeAbbreviations = `
		';' stands for separates marking points.
		'/' stands for alternative answers for the same point.
		'A' stands for accept (for answers correctly cued by the question, or by extra guidance).
		'R' stands for reject (for answers which are not acceptable).
		'I' stands for ignore (for answers which are irrelevant).
		'()' stands for the word / phrase that is not required, but sets the context.
		'AW' stands for alternative wording (where responses vary more than usual).
		'MAX' stands for indicates the maximum number of marks that can be given.
		'ora' stands for or reverse argument.
		'ecf' stands for error carried forward.
		'AVP' stands for alternative valid point.
	`
)

var AiSvr = &AiService{
	baseService: newBaseService(),
	ai:          ai.NewModel(conf.Conf.AiConfig, logger.Logger),
}

type AiService struct {
	baseService
	ai ai.Model
}

func (s *AiService) CheckVocabulary(stem, answer, studentAnswer string) (*model.CheckResponse, error) {
	res := model.CheckResponse{}
	prompt := fmt.Sprintf(AI_PROMPT_VOCABULARY, stem, answer, studentAnswer)
	text, err := s.ai.CreateCompletion(prompt)
	if err != nil {
		return &res, err
	}

	err = json.Unmarshal([]byte(text), &res)
	if err != nil {
		res.Text = text
		logger.Logger.Error("AiService.CheckVocabulary json.Unmarshal error: ", zap.Error(err))
		return &res, nil
	}

	return &res, nil
}

func (s *AiService) CheckQuestion(stem, markScheme, studentAnswer string, fullMark int) (*model.CheckResponse, error) {
	res := model.CheckResponse{}
	prompt := fmt.Sprintf(AI_PROMPT_QUESTION, stem, markScheme, MarkSchemeAbbreviations, studentAnswer, fullMark)
	text, err := s.ai.CreateCompletion(prompt)
	if err != nil {
		return &res, err
	}

	err = json.Unmarshal([]byte(text), &res)
	if err != nil {
		res.Text = text
		logger.Logger.Error("AiService.CheckQuestion json.Unmarshal error: ", zap.Error(err))
		return &res, nil
	}

	return &res, nil
}

// GenerateKnowledgePoints generates knowledge points for a chapter using AI (English output required)
func (s *AiService) GenerateKnowledgePoints(syllabusName, chapterName string) ([]model.AIKnowledgePointData, error) {
	contextInfo := fmt.Sprintf("Syllabus: %s, Chapter: %s", syllabusName, chapterName)

	prompt := fmt.Sprintf(`
You are a curriculum expert. Extract 3-5 core knowledge points for "%s".

Requirements:
1. Each knowledge point must be specific and clearly defined, not overly broad
2. Cover the main examination topics of this chapter
3. Order by importance

Return a strict JSON array with no other text:
[{
    "name": "knowledge point name",
    "description": "1-2 sentences describing the core content of this knowledge point",
    "difficulty": "basic/medium/hard",
    "estimatedMinutes": 30
}]
	`, contextInfo)

	aiResponse, err := s.ai.CreateCompletion(prompt)
	if err != nil {
		return nil, fmt.Errorf("AI generation failed: %w", err)
	}

	// 解析AI响应
	jsonStr := extractJSONFromAIResponse(aiResponse)
	if jsonStr == "" {
		return nil, fmt.Errorf("failed to parse AI response: empty response")
	}
	var kpData []model.AIKnowledgePointData
	err = json.Unmarshal([]byte(jsonStr), &kpData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse AI response: %w", err)
	}

	return kpData, nil
}

// AnalyzeQuestionForKnowledgePoints uses AI to analyze a question and recommend knowledge points (English output required)
func (s *AiService) AnalyzeQuestionForKnowledgePoints(questionStem string, knowledgePointList string) ([]int, error) {
	prompt := fmt.Sprintf(`
You are an education expert. Analyze the following question and determine which knowledge points it covers.

Question:
%s

Available knowledge points:
%s

Requirements:
1. Only select knowledge points directly related to the question
2. Multiple knowledge points may be selected (if the question is comprehensive)
3. When in doubt, do not select

Return JSON format (array of indices only, 1-based):
{"indices": [1, 3]}
	`, questionStem, knowledgePointList)

	aiResponse, err := s.ai.CreateCompletion(prompt)
	if err != nil {
		return nil, fmt.Errorf("AI analysis failed: %w", err)
	}

	// 解析AI响应
	jsonStr := extractJSONFromAIResponse(aiResponse)
	if jsonStr == "" {
		return nil, fmt.Errorf("failed to parse AI response: empty response")
	}
	var result struct {
		Indices []int `json:"indices"`
	}
	err = json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		return nil, fmt.Errorf("failed to parse AI response: %w", err)
	}

	return result.Indices, nil
}
