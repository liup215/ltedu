package service

import (
	"encoding/json"
	"edu/model"
	"edu/repository"
	"errors"
	"fmt"
	"math"
	"strconv"
	"time"
)

var LearningPlanSvr = &LearningPlanService{baseService: newBaseService()}

type LearningPlanService struct {
	baseService
}

// CreatePlan 为学生创建学习计划并记录初始版本
func (svr *LearningPlanService) CreatePlan(req model.StudentLearningPlanCreateRequest, creatorId uint) (*model.StudentLearningPlan, error) {
	if req.ClassId == 0 || req.UserId == 0 {
		return nil, errors.New("班级ID和用户ID不能为空")
	}
	if req.PlanType != model.LearningPlanTypeLong &&
		req.PlanType != model.LearningPlanTypeMid &&
		req.PlanType != model.LearningPlanTypeShort {
		return nil, errors.New("无效的计划类型，必须为 long/mid/short")
	}

	// 验证班级存在
	class, err := repository.ClassRepo.FindByID(req.ClassId)
	if err != nil || class == nil {
		return nil, errors.New("班级不存在")
	}
	if class.ClassType != model.ClassTypeTeaching {
		return nil, errors.New("只有教学班可以管理学习计划")
	}
	if class.SyllabusId == nil {
		return nil, errors.New("该教学班尚未绑定syllabus，请先绑定syllabus")
	}

	// 验证学生存在
	user, err := repository.UserRepo.FindByID(req.UserId)
	if err != nil || user == nil {
		return nil, errors.New("学生不存在")
	}

	plan := &model.StudentLearningPlan{
		ClassId:   req.ClassId,
		UserId:    req.UserId,
		PlanType:  req.PlanType,
		Content:   req.Content,
		Version:   1,
		CreatedBy: creatorId,
	}
	if err := repository.StudentLearningPlanRepo.Create(plan); err != nil {
		return nil, err
	}

	// 记录初始版本
	version := &model.StudentLearningPlanVersion{
		PlanId:    plan.ID,
		Version:   1,
		Content:   req.Content,
		ChangedBy: creatorId,
		Comment:   req.Comment,
	}
	if err := repository.StudentLearningPlanRepo.CreateVersion(version); err != nil {
		return nil, err
	}

	return plan, nil
}

// UpdatePlan 更新学习计划并记录新版本
func (svr *LearningPlanService) UpdatePlan(req model.StudentLearningPlanUpdateRequest, updaterId uint) (*model.StudentLearningPlan, error) {
	if req.ID == 0 {
		return nil, errors.New("无效的ID")
	}
	plan, err := repository.StudentLearningPlanRepo.FindByID(req.ID)
	if err != nil || plan == nil {
		return nil, errors.New("学习计划不存在")
	}

	newVersion := plan.Version + 1
	updated := &model.StudentLearningPlan{
		Model:   model.Model{ID: req.ID},
		Content: req.Content,
		Version: newVersion,
	}
	if err := repository.StudentLearningPlanRepo.Update(updated); err != nil {
		return nil, err
	}

	// 记录新版本
	v := &model.StudentLearningPlanVersion{
		PlanId:    plan.ID,
		Version:   newVersion,
		Content:   req.Content,
		ChangedBy: updaterId,
		Comment:   req.Comment,
	}
	if err := repository.StudentLearningPlanRepo.CreateVersion(v); err != nil {
		return nil, err
	}

	plan.Content = req.Content
	plan.Version = newVersion
	return plan, nil
}

// DeletePlan 删除学习计划
func (svr *LearningPlanService) DeletePlan(id uint) error {
	if id == 0 {
		return errors.New("无效的ID")
	}
	return repository.StudentLearningPlanRepo.Delete(id)
}

// GetPlanById 根据ID获取学习计划
func (svr *LearningPlanService) GetPlanById(id uint) (*model.StudentLearningPlan, error) {
	if id == 0 {
		return nil, errors.New("无效的ID")
	}
	return repository.StudentLearningPlanRepo.FindByID(id)
}

// ListPlans 分页查询学习计划
func (svr *LearningPlanService) ListPlans(q model.StudentLearningPlanQuery) ([]*model.StudentLearningPlan, int64, error) {
	page := q.Page.CheckPage()
	return repository.StudentLearningPlanRepo.FindPage(&q, (page.PageIndex-1)*page.PageSize, page.PageSize)
}

