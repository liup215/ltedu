package cmd

import (
	"edu/cli/client"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var learningPlanCmd = &cobra.Command{
	Use:   "learning-plan",
	Short: "Manage student learning plans (学生学习计划管理)",
}

// ---- list ----

var (
	planListClassID  uint
	planListUserID   uint
	planListPlanType string
	planListPage     int
	planListPageSize int
)

var planListCmd = &cobra.Command{
	Use:   "list",
	Short: "List learning plans (学习计划列表)",
	RunE: func(cmd *cobra.Command, args []string) error {
		c := client.NewClient()
		var result struct {
			List  []map[string]interface{} `json:"list"`
			Total int                      `json:"total"`
		}
		body := map[string]interface{}{
			"pageIndex": planListPage,
			"pageSize":  planListPageSize,
		}
		if planListClassID != 0 {
			body["classId"] = planListClassID
		}
		if planListUserID != 0 {
			body["userId"] = planListUserID
		}
		if planListPlanType != "" {
			body["planType"] = planListPlanType
		}
		if err := c.PostAndDecode("/v1/learning-plan/list", body, &result); err != nil {
			return err
		}
		fmt.Printf("Total: %d\n\n", result.Total)
		headers := []string{"ID", "ClassId", "UserId", "PlanType", "Version"}
		rows := make([][]string, 0, len(result.List))
		for _, item := range result.List {
			rows = append(rows, []string{
				fmtFloat(item["id"]),
				fmtFloat(item["classId"]),
				fmtFloat(item["userId"]),
				fmtStr(item["planType"]),
				fmtFloat(item["version"]),
			})
		}
		printTable(headers, rows)
		return nil
	},
}

// ---- get ----

var planGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get a learning plan by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid id: %s", args[0])
		}
		c := client.NewClient()
		var result interface{}
		if err := c.PostAndDecode("/v1/learning-plan/byId", map[string]interface{}{"id": id}, &result); err != nil {
			return err
		}
		prettyPrint(result)
		return nil
	},
}

// ---- create ----

var (
	planCreateClassID  uint
	planCreateUserID   uint
	planCreatePlanType string
	planCreateContent  string
	planCreateComment  string
)

var planCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a learning plan for a student (为学生创建学习计划)",
	Long:  "Create a learning plan. --plan-type must be one of: long (长期), mid (中期), short (短期)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if planCreateClassID == 0 || planCreateUserID == 0 {
			return fmt.Errorf("--class-id and --user-id are required")
		}
		if planCreatePlanType == "" {
			return fmt.Errorf("--plan-type is required (long/mid/short)")
		}
		c := client.NewClient()
		body := map[string]interface{}{
			"classId":  planCreateClassID,
			"userId":   planCreateUserID,
			"planType": planCreatePlanType,
			"content":  planCreateContent,
			"comment":  planCreateComment,
		}
		var result interface{}
		if err := c.PostAndDecode("/v1/learning-plan/create", body, &result); err != nil {
			return err
		}
		fmt.Println("Learning plan created successfully.")
		prettyPrint(result)
		return nil
	},
}

// ---- edit ----

var (
	planEditID      uint
	planEditContent string
	planEditComment string
)

var planEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Update a learning plan (更新学习计划，自动记录新版本)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if planEditID == 0 {
			return fmt.Errorf("--id is required")
		}
		c := client.NewClient()
		body := map[string]interface{}{
			"id":      planEditID,
			"content": planEditContent,
			"comment": planEditComment,
		}
		var result interface{}
		if err := c.PostAndDecode("/v1/learning-plan/edit", body, &result); err != nil {
			return err
		}
		fmt.Println("Learning plan updated successfully.")
		prettyPrint(result)
		return nil
	},
}

// ---- delete ----

var planDeleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete a learning plan by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid id: %s", args[0])
		}
		c := client.NewClient()
		if err := c.PostAndDecode("/v1/learning-plan/delete", map[string]interface{}{"id": id}, nil); err != nil {
			return err
		}
		fmt.Printf("Learning plan %d deleted successfully.\n", id)
		return nil
	},
}

