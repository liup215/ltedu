package cmd

import (
	"edu/cli/client"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var knowledgePointCmd = &cobra.Command{
	Use:   "knowledge-point",
	Short: "Manage knowledge points (知识点管理)",
}

// ---- list ----

var (
	kpListPage       int
	kpListPageSize   int
	kpListSyllabusID uint
	kpListChapterID  uint
	kpListName       string
	kpListDifficulty string
)

var kpListCmd = &cobra.Command{
	Use:   "list",
	Short: "List knowledge points (requires --syllabus-id or --chapter-id)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if kpListSyllabusID == 0 && kpListChapterID == 0 {
			return fmt.Errorf("--syllabus-id or --chapter-id is required")
		}
		c := client.NewClient()
		var result struct {
			List  []map[string]interface{} `json:"list"`
			Total int                      `json:"total"`
		}
		body := map[string]interface{}{
			"pageIndex":  kpListPage,
			"pageSize":   kpListPageSize,
			"syllabusId": kpListSyllabusID,
			"chapterId":  kpListChapterID,
			"name":       kpListName,
			"difficulty": kpListDifficulty,
		}
		if err := c.PostAndDecode("/v1/knowledge-point/list", body, &result); err != nil {
			return err
		}
		fmt.Printf("Total: %d\n\n", result.Total)
		headers := []string{"ID", "Name", "ChapterId", "Difficulty", "EstimatedMinutes", "OrderIndex"}
		rows := make([][]string, 0, len(result.List))
		for _, item := range result.List {
			rows = append(rows, []string{
				fmtFloat(item["id"]),
				fmtStr(item["name"]),
				fmtFloat(item["chapterId"]),
				fmtStr(item["difficulty"]),
				fmtFloat(item["estimatedMinutes"]),
				fmtFloat(item["orderIndex"]),
			})
		}
		printTable(headers, rows)
		return nil
	},
}

// ---- get ----

var kpGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get a knowledge point by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid id: %s", args[0])
		}
		c := client.NewClient()
		var result interface{}
		if err := c.PostAndDecode("/v1/knowledge-point/byId", map[string]interface{}{"id": id}, &result); err != nil {
			return err
		}
		prettyPrint(result)
		return nil
	},
}

// ---- create ----

var (
	kpCreateName             string
	kpCreateChapterID        uint
	kpCreateDescription      string
	kpCreateDifficulty       string
	kpCreateEstimatedMinutes int
	kpCreateOrderIndex       int
)

var kpCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new knowledge point (requires --name and --chapter-id)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if kpCreateName == "" {
			return fmt.Errorf("--name is required")
		}
		if kpCreateChapterID == 0 {
			return fmt.Errorf("--chapter-id is required")
		}
		c := client.NewClient()
		body := map[string]interface{}{
			"name":             kpCreateName,
			"chapterId":        kpCreateChapterID,
			"description":      kpCreateDescription,
			"difficulty":       kpCreateDifficulty,
			"estimatedMinutes": kpCreateEstimatedMinutes,
			"orderIndex":       kpCreateOrderIndex,
		}
		var result interface{}
		if err := c.PostAndDecode("/v1/knowledge-point/create", body, &result); err != nil {
			return err
		}
		fmt.Println("Knowledge point created successfully.")
		prettyPrint(result)
		return nil
	},
}

// ---- edit ----

var (
	kpEditID                uint
	kpEditName              string
	kpEditDescription       string
	kpEditDifficulty        string
	kpEditEstimatedMinutes  int
	kpEditOrderIndex        int
)

var kpEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a knowledge point (requires --id)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if kpEditID == 0 {
			return fmt.Errorf("--id is required")
		}
		c := client.NewClient()
		body := map[string]interface{}{
			"id":               kpEditID,
			"name":             kpEditName,
			"description":      kpEditDescription,
			"difficulty":       kpEditDifficulty,
			"estimatedMinutes": kpEditEstimatedMinutes,
			"orderIndex":       kpEditOrderIndex,
		}
		if err := c.PostAndDecode("/v1/knowledge-point/edit", body, nil); err != nil {
			return err
		}
		fmt.Println("Knowledge point updated successfully.")
		return nil
	},
}

