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
// Each exam node goes through all 4 phases (新课 → 一轮复习 → 专题综合复习 → 集中刷题).
var phaseNames = []string{"新课学习", "一轮复习", "专题综合复习", "集中刷题"}
var phaseDrillEnabled = []bool{false, true, true, true}

const (
	ExamNodeModeSequential = "sequential" // 顺序模式：先完成考试节点A，再进行考试节点B
	ExamNodeModeParallel   = "parallel"   // 并行模式：同时准备所有考试节点
)

// longPhaseInfo holds computed phase data used when building mid-term plans.
// Each entry represents a time range and the set of chapters scheduled in that range.
type longPhaseInfo struct {
	StartMonth time.Time
	EndMonth   time.Time
	Chapters   []*model.Chapter
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

// buildLongPlanContent generates the JSON content for the long-term plan organised
// around exam nodes. Each exam node gets its own complete 4-phase cycle.
//
// Sequential mode: exam nodes are prepared one after another.
//   - Total months are divided evenly among nodes; the last node absorbs the remainder.
//   - Within each node's time slice, the 4 phases are distributed by phaseRatios.
//
// Parallel mode: all exam nodes are prepared simultaneously throughout the full period.
//   - The 4 phases are distributed over the full period by phaseRatios.
//   - Within each phase, all exam nodes contribute their proportional chapter slices.
//
// Returns the JSON string, a flat list of longPhaseInfo (for mid-term derivation), and an error.
func buildLongPlanContent(ratios []int, startTime time.Time, totalMonths int, mode string, examNodes []*model.SyllabusExamNode) (string, []longPhaseInfo, error) {
	if len(examNodes) == 0 {
		return `{"mode":"` + mode + `","examNodes":[]}`, nil, nil
	}

	// Validate / default mode.
	if mode != ExamNodeModeParallel {
		mode = ExamNodeModeSequential
	}

	// Build active phase descriptors (skip zero-ratio phases).
	type activePhase struct {
		index int // 0-3
		ratio int
	}
	var activePhases []activePhase
	for i, r := range ratios {
		if r > 0 {
			activePhases = append(activePhases, activePhase{index: i, ratio: r})
		}
	}
	if len(activePhases) == 0 {
		return `{"mode":"` + mode + `","examNodes":[]}`, nil, nil
	}

	// Helper: distribute nMonths across phases proportionally.
	distributeMonths := func(nMonths int) []int {
		sumRatio := 0
		for _, p := range activePhases {
			sumRatio += p.ratio
		}
		months := make([]int, len(activePhases))
		alloc := 0
		for i, p := range activePhases {
			m := int(math.Round(float64(nMonths) * float64(p.ratio) / float64(sumRatio)))
			if m < 1 {
				m = 1
			}
			months[i] = m
			alloc += m
		}
		// Clamp to nMonths.
		for i := len(activePhases) - 1; i >= 0 && alloc > nMonths; i-- {
			excess := alloc - nMonths
			reduce := months[i] - 1
			if reduce > excess {
				reduce = excess
			}
			months[i] -= reduce
			alloc -= reduce
		}
		return months
	}

	// Helper: distribute nChapters across phases proportionally.
	distributeChapters := func(chapters []*model.Chapter) []int {
		n := len(chapters)
		sumRatio := 0
		for _, p := range activePhases {
			sumRatio += p.ratio
		}
		ends := make([]int, len(activePhases))
		idx := 0
		for i, p := range activePhases {
			count := int(math.Round(float64(n) * float64(p.ratio) / float64(sumRatio)))
			end := idx + count
			if end > n {
				end = n
			}
			ends[i] = end
			idx = end
		}
		// Assign remainder to last phase.
		if idx < n {
			ends[len(ends)-1] = n
		}
		return ends
	}

	var longInfos []longPhaseInfo

	if mode == ExamNodeModeSequential {
		// --- Sequential: each exam node occupies its own consecutive time slice ---
		type nodePhaseJSON struct {
			Name         string   `json:"name"`
			StartMonth   string   `json:"startMonth"`
			EndMonth     string   `json:"endMonth"`
			Chapters     []string `json:"chapters"`
			DrillEnabled bool     `json:"drillEnabled"`
		}
		type nodeJSON struct {
			ID         uint            `json:"id"`
			Name       string          `json:"name"`
			StartMonth string          `json:"startMonth"`
			EndMonth   string          `json:"endMonth"`
			Phases     []nodePhaseJSON `json:"phases"`
		}

		// Divide total months among nodes; last node absorbs remainder.
		baseMonths := totalMonths / len(examNodes)
		if baseMonths < 1 {
			baseMonths = 1
		}

		var nodesJSON []nodeJSON
		curMonth := startTime

		for ni, node := range examNodes {
			nodeMonths := baseMonths
			if ni == len(examNodes)-1 {
				// Last node gets the remainder so all months are consumed.
				used := baseMonths * ni
				remaining := totalMonths - used
				if remaining > 0 {
					nodeMonths = remaining
				}
			}

			// Distribute phases within this node's time slice.
			phaseMonths := distributeMonths(nodeMonths)

			// Distribute node's chapters across phases.
			nodeChapters := node.Chapters
			chapEnds := distributeChapters(nodeChapters)

			nodeStart := curMonth
			var phasesJSON []nodePhaseJSON
			chapStart := 0
			for i, ap := range activePhases {
				phaseStart := curMonth
				phaseEnd := curMonth.AddDate(0, phaseMonths[i]-1, 0)

				chapSlice := nodeChapters[chapStart:chapEnds[i]]
				chapNames := make([]string, 0, len(chapSlice))
				for _, ch := range chapSlice {
					chapNames = append(chapNames, ch.Name)
				}

				phasesJSON = append(phasesJSON, nodePhaseJSON{
					Name:         phaseNames[ap.index],
					StartMonth:   phaseStart.Format("2006-01"),
					EndMonth:     phaseEnd.Format("2006-01"),
					Chapters:     chapNames,
					DrillEnabled: phaseDrillEnabled[ap.index],
				})
				// Accumulate longPhaseInfo for mid-term derivation.
				longInfos = append(longInfos, longPhaseInfo{
					StartMonth: phaseStart,
					EndMonth:   phaseEnd,
					Chapters:   chapSlice,
				})

				chapStart = chapEnds[i]
				curMonth = phaseEnd.AddDate(0, 1, 0)
			}

			nodeEnd := curMonth.AddDate(0, -1, 0)
			nodesJSON = append(nodesJSON, nodeJSON{
				ID:         node.ID,
				Name:       node.Name,
				StartMonth: nodeStart.Format("2006-01"),
				EndMonth:   nodeEnd.Format("2006-01"),
				Phases:     phasesJSON,
			})
		}

		b, err := json.Marshal(struct {
			Mode  string     `json:"mode"`
			Nodes []nodeJSON `json:"examNodes"`
		}{Mode: mode, Nodes: nodesJSON})
		if err != nil {
			return "", nil, err
		}
		return string(b), longInfos, nil
	}

	// --- Parallel: all exam nodes are prepared simultaneously ---
	// Phases span the full period; within each phase, every exam node contributes chapters.
	type nodeChaptersJSON struct {
		ID       uint     `json:"id"`
		Name     string   `json:"name"`
		Chapters []string `json:"chapters"`
	}
	type parallelPhaseJSON struct {
		Name         string             `json:"name"`
		StartMonth   string             `json:"startMonth"`
		EndMonth     string             `json:"endMonth"`
		DrillEnabled bool               `json:"drillEnabled"`
		ExamNodes    []nodeChaptersJSON `json:"examNodes"`
	}

	phaseMonths := distributeMonths(totalMonths)

	// Precompute chapter distribution for each node across all phases once,
	// so the inner phase loop can look up chapter slices without recomputing.
	type nodeDistrib struct {
		node     *model.SyllabusExamNode
		chapEnds []int // cumulative end indices per active phase
	}
	nodeDistribs := make([]nodeDistrib, len(examNodes))
	for j, node := range examNodes {
		nodeDistribs[j] = nodeDistrib{
			node:     node,
			chapEnds: distributeChapters(node.Chapters),
		}
	}

	var phasesJSON []parallelPhaseJSON
	curMonth := startTime
	for i, ap := range activePhases {
		phaseStart := curMonth
		phaseEnd := curMonth.AddDate(0, phaseMonths[i]-1, 0)

		var nodeEntries []nodeChaptersJSON
		var phaseChapters []*model.Chapter
		for _, nd := range nodeDistribs {
			start := 0
			if i > 0 {
				start = nd.chapEnds[i-1]
			}
			chapSlice := nd.node.Chapters[start:nd.chapEnds[i]]
			phaseChapters = append(phaseChapters, chapSlice...)
			names := make([]string, 0, len(chapSlice))
			for _, ch := range chapSlice {
				names = append(names, ch.Name)
			}
			nodeEntries = append(nodeEntries, nodeChaptersJSON{ID: nd.node.ID, Name: nd.node.Name, Chapters: names})
		}

		phasesJSON = append(phasesJSON, parallelPhaseJSON{
			Name:         phaseNames[ap.index],
			StartMonth:   phaseStart.Format("2006-01"),
			EndMonth:     phaseEnd.Format("2006-01"),
			DrillEnabled: phaseDrillEnabled[ap.index],
			ExamNodes:    nodeEntries,
		})
		longInfos = append(longInfos, longPhaseInfo{
			StartMonth: phaseStart,
			EndMonth:   phaseEnd,
			Chapters:   phaseChapters,
		})

		curMonth = phaseEnd.AddDate(0, 1, 0)
	}

	b, err := json.Marshal(struct {
		Mode   string              `json:"mode"`
		Phases []parallelPhaseJSON `json:"phases"`
	}{Mode: mode, Phases: phasesJSON})
	if err != nil {
		return "", nil, err
	}
	return string(b), longInfos, nil
}

// buildLongPlanContentPerNode builds the long-term plan from per-node independent time ranges (Plan A).
// Each ExamNodeSchedule specifies its own startMonth/endMonth; nodes are processed independently.
// Nodes with overlapping time ranges run in parallel (by calendar), nodes with non-overlapping ranges
// run sequentially — this is implicit in the schedule and requires no special mode flag.
//
// Each node's time span is divided among active phases using phaseRatios proportionally.
// Chapters of each node are also divided among phases by the same ratios.
// When PhaseRatios don't sum to 100, they are treated as relative weights.
// Zero-ratio phases are skipped.
func buildLongPlanContentPerNode(ratios []int, schedules []model.ExamNodeSchedule, nodeMap map[uint]*model.SyllabusExamNode) (string, []longPhaseInfo, error) {
	type activePhase struct {
		index int
		ratio int
	}
	var activePhases []activePhase
	for i, r := range ratios {
		if r > 0 {
			activePhases = append(activePhases, activePhase{index: i, ratio: r})
		}
	}

	distributeMonths := func(nMonths int) []int {
		sumRatio := 0
		for _, p := range activePhases {
			sumRatio += p.ratio
		}
		months := make([]int, len(activePhases))
		alloc := 0
		for i, p := range activePhases {
			m := int(math.Round(float64(nMonths) * float64(p.ratio) / float64(sumRatio)))
			if m < 1 {
				m = 1
			}
			months[i] = m
			alloc += m
		}
		for i := len(activePhases) - 1; i >= 0 && alloc > nMonths; i-- {
			excess := alloc - nMonths
			reduce := months[i] - 1
			if reduce > excess {
				reduce = excess
			}
			months[i] -= reduce
			alloc -= reduce
		}
		return months
	}

	distributeChapters := func(chapters []*model.Chapter) []int {
		n := len(chapters)
		sumRatio := 0
		for _, p := range activePhases {
			sumRatio += p.ratio
		}
		ends := make([]int, len(activePhases))
		idx := 0
		for i, p := range activePhases {
			count := int(math.Round(float64(n) * float64(p.ratio) / float64(sumRatio)))
			end := idx + count
			if end > n {
				end = n
			}
			ends[i] = end
			idx = end
		}
		if idx < n {
			ends[len(ends)-1] = n
		}
		return ends
	}

	type nodePhaseJSON struct {
		Name         string   `json:"name"`
		StartMonth   string   `json:"startMonth"`
		EndMonth     string   `json:"endMonth"`
		Chapters     []string `json:"chapters"`
		DrillEnabled bool     `json:"drillEnabled"`
	}
	type nodeJSON struct {
		ID         uint            `json:"id"`
		Name       string          `json:"name"`
		StartMonth string          `json:"startMonth"`
		EndMonth   string          `json:"endMonth"`
		Phases     []nodePhaseJSON `json:"phases"`
	}

	var nodesJSON []nodeJSON
	var longInfos []longPhaseInfo

	for _, sched := range schedules {
		node, ok := nodeMap[sched.ExamNodeId]
		if !ok {
			return "", nil, fmt.Errorf("考试节点 %d 不存在", sched.ExamNodeId)
		}
		nodeStart, err := parseYearMonth(sched.StartMonth)
		if err != nil {
			return "", nil, fmt.Errorf("考试节点 %d 的 startMonth 无效: %v", sched.ExamNodeId, err)
		}
		nodeEnd, err := parseYearMonth(sched.EndMonth)
		if err != nil {
			return "", nil, fmt.Errorf("考试节点 %d 的 endMonth 无效: %v", sched.ExamNodeId, err)
		}
		if nodeEnd.Before(nodeStart) {
			return "", nil, fmt.Errorf("考试节点 %d 的 endMonth 不能早于 startMonth", sched.ExamNodeId)
		}

		nodeMonths := monthsBetween(nodeStart, nodeEnd)
		phaseMonths := distributeMonths(nodeMonths)
		chapEnds := distributeChapters(node.Chapters)

		var phasesJSON []nodePhaseJSON
		curMonth := nodeStart
		chapStart := 0
		for i, ap := range activePhases {
			phaseStart := curMonth
			phaseEnd := curMonth.AddDate(0, phaseMonths[i]-1, 0)

			chapSlice := node.Chapters[chapStart:chapEnds[i]]
			chapNames := make([]string, 0, len(chapSlice))
			for _, ch := range chapSlice {
				chapNames = append(chapNames, ch.Name)
			}

			phasesJSON = append(phasesJSON, nodePhaseJSON{
				Name:         phaseNames[ap.index],
				StartMonth:   phaseStart.Format("2006-01"),
				EndMonth:     phaseEnd.Format("2006-01"),
				Chapters:     chapNames,
				DrillEnabled: phaseDrillEnabled[ap.index],
			})
			longInfos = append(longInfos, longPhaseInfo{
				StartMonth: phaseStart,
				EndMonth:   phaseEnd,
				Chapters:   chapSlice,
			})

			chapStart = chapEnds[i]
			curMonth = phaseEnd.AddDate(0, 1, 0)
		}

		nodesJSON = append(nodesJSON, nodeJSON{
			ID:         node.ID,
			Name:       node.Name,
			StartMonth: nodeStart.Format("2006-01"),
			EndMonth:   nodeEnd.Format("2006-01"),
			Phases:     phasesJSON,
		})
	}

	b, err := json.Marshal(struct {
		Mode  string     `json:"mode"`
		Nodes []nodeJSON `json:"examNodes"`
	}{Mode: "per-node", Nodes: nodesJSON})
	if err != nil {
		return "", nil, err
	}
	return string(b), longInfos, nil
}

// buildMidPlanContent generates the JSON content for the mid-term plan (8-week schedule)
// using chapters from the first 2 months of the long-term plan.
func buildMidPlanContent(longPhases []longPhaseInfo, startTime time.Time, allChapters []*model.Chapter) (string, []midWeekInfo, error) {
	// Collect root chapters scheduled in the first 2 months of the long-term plan.
	// afterFirst2Months is the exclusive upper bound; phases starting on or after this are outside
	// the first 2 months. Using !Before() is equivalent to >= (i.e. on or after).
	afterFirst2Months := startTime.AddDate(0, 2, 0)
	var first2MonthChapters []*model.Chapter
	for _, p := range longPhases {
		if !p.StartMonth.Before(afterFirst2Months) { // p.StartMonth >= afterFirst2Months → stop
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
// 以考试节点为核心组织计划：每个考试节点都包含完整的新课→一轮复习→专题复习→集中刷题四个阶段。
// 支持顺序模式（sequential，先完成一个考试节点再进入下一个）和并行模式（parallel，同时准备所有节点）。
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

	// Fetch exam nodes for the syllabus (ordered by sort_order ASC).
	examNodes, err := repository.ExamNodeRepo.FindBySyllabusID(req.SyllabusId)
	if err != nil {
		return nil, fmt.Errorf("获取考试节点失败: %v", err)
	}
	if len(examNodes) == 0 {
		return nil, errors.New("考纲暂无考试节点，请先添加考试节点后再生成计划")
	}

	// Fetch all chapters for the syllabus (used by mid/short plan derivation).
	allChapters, err := repository.ChapterRepo.FindAll(&model.ChapterQuery{SyllabusId: req.SyllabusId})
	if err != nil {
		return nil, fmt.Errorf("获取章节失败: %v", err)
	}

	var longContent string
	var longPhases []longPhaseInfo
	var startTime time.Time

	if len(req.ExamNodes) > 0 {
		// Plan A: per-node independent time ranges.
		// Build a map from examNodeId -> examNode for fast lookup.
		nodeMap := make(map[uint]*model.SyllabusExamNode, len(examNodes))
		for _, n := range examNodes {
			nodeMap[n.ID] = n
		}
		longContent, longPhases, err = buildLongPlanContentPerNode(req.PhaseRatios, req.ExamNodes, nodeMap)
		if err != nil {
			return nil, fmt.Errorf("生成长期计划失败: %v", err)
		}
		// Use the earliest startMonth across all schedules as the anchor for mid/short plans.
		for i, sched := range req.ExamNodes {
			t, parseErr := parseYearMonth(sched.StartMonth)
			if parseErr != nil {
				return nil, fmt.Errorf("考试节点 %d 的 startMonth 无效: %v", sched.ExamNodeId, parseErr)
			}
			if i == 0 || t.Before(startTime) {
				startTime = t
			}
		}
	} else {
		// Fallback: global startMonth/endMonth with sequential distribution.
		if req.StartMonth == "" || req.EndMonth == "" {
			return nil, errors.New("examNodes 为空时 startMonth 和 endMonth 不能为空")
		}
		startTime, err = parseYearMonth(req.StartMonth)
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
		longContent, longPhases, err = buildLongPlanContent(req.PhaseRatios, startTime, totalMonths, ExamNodeModeSequential, examNodes)
		if err != nil {
			return nil, fmt.Errorf("生成长期计划失败: %v", err)
		}
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