// ---- versions ----

var (
	planVersionsPlanID   uint
	planVersionsPage     int
	planVersionsPageSize int
)

var planVersionsCmd = &cobra.Command{
	Use:   "versions",
	Short: "List version history of a learning plan (查看学习计划历史版本)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if planVersionsPlanID == 0 {
			return fmt.Errorf("--plan-id is required")
		}
		c := client.NewClient()
		var result struct {
			List  []map[string]interface{} `json:"list"`
			Total int                      `json:"total"`
		}
		body := map[string]interface{}{
			"planId":    planVersionsPlanID,
			"pageIndex": planVersionsPage,
			"pageSize":  planVersionsPageSize,
		}
		if err := c.PostAndDecode("/v1/learning-plan/versions", body, &result); err != nil {
			return err
		}
		fmt.Printf("Total versions: %d\n\n", result.Total)
		headers := []string{"Version", "ChangedBy", "Comment", "CreatedAt"}
		rows := make([][]string, 0, len(result.List))
		for _, item := range result.List {
			rows = append(rows, []string{
				fmtFloat(item["version"]),
				fmtFloat(item["changedBy"]),
				fmtStr(item["comment"]),
				fmtStr(item["createdAt"]),
			})
		}
		printTable(headers, rows)
		return nil
	},
}

// ---- rollback ----

var (
	planRollbackPlanID  uint
	planRollbackVersion int
	planRollbackComment string
)

var planRollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "Rollback a learning plan to a previous version (回滚学习计划到历史版本)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if planRollbackPlanID == 0 || planRollbackVersion == 0 {
			return fmt.Errorf("--plan-id and --version are required")
		}
		c := client.NewClient()
		body := map[string]interface{}{
			"planId":  planRollbackPlanID,
			"version": planRollbackVersion,
			"comment": planRollbackComment,
		}
		var result interface{}
		if err := c.PostAndDecode("/v1/learning-plan/rollback", body, &result); err != nil {
			return err
		}
		fmt.Printf("Learning plan %d rolled back to version %d successfully.\n", planRollbackPlanID, planRollbackVersion)
		prettyPrint(result)
		return nil
	},
}

// ---- generate-template ----

var (
	genTplClassID      uint
	genTplSyllabusID   uint
	genTplStartMonth   string
	genTplEndMonth     string
	genTplPhaseRatios  string
	genTplExamNodeMode string
	genTplComment      string
)

var planGenerateTemplateCmd = &cobra.Command{
	Use:   "generate-template",
	Short: "Batch-generate template learning plans for all students in a class (批量生成模板学习计划)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if genTplClassID == 0 || genTplSyllabusID == 0 {
			return fmt.Errorf("--class-id and --syllabus-id are required")
		}
		if genTplStartMonth == "" || genTplEndMonth == "" {
			return fmt.Errorf("--start-month and --end-month are required")
		}

		// Parse comma-separated phase ratios into []int.
		parts := strings.Split(genTplPhaseRatios, ",")
		if len(parts) != 4 {
			return fmt.Errorf("--phase-ratios must be 4 comma-separated integers, e.g. 30,20,20,10")
		}
		ratios := make([]int, 4)
		for i, p := range parts {
			v, err := strconv.Atoi(strings.TrimSpace(p))
			if err != nil {
				return fmt.Errorf("invalid phase ratio value %q: %v", p, err)
			}
			ratios[i] = v
		}

		c := client.NewClient()
		body := map[string]interface{}{
			"classId":      genTplClassID,
			"syllabusId":   genTplSyllabusID,
			"startMonth":   genTplStartMonth,
			"endMonth":     genTplEndMonth,
			"phaseRatios":  ratios,
			"examNodeMode": genTplExamNodeMode,
			"comment":      genTplComment,
		}
		var result struct {
			StudentCount int      `json:"studentCount"`
			Count        int      `json:"count"`
			Errors       []string `json:"errors"`
		}
		if err := c.PostAndDecode("/v1/learning-plan/generateTemplate", body, &result); err != nil {
			return err
		}
		fmt.Printf("Done. Students: %d, Plans created: %d\n", result.StudentCount, result.Count)
		if len(result.Errors) > 0 {
			fmt.Printf("Errors (%d):\n", len(result.Errors))
			for _, e := range result.Errors {
				fmt.Printf("  - %s\n", e)
			}
		}
		return nil
	},
}

