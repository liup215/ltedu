package service

import (
	"edu/conf"
	"edu/lib/ai"
	"edu/lib/logger"
	"edu/model"
	"edu/repository"
	"encoding/json"
	"fmt"
	"strings"
	"unicode"
)

// nluSystemPrompt is the educational domain system prompt used for intent classification
// and entity extraction. It is bilingual (Chinese and English).
const nluSystemPrompt = `You are an expert NLU (Natural Language Understanding) system for an educational platform called LTEdu. 
Your task is to analyze student and teacher queries and extract structured intent and entity information.

Supported intents:
- find_question: User wants to find or search for a question
- check_answer: User wants to check or grade their answer
- explain_concept: User wants a concept or topic explained
- practice: User wants to start a practice session
- get_learning_plan: User wants a learning plan or study schedule
- get_progress: User wants to see their progress or performance
- get_syllabus: User wants syllabus or curriculum information
- general_query: Any other educational query

Supported entity types:
- subject: Academic subject (e.g. Biology, Physics, Chemistry, Mathematics, English, 生物, 物理, 化学, 数学)
- chapter: Specific chapter or section name
- topic: Specific topic within a subject
- question_type: Type of question (multiple choice, short answer, structured, 单选, 简答, 结构化)
- difficulty: Difficulty level (easy, medium, hard, basic, advanced, 简单, 中等, 困难)
- language: Language preference (Chinese, English, 中文, 英文)

Always respond in valid JSON only, with no markdown or extra text.
Format:
{
  "intent": "<intent_name>",
  "confidence": <0.0-1.0>,
  "entities": [{"type": "<entity_type>", "value": "<extracted_value>"}],
  "language": "<zh|en|mixed>",
  "normalized": "<cleaned version of the query>"
}`

// NLUSvr is the singleton NLU service.
var NLUSvr = &NLUService{
	baseService: newBaseService(),
	ai:          ai.NewModel(conf.Conf.AiConfig, logger.Logger),
}

// NLUService handles Natural Language Understanding for educational queries.
type NLUService struct {
	baseService
	ai ai.Model
}

// AnalyzeQuery performs intent classification and entity extraction on a user query.
func (s *NLUService) AnalyzeQuery(req model.NLURequest) (*model.NLUResult, error) {
	messages := []ai.Message{
		{Role: "system", Content: nluSystemPrompt},
		{Role: "user", Content: buildNLUPrompt(req.Query, req.Context)},
	}

	raw, err := s.ai.CreateCompletionWithMessages(messages)
	if err != nil {
		return nil, fmt.Errorf("NLU analysis failed: %w", err)
	}

	result, err := parseNLUResult(raw, req.Query)
	if err != nil {
		// Fall back to language detection + generic intent
		return &model.NLUResult{
			Intent:     model.IntentGeneralQuery,
			Confidence: 0.5,
			Entities:   []model.NLUEntity{},
			Language:   detectLanguage(req.Query),
			Normalized: req.Query,
		}, nil
	}

	return result, nil
}

// SaveFeedback stores user-provided NLU correction feedback for future improvement.
func (s *NLUService) SaveFeedback(userID uint, req model.NLUFeedbackRequest) error {
	feedback := &model.NLUFeedback{
		UserID:          userID,
		Query:           req.Query,
		PredictedIntent: req.PredictedIntent,
		CorrectIntent:   req.CorrectIntent,
		Feedback:        req.Feedback,
	}
	return repository.NLUFeedbackRepo.Create(feedback)
}

// buildNLUPrompt constructs the user message for NLU analysis.
func buildNLUPrompt(query, context string) string {
	if context != "" {
		return fmt.Sprintf("Context: %s\n\nQuery: %s", context, query)
	}
	return fmt.Sprintf("Query: %s", query)
}

// parseNLUResult parses the JSON string returned by the AI model.
// originalQuery is used as a fallback when the model returns an empty normalized field.
func parseNLUResult(raw, originalQuery string) (*model.NLUResult, error) {
	raw = strings.TrimSpace(raw)
	// Strip markdown code fences if present
	if strings.HasPrefix(raw, "```") {
		lines := strings.SplitN(raw, "\n", 2)
		if len(lines) == 2 {
			raw = lines[1]
		}
		raw = strings.TrimSuffix(strings.TrimSpace(raw), "```")
		raw = strings.TrimSpace(raw)
	}

	var result model.NLUResult
	if err := json.Unmarshal([]byte(raw), &result); err != nil {
		return nil, fmt.Errorf("failed to parse NLU response: %w", err)
	}

	// Normalise intent to a known value
	result.Intent = normalizeIntent(result.Intent)
	if result.Language == "" {
		result.Language = detectLanguage(result.Normalized)
	}
	if result.Normalized == "" {
		result.Normalized = originalQuery
	}

	return &result, nil
}

// normalizeIntent maps any unknown intent string to the closest known constant.
func normalizeIntent(intent string) string {
	known := map[string]string{
		model.IntentFindQuestion:   model.IntentFindQuestion,
		model.IntentCheckAnswer:    model.IntentCheckAnswer,
		model.IntentExplainConcept: model.IntentExplainConcept,
		model.IntentPractice:       model.IntentPractice,
		model.IntentLearningPlan:   model.IntentLearningPlan,
		model.IntentGetProgress:    model.IntentGetProgress,
		model.IntentGetSyllabus:    model.IntentGetSyllabus,
		model.IntentGeneralQuery:   model.IntentGeneralQuery,
	}
	if v, ok := known[intent]; ok {
		return v
	}
	return model.IntentGeneralQuery
}

// detectLanguage performs a simple heuristic to detect query language.
func detectLanguage(text string) string {
	chineseCount := 0
	latinCount := 0
	for _, r := range text {
		if unicode.Is(unicode.Han, r) {
			chineseCount++
		} else if r > 0x40 && r < 0x7B { // A-Z / a-z
			latinCount++
		}
	}
	if chineseCount > 0 && latinCount > 0 {
		return model.LanguageMixed
	}
	if chineseCount > 0 {
		return model.LanguageChinese
	}
	return model.LanguageEnglish
}
