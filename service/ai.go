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
