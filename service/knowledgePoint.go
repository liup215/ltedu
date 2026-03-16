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
	"go.uber.org/zap"
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
		newKp.ConfidenceScore = kp.ConfidenceScore
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

// PredictRelevantChapters 预测题目相关的章节
func (s *KnowledgePointService) PredictRelevantChapters(questionStem string, syllabusId uint) ([]uint, error) {
	// 获取考纲的所有章节
	chapters, err := repository.ChapterRepo.FindBySyllabusID(syllabusId)
	if err != nil {
		return nil, fmt.Errorf("failed to get chapters: %w", err)
	}

	if len(chapters) == 0 {
		return nil, errors.New("no chapters found for syllabus")
	}

	// 安全检查：如果章节数量过多，使用分页或限制策略
	const maxChaptersForPrediction = 100
	if len(chapters) > maxChaptersForPrediction {
		logger.Logger.Warn("Syllabus has too many chapters for prediction. Using first N chapters.",
			zap.Uint("syllabusId", syllabusId),
			zap.Int("chapterCount", len(chapters)),
			zap.Int("maxChaptersForPrediction", maxChaptersForPrediction))
		chapters = chapters[:maxChaptersForPrediction]
	}

	// 构建章节列表字符串（比知识点列表小得多）
	chapterList := ""
	for i, ch := range chapters {
		chapterList += fmt.Sprintf("%d. %s\n", i+1, ch.Name)
	}

	// 调用AI服务预测相关章节
	indices, err := s.analyzeQuestionForChapters(questionStem, chapterList)
	if err != nil {
		return nil, err
	}

	// 转换为章节ID
	var relevantChapterIds []uint
	for _, idx := range indices {
		if idx > 0 && idx <= len(chapters) {
			relevantChapterIds = append(relevantChapterIds, chapters[idx-1].ID)
		}
	}

	return relevantChapterIds, nil
}