// ---- delete ----

var kpDeleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete a knowledge point by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid id: %s", args[0])
		}
		c := client.NewClient()
		if err := c.PostAndDecode("/v1/knowledge-point/delete", map[string]interface{}{"id": id}, nil); err != nil {
			return err
		}
		fmt.Printf("Knowledge point %d deleted successfully.\n", id)
		return nil
	},
}

// ---- link-question ----

var (
	kpLinkQuestionKPID         uint
	kpLinkQuestionQuestionID   uint
)

var kpLinkQuestionCmd = &cobra.Command{
	Use:   "link-question",
	Short: "Link a question to a knowledge point (requires --knowledge-point-id and --question-id)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if kpLinkQuestionKPID == 0 {
			return fmt.Errorf("--knowledge-point-id is required")
		}
		if kpLinkQuestionQuestionID == 0 {
			return fmt.Errorf("--question-id is required")
		}
		c := client.NewClient()
		body := map[string]interface{}{
			"knowledgePointId": kpLinkQuestionKPID,
			"questionId":       kpLinkQuestionQuestionID,
		}
		if err := c.PostAndDecode("/v1/knowledge-point/link-question", body, nil); err != nil {
			return err
		}
		fmt.Printf("Question %d linked to knowledge point %d successfully.\n", kpLinkQuestionQuestionID, kpLinkQuestionKPID)
		return nil
	},
}

// ---- unlink-question ----

var (
	kpUnlinkQuestionKPID       uint
	kpUnlinkQuestionQuestionID uint
)

var kpUnlinkQuestionCmd = &cobra.Command{
	Use:   "unlink-question",
	Short: "Remove a question-knowledge point link (requires --knowledge-point-id and --question-id)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if kpUnlinkQuestionKPID == 0 {
			return fmt.Errorf("--knowledge-point-id is required")
		}
		if kpUnlinkQuestionQuestionID == 0 {
			return fmt.Errorf("--question-id is required")
		}
		c := client.NewClient()
		body := map[string]interface{}{
			"knowledgePointId": kpUnlinkQuestionKPID,
			"questionId":       kpUnlinkQuestionQuestionID,
		}
		if err := c.PostAndDecode("/v1/knowledge-point/unlink-question", body, nil); err != nil {
			return err
		}
		fmt.Printf("Question %d unlinked from knowledge point %d successfully.\n", kpUnlinkQuestionQuestionID, kpUnlinkQuestionKPID)
		return nil
	},
}

// ---- generate ----

var (
	kpGenerateChapterID uint
	kpGenerateMode      string
)

var kpGenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "AI-generate knowledge points for a chapter (requires --chapter-id)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if kpGenerateChapterID == 0 {
			return fmt.Errorf("--chapter-id is required")
		}
		c := client.NewClient()
		body := map[string]interface{}{
			"chapterId": kpGenerateChapterID,
			"mode":      kpGenerateMode,
		}
		var result struct {
			Keypoints []map[string]interface{} `json:"keypoints"`
			Count     int                      `json:"count"`
		}
		if err := c.PostAndDecode("/v1/knowledge-points/generate", body, &result); err != nil {
			return err
		}
		fmt.Printf("Generated %d knowledge point(s) for chapter %d:\n\n", result.Count, kpGenerateChapterID)
		headers := []string{"ID", "Name", "Difficulty", "Minutes", "Confidence"}
		rows := make([][]string, 0, len(result.Keypoints))
		for _, kp := range result.Keypoints {
			rows = append(rows, []string{
				fmtFloat(kp["id"]),
				fmtStr(kp["name"]),
				fmtStr(kp["difficulty"]),
				fmtFloat(kp["estimatedMinutes"]),
				fmt.Sprintf("%.2f", toFloat64(kp["confidenceScore"])),
			})
		}
		printTable(headers, rows)
		return nil
	},
}

// ---- auto-link-question ----

var (
	kpAutoLinkQuestionID   uint
	kpAutoLinkChapterID    uint
	kpAutoLinkSyllabusID   uint
	kpAutoLinkIntelligent  bool
)

