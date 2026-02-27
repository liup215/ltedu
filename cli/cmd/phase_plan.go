package cmd

import (
	"edu/cli/client"
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var phasePlanCmd = &cobra.Command{
	Use:   "phase-plan",
	Short: "Manage phase plans within a learning plan (阶段性计划管理)",
}

// ---- list ----

var phasePlanListPlanID uint

var phasePlanListCmd = &cobra.Command{
	Use:   "list",
	Short: "List phase plans for a learning plan (列出学习计划的阶段性计划)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if phasePlanListPlanID == 0 {
			return fmt.Errorf("--plan-id is required")
		}
		c := client.NewClient()
		var result struct {
			List  []map[string]interface{} `json:"list"`
			Total int                      `json:"total"`
		}
		if err := c.PostAndDecode("/v1/learning-plan/phase/list", map[string]interface{}{"planId": phasePlanListPlanID}, &result); err != nil {
			return err
		}
		fmt.Printf("Total: %d\n\n", result.Total)
		headers := []string{"ID", "Title", "ExamNodeId", "StartDate", "EndDate", "SortOrder", "Chapters"}
		rows := make([][]string, 0, len(result.List))
		for _, item := range result.List {
			chapCount := 0
			if chs, ok := item["chapters"].([]interface{}); ok {
				chapCount = len(chs)
			}
			rows = append(rows, []string{
				fmtFloat(item["id"]),
				fmtStr(item["title"]),
				fmtFloat(item["examNodeId"]),
				fmtStr(item["startDate"]),
				fmtStr(item["endDate"]),
				fmtFloat(item["sortOrder"]),
				strconv.Itoa(chapCount),
			})
		}
		printTable(headers, rows)
		return nil
	},
}

// ---- get ----

var phasePlanGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get a phase plan by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid id: %s", args[0])
		}
		c := client.NewClient()
		var result interface{}
		if err := c.PostAndDecode("/v1/learning-plan/phase/byId", map[string]interface{}{"id": id}, &result); err != nil {
			return err
		}
		prettyPrint(result)
		return nil
	},
}

// ---- create ----

var (
	phasePlanCreatePlanID     uint
	phasePlanCreateExamNodeID uint
	phasePlanCreateTitle      string
	phasePlanCreateStartDate  string
	phasePlanCreateEndDate    string
	phasePlanCreateSortOrder  int
)

var phasePlanCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a phase plan for a learning plan (为学习计划创建阶段性计划)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if phasePlanCreatePlanID == 0 || phasePlanCreateExamNodeID == 0 {
			return fmt.Errorf("--plan-id and --exam-node-id are required")
		}
		body := map[string]interface{}{
			"planId":     phasePlanCreatePlanID,
			"examNodeId": phasePlanCreateExamNodeID,
			"title":      phasePlanCreateTitle,
			"sortOrder":  phasePlanCreateSortOrder,
		}
		if phasePlanCreateStartDate != "" {
			t, err := time.Parse("2006-01-02", phasePlanCreateStartDate)
			if err != nil {
				return fmt.Errorf("invalid --start-date format, expected YYYY-MM-DD")
			}
			body["startDate"] = t.Format(time.RFC3339)
		}
		if phasePlanCreateEndDate != "" {
			t, err := time.Parse("2006-01-02", phasePlanCreateEndDate)
			if err != nil {
				return fmt.Errorf("invalid --end-date format, expected YYYY-MM-DD")
			}
			body["endDate"] = t.Format(time.RFC3339)
		}
		c := client.NewClient()
		var result interface{}
		if err := c.PostAndDecode("/v1/learning-plan/phase/create", body, &result); err != nil {
			return err
		}
		fmt.Println("Phase plan created successfully.")
		prettyPrint(result)
		return nil
	},
}

// ---- edit ----

var (
	phasePlanEditID        uint
	phasePlanEditTitle     string
	phasePlanEditStartDate string
	phasePlanEditEndDate   string
	phasePlanEditSortOrder int
)

var phasePlanEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Update a phase plan (更新阶段性计划)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if phasePlanEditID == 0 {
			return fmt.Errorf("--id is required")
		}
		body := map[string]interface{}{
			"id":        phasePlanEditID,
			"title":     phasePlanEditTitle,
			"sortOrder": phasePlanEditSortOrder,
		}
		if phasePlanEditStartDate != "" {
			t, err := time.Parse("2006-01-02", phasePlanEditStartDate)
			if err != nil {
				return fmt.Errorf("invalid --start-date format, expected YYYY-MM-DD")
			}
			body["startDate"] = t.Format(time.RFC3339)
		}
		if phasePlanEditEndDate != "" {
			t, err := time.Parse("2006-01-02", phasePlanEditEndDate)
			if err != nil {
				return fmt.Errorf("invalid --end-date format, expected YYYY-MM-DD")
			}
			body["endDate"] = t.Format(time.RFC3339)
		}
		c := client.NewClient()
		var result interface{}
		if err := c.PostAndDecode("/v1/learning-plan/phase/edit", body, &result); err != nil {
			return err
		}
		fmt.Println("Phase plan updated successfully.")
		prettyPrint(result)
		return nil
	},
}