// GetAllPlans 获取全部学习计划（不分页）
func (svr *LearningPlanService) GetAllPlans(q model.StudentLearningPlanQuery) ([]*model.StudentLearningPlan, error) {
	return repository.StudentLearningPlanRepo.FindAll(&q)
}

// ListPlanVersions 获取学习计划的历史版本列表
func (svr *LearningPlanService) ListPlanVersions(q model.StudentLearningPlanVersionQuery) ([]*model.StudentLearningPlanVersion, int64, error) {
	if q.PlanId == 0 {
		return nil, 0, errors.New("planId不能为空")
	}
	page := q.Page.CheckPage()
	return repository.StudentLearningPlanRepo.FindVersionsByPlanId(q.PlanId, (page.PageIndex-1)*page.PageSize, page.PageSize)
}

// RollbackPlan 回滚学习计划到指定版本
func (svr *LearningPlanService) RollbackPlan(req model.StudentLearningPlanRollbackRequest, operatorId uint) (*model.StudentLearningPlan, error) {
	if req.PlanId == 0 || req.Version == 0 {
		return nil, errors.New("planId和version不能为空")
	}
	plan, err := repository.StudentLearningPlanRepo.FindByID(req.PlanId)
	if err != nil || plan == nil {
		return nil, errors.New("学习计划不存在")
	}
	targetVersion, err := repository.StudentLearningPlanRepo.FindVersionByPlanAndVersion(req.PlanId, req.Version)
	if err != nil || targetVersion == nil {
		return nil, errors.New("目标版本不存在")
	}

	newVersion := plan.Version + 1
	comment := req.Comment
	if comment == "" {
		comment = "回滚到版本 " + strconv.Itoa(req.Version)
	}

	updated := &model.StudentLearningPlan{
		Model:   model.Model{ID: req.PlanId},
		Content: targetVersion.Content,
		Version: newVersion,
	}
	if err := repository.StudentLearningPlanRepo.Update(updated); err != nil {
		return nil, err
	}

	// 记录回滚版本
	v := &model.StudentLearningPlanVersion{
		PlanId:    req.PlanId,
		Version:   newVersion,
		Content:   targetVersion.Content,
		ChangedBy: operatorId,
		Comment:   comment,
	}
	if err := repository.StudentLearningPlanRepo.CreateVersion(v); err != nil {
		return nil, err
	}

	plan.Content = targetVersion.Content
	plan.Version = newVersion
	return plan, nil
}

// --- Batch template plan generation ---

// phaseNames and phaseDrillEnabled define the 4 standard learning phases in order.
var phaseNames = []string{"新课学习", "一轮复习", "题型综合复习", "最终冲刺"}
var phaseDrillEnabled = []bool{false, true, true, true}

// longPhaseInfo holds computed phase data used when building mid-term plans.
type longPhaseInfo struct {
	StartMonth   time.Time
	EndMonth     time.Time
	Chapters     []*model.Chapter
}

// midWeekInfo holds computed week data used when building short-term plans.
type midWeekInfo struct {
	StartDate time.Time
	Chapters  []*model.Chapter
}

// parseYearMonth parses a "YYYY-MM" string to the first day of that month (UTC).
func parseYearMonth(s string) (time.Time, error) {
	return time.Parse("2006-01", s)
}

// monthsBetween returns the inclusive count of months between start and end.
func monthsBetween(start, end time.Time) int {
	years := end.Year() - start.Year()
	months := int(end.Month()) - int(start.Month())
	return years*12 + months + 1
}

// collectAllDescendants performs a BFS over allChapters to find all descendants of parentId.
func collectAllDescendants(allChapters []*model.Chapter, parentId uint) []*model.Chapter {
	var result []*model.Chapter
	queue := []uint{parentId}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for _, ch := range allChapters {
			if ch.ParentId == current {
				result = append(result, ch)
				queue = append(queue, ch.ID)
			}
		}
	}
	return result
}

