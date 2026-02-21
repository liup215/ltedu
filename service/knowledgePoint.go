package service

import (
	"edu/conf"
	"edu/lib/ai"
	"edu/lib/logger"
	"edu/model"
	"edu/repository"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

var KnowledgePointSvr = &KnowledgePointService{
	baseService: newBaseService(),
	aiModel:     ai.NewModel(conf.Conf.AiConfig, logger.Logger),
}

type KnowledgePointService struct {
	baseService
	aiModel ai.Model
}

// AutoGenerateFromChapter AI生成章节知识点（仅支持叶子节点）
func (s *KnowledgePointService) AutoGenerateFromChapter(chapterId uint) ([]model.KnowledgePoint, error) {
	chapter, err := repository.ChapterRepo.FindByID(chapterId)
	if err != nil {
		return nil, fmt.Errorf("chapter not found: %w", err)
	}

	// 检查是否为叶子节点（没有子章节）
	hasChildren, err := repository.ChapterRepo.HasChildren(chapterId)
	if err != nil {
		return nil, fmt.Errorf("failed to check children: %w", err)
	}
	if hasChildren {
		return nil, errors.New("只能为叶子章节生成知识点！该章节存在子章节。")
	}

	// 获取考纲信息
	syllabus, err := repository.SyllabusRepo.FindByID(chapter.SyllabusId)
	if err != nil {
		return nil, fmt.Errorf("syllabus not found: %w", err)
	}

	// 调用AI服务生成知识点数据
	kpData, err := s.generateKnowledgePoints(syllabus.Name, chapter.Name)
	if err != nil {
		return nil, err
	}

	var keypoints []model.KnowledgePoint
	for i, kp := range kpData {
		newKp := model.KnowledgePoint{
			ChapterId:        chapterId,
			Name:             kp.Name,
			Description:      kp.Description,
			Difficulty:       kp.Difficulty,
			EstimatedMinutes: kp.EstimatedMinutes,
			OrderIndex:       i + 1,
		}
		err = repository.KnowledgePointRepo.Create(&newKp)
		if err != nil {
			return nil, fmt.Errorf("failed to create knowledge point: %w", err)
		}
		keypoints = append(keypoints, newKp)
	}

	return keypoints, nil
}

// AutoLinkQuestionToKeypoints AI自动关联题目到知识点
func (s *KnowledgePointService) AutoLinkQuestionToKeypoints(questionId, chapterId, syllabusId uint) ([]uint, error) {
	question, err := repository.QuestionRepo.FindByID(questionId)
	if err != nil || question == nil {
		return nil, fmt.Errorf("question not found: %w", err)
	}

	// 获取候选知识点 - 必须明确指定范围
	var keypoints []model.KnowledgePoint

	if chapterId != 0 {
		// 方式1: 从指定章节获取知识点
		keypoints, err = repository.KnowledgePointRepo.FindByChapterId(chapterId)
		if err != nil {
			return nil, fmt.Errorf("failed to get knowledge points: %w", err)
		}
	} else if syllabusId != 0 {
		// 方式2: 从考纲获取所有章节的知识点
		chapters, err := repository.ChapterRepo.FindBySyllabusID(syllabusId)
		if err != nil {
			return nil, fmt.Errorf("failed to get chapters: %w", err)
		}
		for _, ch := range chapters {
			kps, err := repository.KnowledgePointRepo.FindByChapterId(ch.ID)
			if err == nil {
				keypoints = append(keypoints, kps...)
			}
		}
	} else {
		return nil, errors.New("chapterId or syllabusId is required")
	}

	if len(keypoints) == 0 {
		return nil, errors.New("no knowledge points available for linking")
	}

	// 构建知识点列表字符串
	kpList := ""
	for i, kp := range keypoints {
		kpList += fmt.Sprintf("%d. [%s] %s - %s\n",
			i+1, kp.Chapter.Name, kp.Name, kp.Description)
	}

	// 调用AI服务分析题目
	indices, err := s.analyzeQuestionForKnowledgePoints(question.Stem, kpList)
	if err != nil {
		return nil, err
	}

	// 建立关联
	var linkedIds []uint
	for _, idx := range indices {
		if idx > 0 && idx <= len(keypoints) {
			kp := keypoints[idx-1]
			// 添加关联关系
			err = repository.QuestionRepo.AddKnowledgePoint(questionId, kp.ID)
			if err == nil {
				linkedIds = append(linkedIds, kp.ID)
			}
		}
	}

	return linkedIds, nil
}

// MigrateOptions 迁移选项
type MigrateOptions struct {
	GenerateKeypoints bool `json:"generateKeypoints"`
	LinkQuestions     bool `json:"linkQuestions"`
	BatchSize         int  `json:"batchSize"`
}

// MigrateReport 迁移报告
type MigrateReport struct {
	GeneratedKeypoints int      `json:"generatedKeypoints"`
	LinkedQuestions    int      `json:"linkedQuestions"`
	TotalLinks         int      `json:"totalLinks"`
	Errors             []string `json:"errors"`
}

// AutoMigrateSyllabus 批量自动化处理考纲
func (s *KnowledgePointService) AutoMigrateSyllabus(syllabusId uint, options MigrateOptions) (*MigrateReport, error) {
	report := &MigrateReport{}

	if options.BatchSize == 0 {
		options.BatchSize = 50
	}

	// Step 1: 获取所有章节
	chapters, err := repository.ChapterRepo.FindBySyllabusID(syllabusId)
	if err != nil {
		return report, fmt.Errorf("failed to get chapters: %w", err)
	}

	// Step 2: 仅为叶子章节生成知识点（没有子章节的章节）
	if options.GenerateKeypoints {
		for _, chapter := range chapters {
			// 检查是否为叶子节点（没有子章节）
			hasChildren, err := repository.ChapterRepo.HasChildren(chapter.ID)
			if err != nil {
				report.Errors = append(report.Errors,
					fmt.Sprintf("Chapter %d (%s): failed to check children: %v", chapter.ID, chapter.Name, err))
				continue
			}

			// 跳过非叶子节点
			if hasChildren {
				continue
			}

			// 为叶子节点生成知识点
			kps, err := s.AutoGenerateFromChapter(chapter.ID)
			if err == nil {
				report.GeneratedKeypoints += len(kps)
			} else {
				report.Errors = append(report.Errors,
					fmt.Sprintf("Chapter %d (%s): %v", chapter.ID, chapter.Name, err))
			}
			time.Sleep(1 * time.Second) // 防止AI接口限流
		}
	}

	// Step 3: 关联题目
	// 策略：使用考纲级别的知识点池，让AI从所有知识点中选择
	if options.LinkQuestions {
		query := &model.QuestionQueryRequest{SyllabusId: syllabusId}
		questions, err := repository.QuestionRepo.FindAll(query)
		if err != nil {
			report.Errors = append(report.Errors,
				fmt.Sprintf("Failed to get questions: %v", err))
			return report, err
		}

		for i, question := range questions {
			// 使用考纲ID，让AI从该考纲的所有知识点中选择
			linkedIds, err := s.AutoLinkQuestionToKeypoints(question.ID, 0, syllabusId)
			if err == nil {
				report.LinkedQuestions++
				report.TotalLinks += len(linkedIds)
			} else {
				report.Errors = append(report.Errors,
					fmt.Sprintf("Question %d: %v", question.ID, err))
			}

			// 批量处理时添加延迟
			if (i+1)%options.BatchSize == 0 {
				time.Sleep(2 * time.Second)
			}
		}
	}

	return report, nil
}

// Create 创建知识点
func (s *KnowledgePointService) Create(kp *model.KnowledgePoint) error {
	return repository.KnowledgePointRepo.Create(kp)
}

// Update 更新知识点
func (s *KnowledgePointService) Update(kp *model.KnowledgePoint) error {
	return repository.KnowledgePointRepo.Update(kp)
}

// Delete 删除知识点
func (s *KnowledgePointService) Delete(id uint) error {
	return repository.KnowledgePointRepo.Delete(id)
}

// GetByID 根据ID获取知识点
func (s *KnowledgePointService) GetByID(id uint) (*model.KnowledgePoint, error) {
	return repository.KnowledgePointRepo.FindByID(id)
}

// GetByChapterId 根据章节ID获取知识点列表
func (s *KnowledgePointService) GetByChapterId(chapterId uint) ([]model.KnowledgePoint, error) {
	return repository.KnowledgePointRepo.FindByChapterId(chapterId)
}

// GetBySyllabusId 根据考纲ID获取知识点列表
func (s *KnowledgePointService) GetBySyllabusId(syllabusId uint) ([]model.KnowledgePoint, error) {
	return repository.KnowledgePointRepo.FindBySyllabusId(syllabusId)
}

// GetAll 获取知识点列表（带分页）
func (s *KnowledgePointService) GetAll(query *model.KnowledgePointQuery) ([]model.KnowledgePoint, int64, error) {
	return repository.KnowledgePointRepo.FindAll(query)
}

// generateKnowledgePoints AI生成知识点
func (s *KnowledgePointService) generateKnowledgePoints(syllabusName, chapterName string) ([]model.AIKnowledgePointData, error) {
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

	aiResponse, err := s.aiModel.CreateCompletion(prompt)
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

// analyzeQuestionForKnowledgePoints AI分析题目并推荐知识点
func (s *KnowledgePointService) analyzeQuestionForKnowledgePoints(questionStem string, knowledgePointList string) ([]int, error) {
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

	aiResponse, err := s.aiModel.CreateCompletion(prompt)
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