// ---- delete ----

var phasePlanDeleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete a phase plan by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid id: %s", args[0])
		}
		c := client.NewClient()
		if err := c.PostAndDecode("/v1/learning-plan/phase/delete", map[string]interface{}{"id": id}, nil); err != nil {
			return err
		}
		fmt.Printf("Phase plan %d deleted successfully.\n", id)
		return nil
	},
}

// ---- add-chapter ----

var (
	addPPChapterPhasePlanID uint
	addPPChapterChapterID   uint
)

var phasePlanAddChapterCmd = &cobra.Command{
	Use:   "add-chapter",
	Short: "Add a chapter to a phase plan (为阶段性计划添加章节)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if addPPChapterPhasePlanID == 0 || addPPChapterChapterID == 0 {
			return fmt.Errorf("--phase-plan-id and --chapter-id are required")
		}
		c := client.NewClient()
		body := map[string]interface{}{
			"phasePlanId": addPPChapterPhasePlanID,
			"chapterId":   addPPChapterChapterID,
		}
		if err := c.PostAndDecode("/v1/learning-plan/phase/chapter/add", body, nil); err != nil {
			return err
		}
		fmt.Printf("Chapter %d added to phase plan %d successfully.\n", addPPChapterChapterID, addPPChapterPhasePlanID)
		return nil
	},
}

// ---- remove-chapter ----

var (
	removePPChapterPhasePlanID uint
	removePPChapterChapterID   uint
)

var phasePlanRemoveChapterCmd = &cobra.Command{
	Use:   "remove-chapter",
	Short: "Remove a chapter from a phase plan (从阶段性计划移除章节)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if removePPChapterPhasePlanID == 0 || removePPChapterChapterID == 0 {
			return fmt.Errorf("--phase-plan-id and --chapter-id are required")
		}
		c := client.NewClient()
		body := map[string]interface{}{
			"phasePlanId": removePPChapterPhasePlanID,
			"chapterId":   removePPChapterChapterID,
		}
		if err := c.PostAndDecode("/v1/learning-plan/phase/chapter/remove", body, nil); err != nil {
			return err
		}
		fmt.Printf("Chapter %d removed from phase plan %d successfully.\n", removePPChapterChapterID, removePPChapterPhasePlanID)
		return nil
	},
}

func init() {
	phasePlanListCmd.Flags().UintVar(&phasePlanListPlanID, "plan-id", 0, "Learning plan ID (required)")

	phasePlanCreateCmd.Flags().UintVar(&phasePlanCreatePlanID, "plan-id", 0, "Learning plan ID (required)")
	phasePlanCreateCmd.Flags().UintVar(&phasePlanCreateExamNodeID, "exam-node-id", 0, "Exam node ID (required)")
	phasePlanCreateCmd.Flags().StringVar(&phasePlanCreateTitle, "title", "", "Phase plan title")
	phasePlanCreateCmd.Flags().StringVar(&phasePlanCreateStartDate, "start-date", "", "Start date (YYYY-MM-DD)")
	phasePlanCreateCmd.Flags().StringVar(&phasePlanCreateEndDate, "end-date", "", "End date (YYYY-MM-DD)")
	phasePlanCreateCmd.Flags().IntVar(&phasePlanCreateSortOrder, "sort-order", 0, "Sort order")

	phasePlanEditCmd.Flags().UintVar(&phasePlanEditID, "id", 0, "Phase plan ID (required)")
	phasePlanEditCmd.Flags().StringVar(&phasePlanEditTitle, "title", "", "New title")
	phasePlanEditCmd.Flags().StringVar(&phasePlanEditStartDate, "start-date", "", "New start date (YYYY-MM-DD)")
	phasePlanEditCmd.Flags().StringVar(&phasePlanEditEndDate, "end-date", "", "New end date (YYYY-MM-DD)")
	phasePlanEditCmd.Flags().IntVar(&phasePlanEditSortOrder, "sort-order", 0, "New sort order")

	phasePlanAddChapterCmd.Flags().UintVar(&addPPChapterPhasePlanID, "phase-plan-id", 0, "Phase plan ID (required)")
	phasePlanAddChapterCmd.Flags().UintVar(&addPPChapterChapterID, "chapter-id", 0, "Chapter ID (required)")

	phasePlanRemoveChapterCmd.Flags().UintVar(&removePPChapterPhasePlanID, "phase-plan-id", 0, "Phase plan ID (required)")
	phasePlanRemoveChapterCmd.Flags().UintVar(&removePPChapterChapterID, "chapter-id", 0, "Chapter ID (required)")

	phasePlanCmd.AddCommand(phasePlanListCmd)
	phasePlanCmd.AddCommand(phasePlanGetCmd)
	phasePlanCmd.AddCommand(phasePlanCreateCmd)
	phasePlanCmd.AddCommand(phasePlanEditCmd)
	phasePlanCmd.AddCommand(phasePlanDeleteCmd)
	phasePlanCmd.AddCommand(phasePlanAddChapterCmd)
	phasePlanCmd.AddCommand(phasePlanRemoveChapterCmd)
}