// buildLongPlanContent generates the JSON content for the long-term plan and returns
// phase info for downstream mid-term generation.
func buildLongPlanContent(ratios []int, startTime time.Time, totalMonths int, rootChapters []*model.Chapter) (string, []longPhaseInfo, error) {
	type phaseAlloc struct {
		index      int
		ratio      int
		monthCount int
		chapStart  int
		chapEnd    int
	}

	// Determine active phases and their month allocations.
	var activePhases []phaseAlloc
	for i, r := range ratios {
		if r > 0 {
			mc := int(math.Round(float64(totalMonths) * float64(r) / 100.0))
			if mc < 1 {
				mc = 1
			}
			activePhases = append(activePhases, phaseAlloc{index: i, ratio: r, monthCount: mc})
		}
	}

	if len(activePhases) == 0 {
		return `{"phases":[]}`, nil, nil
	}

	// Clamp total allocated months to totalMonths to avoid phases extending beyond endMonth.
	// Reduce month counts from the last phase backwards until the sum fits.
	allocTotal := 0
	for _, p := range activePhases {
		allocTotal += p.monthCount
	}
	for i := len(activePhases) - 1; i >= 0 && allocTotal > totalMonths; i-- {
		excess := allocTotal - totalMonths
		reduce := activePhases[i].monthCount - 1 // keep at least 1
		if reduce > excess {
			reduce = excess
		}
		activePhases[i].monthCount -= reduce
		allocTotal -= reduce
	}

	// Distribute chapters proportionally across active phases by ratio.
	nChapters := len(rootChapters)
	sumActiveRatios := 0
	for _, p := range activePhases {
		sumActiveRatios += p.ratio
	}
	chapIdx := 0
	for i := range activePhases {
		count := int(math.Round(float64(nChapters) * float64(activePhases[i].ratio) / float64(sumActiveRatios)))
		activePhases[i].chapStart = chapIdx
		end := chapIdx + count
		if end > nChapters {
			end = nChapters
		}
		activePhases[i].chapEnd = end
		chapIdx = end
	}
	// Assign any remainder to the last active phase.
	if chapIdx < nChapters {
		activePhases[len(activePhases)-1].chapEnd = nChapters
	}

	type phaseJSON struct {
		Name         string   `json:"name"`
		StartMonth   string   `json:"startMonth"`
		EndMonth     string   `json:"endMonth"`
		Chapters     []string `json:"chapters"`
		DrillEnabled bool     `json:"drillEnabled"`
	}

	var phasesJSON []phaseJSON
	var longInfos []longPhaseInfo

	curMonth := startTime
	for _, p := range activePhases {
		endMonth := curMonth.AddDate(0, p.monthCount-1, 0)

		chapNames := make([]string, 0, p.chapEnd-p.chapStart)
		chapModels := make([]*model.Chapter, 0, p.chapEnd-p.chapStart)
		for _, ch := range rootChapters[p.chapStart:p.chapEnd] {
			chapNames = append(chapNames, ch.Name)
			chapModels = append(chapModels, ch)
		}

		phasesJSON = append(phasesJSON, phaseJSON{
			Name:         phaseNames[p.index],
			StartMonth:   curMonth.Format("2006-01"),
			EndMonth:     endMonth.Format("2006-01"),
			Chapters:     chapNames,
			DrillEnabled: phaseDrillEnabled[p.index],
		})
		longInfos = append(longInfos, longPhaseInfo{
			StartMonth: curMonth,
			EndMonth:   endMonth,
			Chapters:   chapModels,
		})

		curMonth = endMonth.AddDate(0, 1, 0)
	}

	b, err := json.Marshal(struct {
		Phases []phaseJSON `json:"phases"`
	}{Phases: phasesJSON})
	if err != nil {
		return "", nil, err
	}
	return string(b), longInfos, nil
}

