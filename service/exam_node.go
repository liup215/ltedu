package service

import (
	"edu/model"
	"edu/repository"
	"errors"
)

var ExamNodeSvr = &ExamNodeService{baseService: newBaseService()}

type ExamNodeService struct {
	baseService
}

// CreateExamNode 创建考试节点
func (svr *ExamNodeService) CreateExamNode(req model.SyllabusExamNodeCreateRequest) (*model.SyllabusExamNode, error) {
	syl, err := repository.SyllabusRepo.FindByID(req.SyllabusId)
	if err != nil || syl == nil {
		return nil, errors.New("考纲不存在")
	}
	node := &model.SyllabusExamNode{
		SyllabusId:  req.SyllabusId,
		Name:        req.Name,
		Description: req.Description,
		SortOrder:   req.SortOrder,
	}
	if err := repository.ExamNodeRepo.Create(node); err != nil {
		return nil, err
	}
	return node, nil
}

// UpdateExamNode 更新考试节点基本信息
func (svr *ExamNodeService) UpdateExamNode(req model.SyllabusExamNodeUpdateRequest) (*model.SyllabusExamNode, error) {
	if req.ID == 0 {
		return nil, errors.New("无效的ID")
	}
	node := &model.SyllabusExamNode{
		Model:       model.Model{ID: req.ID},
		Name:        req.Name,
		Description: req.Description,
		SortOrder:   req.SortOrder,
	}
	if err := repository.ExamNodeRepo.Update(node); err != nil {
		return nil, err
	}
	return repository.ExamNodeRepo.FindByID(req.ID)
}

// DeleteExamNode 删除考试节点
func (svr *ExamNodeService) DeleteExamNode(id uint) error {
	if id == 0 {
		return errors.New("无效的ID")
	}
	return repository.ExamNodeRepo.Delete(id)
}

// GetExamNode 根据ID获取考试节点
func (svr *ExamNodeService) GetExamNode(id uint) (*model.SyllabusExamNode, error) {
	if id == 0 {
		return nil, errors.New("无效的ID")
	}
	return repository.ExamNodeRepo.FindByID(id)
}

// ListExamNodes 获取某个Syllabus下的所有考试节点
func (svr *ExamNodeService) ListExamNodes(syllabusId uint) ([]*model.SyllabusExamNode, error) {
	if syllabusId == 0 {
		return nil, errors.New("syllabusId不能为空")
	}
	return repository.ExamNodeRepo.FindBySyllabusID(syllabusId)
}

// AddChapterToExamNode 为考试节点添加章节，递归包含所有子章节
func (svr *ExamNodeService) AddChapterToExamNode(req model.SyllabusExamNodeAddChapterRequest) error {
	if req.ExamNodeId == 0 || req.ChapterId == 0 {
		return errors.New("examNodeId和chapterId不能为空")
	}
	node, err := repository.ExamNodeRepo.FindByID(req.ExamNodeId)
	if err != nil || node == nil {
		return errors.New("考试节点不存在")
	}

	// Collect the root chapter and all its descendants within the same syllabus
	chapterIds, err := svr.collectAllDescendants(node.SyllabusId, req.ChapterId)
	if err != nil {
		return err
	}
	return repository.ExamNodeRepo.AddChapters(req.ExamNodeId, chapterIds)
}

// RemoveChapterFromExamNode 从考试节点移除指定章节
func (svr *ExamNodeService) RemoveChapterFromExamNode(req model.SyllabusExamNodeRemoveChapterRequest) error {
	if req.ExamNodeId == 0 || req.ChapterId == 0 {
		return errors.New("examNodeId和chapterId不能为空")
	}
	return repository.ExamNodeRepo.RemoveChapter(req.ExamNodeId, req.ChapterId)
}

// AddPaperCodeToExamNode 为考试节点添加试卷代码
func (svr *ExamNodeService) AddPaperCodeToExamNode(req model.SyllabusExamNodeAddPaperCodeRequest) error {
	if req.ExamNodeId == 0 || req.PaperCodeId == 0 {
		return errors.New("examNodeId和paperCodeId不能为空")
	}
	node, err := repository.ExamNodeRepo.FindByID(req.ExamNodeId)
	if err != nil || node == nil {
		return errors.New("考试节点不存在")
	}
	return repository.ExamNodeRepo.AddPaperCode(req.ExamNodeId, req.PaperCodeId)
}

// RemovePaperCodeFromExamNode 从考试节点移除试卷代码
func (svr *ExamNodeService) RemovePaperCodeFromExamNode(req model.SyllabusExamNodeRemovePaperCodeRequest) error {
	if req.ExamNodeId == 0 || req.PaperCodeId == 0 {
		return errors.New("examNodeId和paperCodeId不能为空")
	}
	return repository.ExamNodeRepo.RemovePaperCode(req.ExamNodeId, req.PaperCodeId)
}

// collectAllDescendants 收集指定章节及其所有后代章节的ID（在同一Syllabus范围内）
func (svr *ExamNodeService) collectAllDescendants(syllabusId uint, rootChapterId uint) ([]uint, error) {
	allChapters, err := repository.ChapterRepo.FindBySyllabusID(syllabusId)
	if err != nil {
		return nil, err
	}

	// Validate rootChapterId belongs to this syllabus and build child index
	chapterSet := make(map[uint]struct{}, len(allChapters))
	children := make(map[uint][]uint, len(allChapters))
	for _, ch := range allChapters {
		chapterSet[ch.ID] = struct{}{}
		children[ch.ParentId] = append(children[ch.ParentId], ch.ID)
	}
	if _, ok := chapterSet[rootChapterId]; !ok {
		return nil, errors.New("章节不存在或不属于该考纲")
	}

	// BFS using an index to avoid O(n²) slice reslicing
	queue := []uint{rootChapterId}
	var result []uint
	for i := 0; i < len(queue); i++ {
		current := queue[i]
		result = append(result, current)
		queue = append(queue, children[current]...)
	}
	return result, nil
}