var kpAutoLinkQuestionCmd = &cobra.Command{
	Use:   "auto-link-question",
	Short: "Auto-link a question to knowledge points using AI (requires --question-id and --syllabus-id or --chapter-id)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if kpAutoLinkQuestionID == 0 {
			return fmt.Errorf("--question-id is required")
		}
		if kpAutoLinkSyllabusID == 0 && kpAutoLinkChapterID == 0 {
			return fmt.Errorf("--syllabus-id or --chapter-id is required")
		}
		c := client.NewClient()
		var endpoint string
		var body map[string]interface{}
		if kpAutoLinkIntelligent {
			if kpAutoLinkSyllabusID == 0 {
				return fmt.Errorf("--syllabus-id is required for intelligent mode")
			}
			endpoint = "/v1/question/auto-link-keypoints-intelligent"
			body = map[string]interface{}{
				"questionId": kpAutoLinkQuestionID,
				"syllabusId": kpAutoLinkSyllabusID,
			}
		} else {
			endpoint = "/v1/question/auto-link-keypoints"
			body = map[string]interface{}{
				"questionId": kpAutoLinkQuestionID,
				"chapterId":  kpAutoLinkChapterID,
				"syllabusId": kpAutoLinkSyllabusID,
			}
		}
		var result struct {
			LinkedKeypoints []interface{} `json:"linkedKeypoints"`
			Count           int           `json:"count"`
		}
		if err := c.PostAndDecode(endpoint, body, &result); err != nil {
			return err
		}
		fmt.Printf("Question %d linked to %d knowledge points.\n", kpAutoLinkQuestionID, result.Count)
		prettyPrint(result.LinkedKeypoints)
		return nil
	},
}

// ---- auto-migrate ----

var (
	kpAutoMigrateSyllabusID      uint
	kpAutoMigrateGenerateKP      bool
	kpAutoMigrateLinkQuestions   bool
	kpAutoMigrateBatchSize       int
)

var kpAutoMigrateCmd = &cobra.Command{
	Use:   "auto-migrate",
	Short: "Batch auto-migrate knowledge points for a syllabus (requires --syllabus-id, admin only)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if kpAutoMigrateSyllabusID == 0 {
			return fmt.Errorf("--syllabus-id is required")
		}
		c := client.NewClient()
		body := map[string]interface{}{
			"syllabusId": kpAutoMigrateSyllabusID,
			"options": map[string]interface{}{
				"generateKeypoints": kpAutoMigrateGenerateKP,
				"linkQuestions":     kpAutoMigrateLinkQuestions,
				"batchSize":         kpAutoMigrateBatchSize,
			},
		}
		var result interface{}
		if err := c.PostAndDecode("/v1/syllabus/auto-migrate-keypoints", body, &result); err != nil {
			return err
		}
		fmt.Println("Auto-migration completed.")
		prettyPrint(result)
		return nil
	},
}

// toFloat64 safely converts an interface{} JSON number to float64.
func toFloat64(v interface{}) float64 {
	if v == nil {
		return 0
	}
	if f, ok := v.(float64); ok {
		return f
	}
	return 0
}