// buildMidPlanContent generates the JSON content for the mid-term plan (8-week schedule)
// using chapters from the first 2 months of the long-term plan.
func buildMidPlanContent(longPhases []longPhaseInfo, startTime time.Time, allChapters []*model.Chapter) (string, []midWeekInfo, error) {
	// Collect root chapters scheduled in the first 2 months of the long-term plan.
	// afterFirst2Months is the exclusive upper bound: phases starting on or after this are excluded.
	afterFirst2Months := startTime.AddDate(0, 2, 0)
	var first2MonthChapters []*model.Chapter
	for _, p := range longPhases {
		if !p.StartMonth.Before(afterFirst2Months) {
			break
		}
		first2MonthChapters = append(first2MonthChapters, p.Chapters...)
	}

	// Get direct children (sub-chapters) of each first-2-month root chapter.
	var subChapters []*model.Chapter
	for _, ch := range first2MonthChapters {
		for _, sc := range allChapters {
			if sc.ParentId == ch.ID {
				subChapters = append(subChapters, sc)
			}
		}
	}

	// Find first Monday on or after startTime.
	weekStart := startTime
	for weekStart.Weekday() != time.Monday {
		weekStart = weekStart.AddDate(0, 0, 1)
	}

	type weekEntry struct {
		Week      int      `json:"week"`
		StartDate string   `json:"startDate"`
		EndDate   string   `json:"endDate"`
		Chapters  []string `json:"chapters"`
	}

	const nWeeks = 8
	// nSubChaps may be 0 if no sub-chapters exist; the loop still produces 8 week entries with
	// empty chapter lists, which is valid (the plan records the schedule skeleton).
	nSubChaps := len(subChapters)
	var weeksJSON []weekEntry
	var midInfos []midWeekInfo

	for w := 0; w < nWeeks; w++ {
		wStart := weekStart.AddDate(0, 0, w*7)
		wEnd := wStart.AddDate(0, 0, 6)

		// Distribute sub-chapters evenly across weeks using integer division.
		start := w * nSubChaps / nWeeks
		end := (w + 1) * nSubChaps / nWeeks

		chapNames := make([]string, 0, end-start)
		chapModels := make([]*model.Chapter, 0, end-start)
		for _, ch := range subChapters[start:end] {
			chapNames = append(chapNames, ch.Name)
			chapModels = append(chapModels, ch)
		}

		weeksJSON = append(weeksJSON, weekEntry{
			Week:      w + 1,
			StartDate: wStart.Format("2006-01-02"),
			EndDate:   wEnd.Format("2006-01-02"),
			Chapters:  chapNames,
		})
		midInfos = append(midInfos, midWeekInfo{
			StartDate: wStart,
			Chapters:  chapModels,
		})
	}

	b, err := json.Marshal(struct {
		Weeks []weekEntry `json:"weeks"`
	}{Weeks: weeksJSON})
	if err != nil {
		return "", nil, err
	}
	return string(b), midInfos, nil
}