func init() {
	planListCmd.Flags().UintVar(&planListClassID, "class-id", 0, "Filter by class ID")
	planListCmd.Flags().UintVar(&planListUserID, "user-id", 0, "Filter by user ID")
	planListCmd.Flags().StringVar(&planListPlanType, "plan-type", "", "Filter by plan type (long/mid/short)")
	planListCmd.Flags().IntVar(&planListPage, "page", 1, "Page number")
	planListCmd.Flags().IntVar(&planListPageSize, "page-size", 20, "Page size")

	planCreateCmd.Flags().UintVar(&planCreateClassID, "class-id", 0, "Class ID (required)")
	planCreateCmd.Flags().UintVar(&planCreateUserID, "user-id", 0, "User (student) ID (required)")
	planCreateCmd.Flags().StringVar(&planCreatePlanType, "plan-type", "", "Plan type: long/mid/short (required)")
	planCreateCmd.Flags().StringVar(&planCreateContent, "content", "", "Plan content")
	planCreateCmd.Flags().StringVar(&planCreateComment, "comment", "", "Initial version comment")

	planEditCmd.Flags().UintVar(&planEditID, "id", 0, "Learning plan ID (required)")
	planEditCmd.Flags().StringVar(&planEditContent, "content", "", "New plan content")
	planEditCmd.Flags().StringVar(&planEditComment, "comment", "", "Version comment")

	planVersionsCmd.Flags().UintVar(&planVersionsPlanID, "plan-id", 0, "Learning plan ID (required)")
	planVersionsCmd.Flags().IntVar(&planVersionsPage, "page", 1, "Page number")
	planVersionsCmd.Flags().IntVar(&planVersionsPageSize, "page-size", 20, "Page size")

	planRollbackCmd.Flags().UintVar(&planRollbackPlanID, "plan-id", 0, "Learning plan ID (required)")
	planRollbackCmd.Flags().IntVar(&planRollbackVersion, "version", 0, "Target version number (required)")
	planRollbackCmd.Flags().StringVar(&planRollbackComment, "comment", "", "Rollback comment")

	planGenerateTemplateCmd.Flags().UintVar(&genTplClassID, "class-id", 0, "Target teaching class ID (required)")
	planGenerateTemplateCmd.Flags().UintVar(&genTplSyllabusID, "syllabus-id", 0, "Syllabus ID to base plans on (required)")
	planGenerateTemplateCmd.Flags().StringVar(&genTplStartMonth, "start-month", "", "Start month in YYYY-MM format (required)")
	planGenerateTemplateCmd.Flags().StringVar(&genTplEndMonth, "end-month", "", "End month in YYYY-MM format (required)")
	planGenerateTemplateCmd.Flags().StringVar(&genTplPhaseRatios, "phase-ratios", "30,20,20,10", "Comma-separated phase ratios summing to <=100, e.g. 30,20,20,10")
	planGenerateTemplateCmd.Flags().StringVar(&genTplExamNodeMode, "exam-node-mode", "sequential", "How to arrange exam nodes: 'sequential' (one after another) or 'parallel' (simultaneously)")
	planGenerateTemplateCmd.Flags().StringVar(&genTplComment, "comment", "", "Version comment for the initial plans")

	learningPlanCmd.AddCommand(planListCmd)
	learningPlanCmd.AddCommand(planGetCmd)
	learningPlanCmd.AddCommand(planCreateCmd)
	learningPlanCmd.AddCommand(planEditCmd)
	learningPlanCmd.AddCommand(planDeleteCmd)
	learningPlanCmd.AddCommand(planVersionsCmd)
	learningPlanCmd.AddCommand(planRollbackCmd)
	learningPlanCmd.AddCommand(planGenerateTemplateCmd)
}
