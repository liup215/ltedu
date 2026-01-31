// ltedu-api/service/practice.go

package service

import (
	"edu/model"
	"edu/repository"
	"encoding/json"
	"errors"
	"math/rand"
	"time"
)

var PracticeSvr = &PracticeService{baseService: newBaseService()}

type PracticeService struct {
	baseService
}

func (svr *PracticeService) GenerateQuickPractice(req model.PracticeQuickRequest) ([]uint, error) {
	if req.QuestionCount <= 0 {
		return nil, errors.New("questionCount must be greater than 0")
	}
	if req.SyllabusId == 0 {
		return nil, errors.New("syllabusId is required")
	}

	// Query questions by syllabusId and status=normal
	query := &model.QuestionQueryRequest{
		SyllabusId: req.SyllabusId,
		Status:     model.QUESTION_STATE_NORMAL,
	}
	questions, err := repository.QuestionRepo.FindAll(query)
	if err != nil {
		return nil, err
	}
	if len(questions) == 0 {
		return nil, errors.New("no questions found for this syllabus")
	}

	// Collect question IDs
	ids := make([]uint, len(questions))
	for i, q := range questions {
		ids[i] = q.ID
	}

	// Shuffle and select up to QuestionCount
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(ids), func(i, j int) { ids[i], ids[j] = ids[j], ids[i] })
	if len(ids) > req.QuestionCount {
		ids = ids[:req.QuestionCount]
	}

	return ids, nil
}

func (svr *PracticeService) GeneratePaperPractice(req model.PracticePaperRequest) ([]uint, error) {
	if req.PaperId == 0 {
		return nil, errors.New("paperId is required")
	}

	// Query questions by pastPaperId and status=normal
	query := &model.QuestionQueryRequest{
		PastPaperId: req.PaperId,
		Status:      model.QUESTION_STATE_NORMAL,
	}
	questions, err := repository.QuestionRepo.FindAll(query)
	if err != nil {
		return nil, err
	}
	if len(questions) == 0 {
		return nil, errors.New("no questions found for this paper")
	}

	// Collect question IDs
	ids := make([]uint, len(questions))
	for i, q := range questions {
		ids[i] = q.ID
	}

	return ids, nil
}

// GradePracticeSubmission grades the student's answers and returns detailed results

func (svr *PracticeService) GradePracticeSubmission(sub model.PracticeGradeRequest) (*model.PracticeResult, error) {
	ids := make([]uint, len(sub))
	for i, ans := range sub {
		ids[i] = ans.QuestionID
	}
	questions, err := repository.QuestionRepo.FindByIDs(ids)
	if err != nil {
		return nil, err
	}
	if len(questions) == 0 {
		return nil, errors.New("no questions found")
	}

	score := 0
	total := len(sub)
	results := make([]model.PracticeResultItem, total)

	questionMap := make(map[uint]*model.Question)
	for _, q := range questions {
		q.Format()
		questionMap[q.ID] = q
	}

	for i, ans := range sub {
		q := questionMap[ans.QuestionID]
		if q == nil || len(q.QuestionContents) == 0 {
			results[i] = model.PracticeResultItem{
				QuestionID: ans.QuestionID,
				SubResults: []model.PracticeSubResultItem{},
			}
			continue
		}

		subResults := make([]model.PracticeSubResultItem, len(q.QuestionContents))
		for j, qc := range q.QuestionContents {
			var isCorrect *bool
			var correctAnswer, modelAnswer string
			studentAnswer := ""
			// Find student answer for this QuestionContent
			for _, part := range ans.Answers {
				if part.QuestionContentId == uint(j) {
					studentAnswer = part.Answer
					break
				}
			}

			switch qc.QuestionTypeId {
			case model.QUESTION_TYPE_SINGLE_CHOICE:
				correctAnswer = qc.SingleChoice.Answer
				is := studentAnswer == correctAnswer
				isCorrect = &is
				if is {
					score++
				}
			case model.QUESTION_TYPE_MULTIPLE_CHOICE:
				correctSlice := qc.MultipleChoice.Answer
				studentSlice := jsonStringToSlice(studentAnswer)
				match := slicesEqual(correctSlice, studentSlice)
				isCorrect = &match
				if match {
					score++
				}
				correctAnswer = sliceToString(correctSlice)
			case model.QUESTION_TYPE_TRUE_FALSE:
				correctAnswer = ""
				if qc.TrueOrFalse.Answer == 1 {
					correctAnswer = "true"
				} else if qc.TrueOrFalse.Answer == 2 {
					correctAnswer = "false"
				}
				is := studentAnswer == correctAnswer
				isCorrect = &is
				if is {
					score++
				}
			case model.QUESTION_TYPE_GAP_FILLING:
				correctSlice := qc.GapFilling.Answer
				found := false
				for _, ca := range correctSlice {
					if studentAnswer == ca {
						found = true
						break
					}
				}
				isCorrect = &found
				if found {
					score++
				}
				correctAnswer = sliceToString(correctSlice)
			case model.QUESTION_TYPE_SHORT_ANSWER, model.QUESTION_TYPE_STRUCTURED:
				isCorrect = nil
				modelAnswer = qc.ShortAnswer.Answer
				correctAnswer = ""
			default:
				isCorrect = nil
			}

			subResults[j] = model.PracticeSubResultItem{
				QuestionContentId: uint(j), // use index as unique id
				QuestionType:      int(qc.QuestionTypeId),
				CorrectAnswer:     correctAnswer,
				StudentAnswer:     studentAnswer,
				IsCorrect:         isCorrect,
				ModelAnswer:       modelAnswer,
			}
		}

		results[i] = model.PracticeResultItem{
			QuestionID: ans.QuestionID,
			SubResults: subResults,
		}
	}

	return &model.PracticeResult{
		Score:   score,
		Total:   total,
		Results: results,
	}, nil
}

// Helper functions for answer comparison
func jsonStringToSlice(s string) []string {
	var arr []string
	_ = json.Unmarshal([]byte(s), &arr)
	return arr
}

func sliceToString(arr []string) string {
	b, _ := json.Marshal(arr)
	return string(b)
}

func slicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	m := make(map[string]bool)
	for _, v := range a {
		m[v] = true
	}
	for _, v := range b {
		if !m[v] {
			return false
		}
	}
	return true
}