// buildShortPlanContent generates the JSON content for the short-term plan (14-day schedule)
// using all descendants of chapters from the first 2 weeks of the mid-term plan.
func buildShortPlanContent(midWeeks []midWeekInfo, allChapters []*model.Chapter) (string, error) {
	// Chapters from weeks 1-2.
	var first2WeekChapters []*model.Chapter
	for i := 0; i < 2 && i < len(midWeeks); i++ {
		first2WeekChapters = append(first2WeekChapters, midWeeks[i].Chapters...)
	}

	// Recursively collect all descendants of each week-1/2 chapter.
	var leafChapters []*model.Chapter
	for _, ch := range first2WeekChapters {
		leafChapters = append(leafChapters, collectAllDescendants(allChapters, ch.ID)...)
	}

	// Determine the start date (day 1 = start of week 1).
	var dayStart time.Time
	if len(midWeeks) > 0 {
		dayStart = midWeeks[0].StartDate
	}

	type dayEntry struct {
		Day      int      `json:"day"`
		Date     string   `json:"date"`
		Chapters []string `json:"chapters"`
	}

	const nDays = 14
	// nChaps may be 0 if no descendants exist; each day entry will simply have an empty chapter
	// list, preserving the 14-day schedule skeleton.
	nChaps := len(leafChapters)
	var daysJSON []dayEntry

	for d := 0; d < nDays; d++ {
		start := d * nChaps / nDays
		end := (d + 1) * nChaps / nDays

		chapNames := make([]string, 0, end-start)
		for _, ch := range leafChapters[start:end] {
			chapNames = append(chapNames, ch.Name)
		}

		daysJSON = append(daysJSON, dayEntry{
			Day:      d + 1,
			Date:     dayStart.AddDate(0, 0, d).Format("2006-01-02"),
			Chapters: chapNames,
		})
	}

	b, err := json.Marshal(struct {
		Days []dayEntry `json:"days"`
	}{Days: daysJSON})
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// GenerateTemplatePlans 批量为班级所有学生生成模板学习计划（长期/中期/短期）。
// 每个学生已存在对应类型计划时跳过，其余错误收集后继续处理。
func (svr *LearningPlanService) GenerateTemplatePlans(req model.GeneratePlansRequest, creatorId uint) (*model.GeneratePlansResult, error) {
	// Validate phase ratios.
	if len(req.PhaseRatios) != 4 {
		return nil, errors.New("phaseRatios 必须包含恰好 4 个元素")
	}
	sum := 0
	for _, r := range req.PhaseRatios {
		if r < 0 {
			return nil, errors.New("phaseRatios 中的值不能为负数")
		}
		sum += r
	}
	if sum > 100 {
		return nil, errors.New("phaseRatios 之和不能超过 100")
	}

	startTime, err := parseYearMonth(req.StartMonth)
	if err != nil {
		return nil, fmt.Errorf("无效的 startMonth: %v", err)
	}
	endTime, err := parseYearMonth(req.EndMonth)
	if err != nil {
		return nil, fmt.Errorf("无效的 endMonth: %v", err)
	}
	if endTime.Before(startTime) {
		return nil, errors.New("endMonth 不能早于 startMonth")
	}

	totalMonths := monthsBetween(startTime, endTime)

	// Fetch all chapters for the syllabus once.
	allChapters, err := repository.ChapterRepo.FindAll(&model.ChapterQuery{SyllabusId: req.SyllabusId})
	if err != nil {
		return nil, fmt.Errorf("获取章节失败: %v", err)
	}

	// Extract root chapters (parentId == 0) in ascending creation order.
	var rootChapters []*model.Chapter
	for _, ch := range allChapters {
		if ch.ParentId == 0 {
			rootChapters = append(rootChapters, ch)
		}
	}
	// FindAll returns DESC order; reverse to get ascending (creation) order.
	for i, j := 0, len(rootChapters)-1; i < j; i, j = i+1, j-1 {
		rootChapters[i], rootChapters[j] = rootChapters[j], rootChapters[i]
	}

	// Build the three plan contents.
	longContent, longPhases, err := buildLongPlanContent(req.PhaseRatios, startTime, totalMonths, rootChapters)
	if err != nil {
		return nil, fmt.Errorf("生成长期计划失败: %v", err)
	}
	midContent, midWeeks, err := buildMidPlanContent(longPhases, startTime, allChapters)
	if err != nil {
		return nil, fmt.Errorf("生成中期计划失败: %v", err)
	}
	shortContent, err := buildShortPlanContent(midWeeks, allChapters)
	if err != nil {
		return nil, fmt.Errorf("生成短期计划失败: %v", err)
	}

	planContents := map[string]string{
		model.LearningPlanTypeLong:  longContent,
		model.LearningPlanTypeMid:   midContent,
		model.LearningPlanTypeShort: shortContent,
	}

	// Get all students in the class.
	students, err := repository.ClassRepo.FindStudents(req.ClassId)
	if err != nil {
		return nil, fmt.Errorf("获取班级学生失败: %v", err)
	}

	result := &model.GeneratePlansResult{
		StudentCount: len(students),
	}

	for _, student := range students {
		for _, planType := range []string{model.LearningPlanTypeLong, model.LearningPlanTypeMid, model.LearningPlanTypeShort} {
			// Skip if a plan of this type already exists for the student in this class.
			existing, err := repository.StudentLearningPlanRepo.FindAll(&model.StudentLearningPlanQuery{
				ClassId:  req.ClassId,
				UserId:   student.ID,
				PlanType: planType,
			})
			if err != nil {
				result.Errors = append(result.Errors, fmt.Sprintf("user %d planType %s: 检查已有计划出错: %v", student.ID, planType, err))
				continue
			}
			if len(existing) > 0 {
				continue
			}

			plan := &model.StudentLearningPlan{
				ClassId:   req.ClassId,
				UserId:    student.ID,
				PlanType:  planType,
				Content:   planContents[planType],
				Version:   1,
				CreatedBy: creatorId,
			}
			if err := repository.StudentLearningPlanRepo.Create(plan); err != nil {
				result.Errors = append(result.Errors, fmt.Sprintf("user %d planType %s: 创建计划出错: %v", student.ID, planType, err))
				continue
			}

			// Record the initial version.
			v := &model.StudentLearningPlanVersion{
				PlanId:    plan.ID,
				Version:   1,
				Content:   planContents[planType],
				ChangedBy: creatorId,
				Comment:   req.Comment,
			}
			if err := repository.StudentLearningPlanRepo.CreateVersion(v); err != nil {
				result.Errors = append(result.Errors, fmt.Sprintf("user %d planType %s: 创建版本记录出错: %v", student.ID, planType, err))
				continue
			}

			result.Count++
		}
	}

	return result, nil
}
