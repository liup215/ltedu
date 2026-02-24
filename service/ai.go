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

// GenerateKnowledgePoints AI生成知识点（用于KnowledgePointService）
func (s *AiService) GenerateKnowledgePoints(syllabusName, chapterName string) ([]model.AIKnowledgePointData, error) {
	contextInfo := fmt.Sprintf("考纲: %s, 章节: %s", syllabusName, chapterName)

	prompt := fmt.Sprintf(`
你是考纲专家。请为"%s"提取3-5个核心知识点。

要求：
1. 知识点要具体明确，不要过于宽泛
2. 覆盖该章节的主要考点
3. 按重要性排序

返回严格的JSON数组格式，无其他文字：
[{
    "name": "知识点名称",
    "description": "1-2句话描述该知识点的核心内容",
    "difficulty": "basic/medium/hard",
    "estimatedMinutes": 30
}]
	`, contextInfo)

	aiResponse, err := s.ai.CreateCompletion(prompt)
	if err != nil {
		return nil, fmt.Errorf("AI generation failed: %w", err)
	}

	// 解析AI响应
	var kpData []model.AIKnowledgePointData
	err = json.Unmarshal([]byte(aiResponse), &kpData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse AI response: %w", err)
	}

	return kpData, nil
}

// AnalyzeQuestionForKnowledgePoints AI分析题目并推荐知识点（用于KnowledgePointService）
func (s *AiService) AnalyzeQuestionForKnowledgePoints(questionStem string, knowledgePointList string) ([]int, error) {
	prompt := fmt.Sprintf(`
你是教育专家。请分析以下题目，判断它涉及哪些知识点。

题目内容：
%s

可选知识点列表：
%s

要求：
1. 仅选择与题目直接相关的知识点
2. 可以选择多个知识点（如果题目是综合题）
3. 如果不确定，宁可不选

返回JSON格式（仅包含序号数组，从1开始）：
{"indices": [1, 3]}
	`, questionStem, knowledgePointList)

	aiResponse, err := s.ai.CreateCompletion(prompt)
	if err != nil {
		return nil, fmt.Errorf("AI analysis failed: %w", err)
	}

	// 解析AI响应
	var result struct {
		Indices []int `json:"indices"`
	}
	err = json.Unmarshal([]byte(aiResponse), &result)
	if err != nil {
		return nil, fmt.Errorf("failed to parse AI response: %w", err)
	}

	return result.Indices, nil
}
