package service

import (
	"crypto/md5"
	"edu/lib/storage"
	"edu/model"
	"edu/repository"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

var QuestionSvr = &QuestionService{baseService: newBaseService()}

type QuestionService struct {
	baseService
}

var (
	base64Regex = regexp.MustCompile(`src=["'](data:image/([^;]+);base64,([^"']+))["']`)
)

func (svr *QuestionService) SelectQuestionAllCount(q model.QuestionQueryRequest) (int64, error) {
	return repository.QuestionRepo.Count(&q)
}

func (svr *QuestionService) SelectQuestionAll(q model.QuestionQueryRequest) ([]*model.Question, error) {
	list, err := repository.QuestionRepo.FindAll(&q)
	for _, q := range list {
		q.Format()
	}
	return list, err
}

func (svr *QuestionService) SelectQuestionList(q model.QuestionQueryRequest) ([]*model.Question, int64, error) {
	page := q.Page.CheckPage()
	list, total, err := repository.QuestionRepo.FindPage(&q, (page.PageIndex-1)*page.PageSize, page.PageSize)
	for _, q := range list {
		q.Format()
	}
	return list, total, err
}

func (svr *QuestionService) SelectQuestionById(id uint) (*model.Question, error) {
	if id == 0 {
		return nil, errors.New("ID不能为空")
	}
	q, err := repository.QuestionRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	q.Format()
	return q, nil
}

func (svr *QuestionService) CreateQuestion(q model.Question, uid uint) (uint, error) {
	if uid == 0 {
		return 0, errors.New("用户ID不能为空")
	}
	if q.ID != 0 {
		q.ID = 0
	}
	if q.Status == 0 {
		q.Status = model.QUESTION_STATE_FORBIDDEN
	}
	contentsStringBytes, err := json.Marshal(q.QuestionContents)
	if err != nil {
		return 0, err
	}
	q.QuestionContentsString = string(contentsStringBytes)
	e := repository.QuestionRepo.Create(&q)
	if e != nil {
		return 0, e
	}
	return q.ID, nil
}

func (svr *QuestionService) EditQuestion(q model.Question) error {
	if q.ID == 0 {
		return errors.New("无效的ID")
	}
	if q.Status == 0 {
		q.Status = model.QUESTION_STATE_FORBIDDEN
	}
	contentsStringBytes, err := json.Marshal(q.QuestionContents)
	if err != nil {
		return err
	}
	q.QuestionContentsString = string(contentsStringBytes)
	return repository.QuestionRepo.Update(&q)
}

func (svr *QuestionService) DeleteQuestion(id uint) error {
	if id == 0 {
		return errors.New("无效的ID")
	}
	return repository.QuestionRepo.Delete(id)
}

// 生成练习题
func (svr *QuestionService) GenerateQuestionExercise(query model.QuestionQueryRequest) (list []*model.Question, err error) {
	total, err := repository.QuestionRepo.Count(&query)
	if err != nil {
		return
	}
	if total < 20 {
		list, err = repository.QuestionRepo.FindAll(&query)
		return
	}
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randNums := rand.Perm(int(total))
	randNums = randNums[:20]
	for _, num := range randNums {
		questions, _, err := repository.QuestionRepo.FindPage(&query, num, 1)
		if err != nil || len(questions) == 0 {
			continue
		}
		q := questions[0]
		err = q.Format()
		if err != nil {
			continue
		}
		list = append(list, q)
	}
	return
}

// 兼容 controller 空实现
func (svr *QuestionService) AddQuestionChapter(o interface{}) error {
	return nil
}
func (svr *QuestionService) DeleteQuestionChapter(o interface{}) error {
	return nil
}

func (svr *QuestionService) MigrateBase64Images() error {
	// Initialize Config Service to get storage settings
	cfg, err := ConfigSvr.GetImageUploadConfigRaw()
	if err != nil {
		fmt.Printf("Failed to get image upload config: %v\nUsing local storage.\n", err)
		cfg = model.ImageUploadConfig{Disk: model.LTEDU_CONFIG_IMAGE_UPLOAD_DISK_PUBLIC}
	} else if cfg.Disk == "" {
		cfg.Disk = model.LTEDU_CONFIG_IMAGE_UPLOAD_DISK_PUBLIC
	}

	// Create Storage
	store, err := storage.NewStorage(cfg)
	if err != nil {
		return fmt.Errorf("failed to create storage: %v", err)
	}

	// Get all questions
	// Note: For large datasets, this should be paginated.
	// Given the context, we will fetch all but in a real-world scenario, iterating with limit/offset is better.
	// Since repository methods might not support iteration easily without direct DB access,
	// we'll assume the dataset fits in memory or use a loop with manual pagination if needed.
	// We'll implementation pagination similar to the CLI tool.

	pageIndex := 1
	pageSize := 100

	for {
		list, _, err := repository.QuestionRepo.FindPage(&model.QuestionQueryRequest{}, (pageIndex-1)*pageSize, pageSize)
		if err != nil {
			return fmt.Errorf("failed to fetch questions: %v", err)
		}
		if len(list) == 0 {
			break
		}

		for _, q := range list {
			changed := false

			// Process Stem
			newStem, c := svr.processContent(q.Stem, store)
			if c {
				q.Stem = newStem
				changed = true
			}

			// Process QuestionContents
			if q.QuestionContentsString != "" {
				var qcs []model.QuestionContent
				if err := json.Unmarshal([]byte(q.QuestionContentsString), &qcs); err == nil {
					qcsChanged := false
					for i := range qcs {
						// Process Analyze
						if qcs[i].Analyze != "" {
							newAnalyze, c := svr.processContent(qcs[i].Analyze, store)
							if c {
								qcs[i].Analyze = newAnalyze
								qcsChanged = true
							}
						}
						// Process Answer
						if qcs[i].Answer != "" {
							newAnswer, c := svr.processContent(qcs[i].Answer, store)
							if c {
								qcs[i].Answer = newAnswer
								qcsChanged = true
							}
						}

						// Process Options (Single Choice)
						for j := range qcs[i].SingleChoice.Options {
							newContent, c := svr.processContent(qcs[i].SingleChoice.Options[j].Content, store)
							if c {
								qcs[i].SingleChoice.Options[j].Content = newContent
								qcsChanged = true
							}
						}
						// Process Options (Multiple Choice)
						for j := range qcs[i].MultipleChoice.Options {
							newContent, c := svr.processContent(qcs[i].MultipleChoice.Options[j].Content, store)
							if c {
								qcs[i].MultipleChoice.Options[j].Content = newContent
								qcsChanged = true
							}
						}
					}

					if qcsChanged {
						bs, _ := json.Marshal(qcs)
						q.QuestionContentsString = string(bs)
						changed = true
					}
				}
			}

			if changed {
				if err := repository.QuestionRepo.Update(q); err != nil {
					fmt.Printf("Failed to save question %d: %v\n", q.ID, err)
				}
			}
		}
		pageIndex++
	}
	return nil
}

func (svr *QuestionService) processContent(content string, store storage.Storage) (string, bool) {
	changed := false
	newContent := base64Regex.ReplaceAllStringFunc(content, func(match string) string {
		submatches := base64Regex.FindStringSubmatch(match)
		if len(submatches) != 4 {
			return match
		}
		// submatches[0] is full match
		// submatches[1] is "data:image/ext;base64,content"
		// submatches[2] is ext (e.g. png)
		// submatches[3] is base64 content

		ext := submatches[2]
		b64Data := submatches[3]

		// Decode base64
		data, err := base64.StdEncoding.DecodeString(b64Data)
		if err != nil {
			fmt.Printf("Failed to decode base64: %v\n", err)
			return match
		}

		// Calculate MD5
		md5hash := fmt.Sprintf("%x", md5.Sum(data))

		// Use a simpler valid extension fix if needed (e.g. jpeg -> jpg)
		if ext == "jpeg" {
			ext = "jpg"
		}

		objectName := fmt.Sprintf("uploads/%s/%s.%s", strings.Join(strings.Split(md5hash, "")[0:5], "/"), md5hash, ext)

		// Create temp file
		tmpFile := fmt.Sprintf("temp_%s.%s", md5hash, ext)
		if err := os.WriteFile(tmpFile, data, 0644); err != nil {
			fmt.Printf("Failed to write temp file: %v\n", err)
			return match
		}
		defer os.Remove(tmpFile)

		// Upload
		url, err := store.Upload(objectName, tmpFile)
		if err != nil {
			fmt.Printf("Failed to upload image: %v\n", err)
			return match
		}

		changed = true
		return fmt.Sprintf(`src="%s"`, url)
	})

	return newContent, changed
}