func init() {
	kpListCmd.Flags().IntVar(&kpListPage, "page", 1, "Page number")
	kpListCmd.Flags().IntVar(&kpListPageSize, "page-size", 20, "Page size")
	kpListCmd.Flags().UintVar(&kpListSyllabusID, "syllabus-id", 0, "Filter by syllabus ID")
	kpListCmd.Flags().UintVar(&kpListChapterID, "chapter-id", 0, "Filter by chapter ID")
	kpListCmd.Flags().StringVar(&kpListName, "name", "", "Filter by name")
	kpListCmd.Flags().StringVar(&kpListDifficulty, "difficulty", "", "Filter by difficulty (basic/medium/hard)")

	kpCreateCmd.Flags().StringVar(&kpCreateName, "name", "", "Knowledge point name (required)")
	kpCreateCmd.Flags().UintVar(&kpCreateChapterID, "chapter-id", 0, "Chapter ID (required)")
	kpCreateCmd.Flags().StringVar(&kpCreateDescription, "description", "", "Description")
	kpCreateCmd.Flags().StringVar(&kpCreateDifficulty, "difficulty", "basic", "Difficulty (basic/medium/hard)")
	kpCreateCmd.Flags().IntVar(&kpCreateEstimatedMinutes, "estimated-minutes", 0, "Estimated study minutes")
	kpCreateCmd.Flags().IntVar(&kpCreateOrderIndex, "order-index", 0, "Order index")

	kpEditCmd.Flags().UintVar(&kpEditID, "id", 0, "Knowledge point ID (required)")
	kpEditCmd.Flags().StringVar(&kpEditName, "name", "", "New name")
	kpEditCmd.Flags().StringVar(&kpEditDescription, "description", "", "New description")
	kpEditCmd.Flags().StringVar(&kpEditDifficulty, "difficulty", "", "New difficulty (basic/medium/hard)")
	kpEditCmd.Flags().IntVar(&kpEditEstimatedMinutes, "estimated-minutes", 0, "New estimated study minutes")
	kpEditCmd.Flags().IntVar(&kpEditOrderIndex, "order-index", 0, "New order index")

	kpLinkQuestionCmd.Flags().UintVar(&kpLinkQuestionKPID, "knowledge-point-id", 0, "Knowledge point ID (required)")
	kpLinkQuestionCmd.Flags().UintVar(&kpLinkQuestionQuestionID, "question-id", 0, "Question ID (required)")

	kpUnlinkQuestionCmd.Flags().UintVar(&kpUnlinkQuestionKPID, "knowledge-point-id", 0, "Knowledge point ID (required)")
	kpUnlinkQuestionCmd.Flags().UintVar(&kpUnlinkQuestionQuestionID, "question-id", 0, "Question ID (required)")

	kpGenerateCmd.Flags().UintVar(&kpGenerateChapterID, "chapter-id", 0, "Chapter ID (required)")
	kpGenerateCmd.Flags().StringVar(&kpGenerateMode, "mode", "auto", "Generation mode (auto/manual)")

	kpAutoLinkQuestionCmd.Flags().UintVar(&kpAutoLinkQuestionID, "question-id", 0, "Question ID (required)")
	kpAutoLinkQuestionCmd.Flags().UintVar(&kpAutoLinkChapterID, "chapter-id", 0, "Chapter ID (scope for linking)")
	kpAutoLinkQuestionCmd.Flags().UintVar(&kpAutoLinkSyllabusID, "syllabus-id", 0, "Syllabus ID (scope for linking)")
	kpAutoLinkQuestionCmd.Flags().BoolVar(&kpAutoLinkIntelligent, "intelligent", false, "Use two-phase intelligent linking (requires --syllabus-id)")

	kpAutoMigrateCmd.Flags().UintVar(&kpAutoMigrateSyllabusID, "syllabus-id", 0, "Syllabus ID (required)")
	kpAutoMigrateCmd.Flags().BoolVar(&kpAutoMigrateGenerateKP, "generate-keypoints", true, "Generate knowledge points for chapters")
	kpAutoMigrateCmd.Flags().BoolVar(&kpAutoMigrateLinkQuestions, "link-questions", true, "Auto-link existing questions to knowledge points")
	kpAutoMigrateCmd.Flags().IntVar(&kpAutoMigrateBatchSize, "batch-size", 10, "Batch size for processing")

	knowledgePointCmd.AddCommand(kpListCmd)
	knowledgePointCmd.AddCommand(kpGetCmd)
	knowledgePointCmd.AddCommand(kpCreateCmd)
	knowledgePointCmd.AddCommand(kpEditCmd)
	knowledgePointCmd.AddCommand(kpDeleteCmd)
	knowledgePointCmd.AddCommand(kpLinkQuestionCmd)
	knowledgePointCmd.AddCommand(kpUnlinkQuestionCmd)
	knowledgePointCmd.AddCommand(kpGenerateCmd)
	knowledgePointCmd.AddCommand(kpAutoLinkQuestionCmd)
	knowledgePointCmd.AddCommand(kpAutoMigrateCmd)
}