// AutoLinkQuestionToKeypointsIntelligent 智能关联题目到知识点（两阶段方法）
func (s *KnowledgePointService) AutoLinkQuestionToKeypointsIntelligent(questionId, syllabusId uint) ([]uint, error) {
	question, err := repository.QuestionRepo.FindByID(questionId)
	if err != nil || question == nil {
		return nil, fmt.Errorf("question not found: %w", err)
	}

	if syllabusId == 0 {
		return nil, errors.New("syllabusId is required")
	}

	// 阶段1: 预测相关章节
	relevantChapterIds, err := s.PredictRelevantChapters(question.Stem, syllabusId)
	if err != nil {
		logger.Logger.Warn("Failed to predict relevant chapters for question. Falling back to all chapters.",
			zap.Uint("questionId", questionId),
			zap.Error(err))
		// 回退到所有章节
		chapters, _ := repository.ChapterRepo.FindBySyllabusID(syllabusId)
		for _, ch := range chapters {
			relevantChapterIds = append(relevantChapterIds, ch.ID)
		}
	}

	if len(relevantChapterIds) == 0 {
		// 如果仍然没有章节，回退到所有章节
		chapters, _ := repository.ChapterRepo.FindBySyllabusID(syllabusId)
		for _, ch := range chapters {
			relevantChapterIds = append(relevantChapterIds, ch.ID)
		}
	}

	// 阶段2: 仅从相关章节加载知识点
	var relevantKeypoints []model.KnowledgePoint
	for _, chapterId := range relevantChapterIds {
		kps, _ := repository.KnowledgePointRepo.FindByChapterId(chapterId)
		relevantKeypoints = append(relevantKeypoints, kps...)
	}

	if len(relevantKeypoints) == 0 {
		return nil, errors.New("no knowledge points available for linking")
	}

	// 安全检查：如果相关知识点数量过多，使用限制策略
	const maxKeypointsForLinking = 50
	if len(relevantKeypoints) > maxKeypointsForLinking {
		logger.Logger.Warn("Question has too many relevant knowledge points for linking. Using first N knowledge points.",
			zap.Uint("questionId", questionId),
			zap.Int("knowledgePointCount", len(relevantKeypoints)),
			zap.Int("maxKeypointsForLinking", maxKeypointsForLinking))
		relevantKeypoints = relevantKeypoints[:maxKeypointsForLinking]
	}

	// 构建知识点列表字符串（现在小得多）
	kpList := ""
	for i, kp := range relevantKeypoints {
		kpList += fmt.Sprintf("%d. [%s] %s - %s\n",
			i+1, kp.Chapter.Name, kp.Name, kp.Description)
	}

	// 调用AI服务分析题目（使用减少的上下文）
	indices, err := s.analyzeQuestionForKnowledgePoints(question.Stem, kpList)
	if err != nil {
		return nil, err
	}

	// 建立关联
	var linkedIds []uint
	for _, idx := range indices {
		if idx > 0 && idx <= len(relevantKeypoints) {
			kp := relevantKeypoints[idx-1]
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
	return s.AutoMigrateSyllabusWithProgress(syllabusId, options, nil, nil, nil)
}

// AutoMigrateSyllabusWithProgress 批量自动化处理考纲（支持进度回调和断点续传）
// skipChapterIds: set of chapter IDs already successfully processed (skip them on resume)
// onProgress:    called after each item with (done, total)
// onChapterDone: called with a chapter ID after it is successfully processed (for resume tracking)
func (s *KnowledgePointService) AutoMigrateSyllabusWithProgress(
	syllabusId uint,
	options MigrateOptions,
	skipChapterIds map[uint]bool,
	onProgress func(done, total int),
	onChapterDone func(chapterId uint),
) (*MigrateReport, error) {
	report := &MigrateReport{}

	if options.BatchSize == 0 {
		options.BatchSize = 50
	}

	// Step 1: 获取所有章节
	chapters, err := repository.ChapterRepo.FindBySyllabusID(syllabusId)
	if err != nil {
		return report, fmt.Errorf("failed to get chapters: %w", err)
	}

	// Compute leaf chapters first for progress tracking
	var leafChapters []*model.Chapter
	for _, chapter := range chapters {
		hasChildren, err := repository.ChapterRepo.HasChildren(chapter.ID)
		if err != nil {
			report.Errors = append(report.Errors,
				fmt.Sprintf("Chapter %d (%s): failed to check children: %v", chapter.ID, chapter.Name, err))
			continue
		}
		if hasChildren {
			continue
		}
		leafChapters = append(leafChapters, chapter)
	}

	// Compute question count for progress tracking
	var questionCount int
	if options.LinkQuestions {
		query := &model.QuestionQueryRequest{SyllabusId: syllabusId}
		questions, err := repository.QuestionRepo.FindAll(query)
		if err == nil {
			questionCount = len(questions)
		}
	}

	total := 0
	if options.GenerateKeypoints {
		total += len(leafChapters)
	}
	if options.LinkQuestions {
		total += questionCount
	}

	// Count already-processed chapters so progress starts at the right offset
	alreadyDone := 0
	if options.GenerateKeypoints && skipChapterIds != nil {
		for _, ch := range leafChapters {
			if skipChapterIds[ch.ID] {
				alreadyDone++
			}
		}
	}
	done := alreadyDone

	if onProgress != nil {
		onProgress(done, total)
	}

	// Step 2: 仅为叶子章节生成知识点（没有子章节的章节）
	if options.GenerateKeypoints {
		for _, chapter := range leafChapters {
			// 跳过上次已成功处理的章节（断点续传）
			if skipChapterIds != nil && skipChapterIds[chapter.ID] {
				continue
			}

			// 先删除该章节已有的知识点，避免重复数据（幂等处理）
			if err := repository.KnowledgePointRepo.DeleteByChapterId(chapter.ID); err != nil {
				report.Errors = append(report.Errors,
					fmt.Sprintf("Chapter %d (%s): failed to clear existing keypoints, skipping: %v", chapter.ID, chapter.Name, err))
				done++
				if onProgress != nil {
					onProgress(done, total)
				}
				continue
			}

			kps, err := s.AutoGenerateFromChapter(chapter.ID)
			if err == nil {
				report.GeneratedKeypoints += len(kps)
				// 通知调用方该章节已成功处理（用于断点续传记录）
				if onChapterDone != nil {
					onChapterDone(chapter.ID)
				}
			} else {
				report.Errors = append(report.Errors,
					fmt.Sprintf("Chapter %d (%s): %v", chapter.ID, chapter.Name, err))
			}
			done++
			if onProgress != nil {
				onProgress(done, total)
			}
			time.Sleep(1 * time.Second) // 防止AI接口限流
		}
	}

	// Step 3: 关联题目 - 使用智能两阶段方法
	// 策略：先预测相关章节，再从相关章节的知识点中选择
	if options.LinkQuestions {
		query := &model.QuestionQueryRequest{SyllabusId: syllabusId}
		questions, err := repository.QuestionRepo.FindAll(query)
		if err != nil {
			report.Errors = append(report.Errors,
				fmt.Sprintf("Failed to get questions: %v", err))
			return report, err
		}

		for i, question := range questions {
			// 使用智能方法，避免加载所有知识点
			linkedIds, err := s.AutoLinkQuestionToKeypointsIntelligent(question.ID, syllabusId)
			if err == nil {
				report.LinkedQuestions++
				report.TotalLinks += len(linkedIds)
			} else {
				report.Errors = append(report.Errors,
					fmt.Sprintf("Question %d: %v", question.ID, err))
			}

			done++
			if onProgress != nil {
				onProgress(done, total)
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

// LinkQuestion 手动关联题目到知识点
func (s *KnowledgePointService) LinkQuestion(knowledgePointId uint, questionId uint) error {
	return repository.KnowledgePointRepo.LinkQuestion(knowledgePointId, questionId)
}

// UnlinkQuestion 取消题目与知识点的关联
func (s *KnowledgePointService) UnlinkQuestion(knowledgePointId uint, questionId uint) error {
	return repository.KnowledgePointRepo.UnlinkQuestion(knowledgePointId, questionId)
}

// generateKnowledgePoints AI生成知识点
func (s *KnowledgePointService) generateKnowledgePoints(syllabusName, chapterName string) ([]model.AIKnowledgePointData, error) {
	contextInfo := fmt.Sprintf("考纲: %s, 章节: %s", syllabusName, chapterName)

	prompt := fmt.Sprintf(`
你是考纲专家。请为"%s"提取核心知识点。

要求：
1. 知识点要具体明确，不要过于宽泛
2. 覆盖该章节的主要考点
3. 按重要性排序
4. 使用英文回答

返回严格的JSON数组格式，无其他文字：
[{
    "name": "知识点名称",
    "description": "1-2句话描述该知识点的核心内容",
    "difficulty": "basic/medium/hard",
    "estimatedMinutes": 30,
    "confidenceScore": 0.95
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

// analyzeQuestionForChapters AI分析题目并预测相关章节
func (s *KnowledgePointService) analyzeQuestionForChapters(questionStem string, chapterList string) ([]int, error) {
	prompt := fmt.Sprintf(`
你是教育专家。请分析以下题目，判断它属于哪些章节。

题目内容：
%s

可用章节列表：
%s

要求：
1. 仅选择与题目直接相关的章节
2. 可以选择多个章节（如果题目是综合题）
3. 如果不确定，宁可不选

返回JSON格式（仅包含序号数组，从1开始）：
{"chapterIndices": [1, 3]}
`, questionStem, chapterList)

	aiResponse, err := s.aiModel.CreateCompletion(prompt)
	if err != nil {
		return nil, fmt.Errorf("AI chapter analysis failed: %w", err)
	}

	// 解析AI响应
	var result struct {
		ChapterIndices []int `json:"chapterIndices"`
	}
	err = json.Unmarshal([]byte(aiResponse), &result)
	if err != nil {
		// 尝试备用解析格式
		var altResult struct {
			Indices []int `json:"indices"`
		}
		err2 := json.Unmarshal([]byte(aiResponse), &altResult)
		if err2 != nil {
			return nil, fmt.Errorf("failed to parse AI chapter response: %w", err)
		}
		return altResult.Indices, nil
	}

	return result.ChapterIndices, nil
}
