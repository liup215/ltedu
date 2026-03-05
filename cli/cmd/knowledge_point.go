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
)

var kpListCmd = &cobra.Command{
	Use:   "list",
	Short: "List knowledge points (optionally filter by --syllabus-id or --chapter-id)",
	RunE: func(cmd *cobra.Command, args []string) error {
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
		}
		if err := c.PostAndDecode("/v1/knowledge-point/list", body, &result); err != nil {
			return err
		}
		fmt.Printf("Total: %d\n\n", result.Total)
		headers := []string{"ID", "Name", "ChapterId", "Difficulty", "EstimatedMinutes"}
		rows := make([][]string, 0, len(result.List))
		for _, item := range result.List {
			rows = append(rows, []string{
				fmtFloat(item["id"]),
				fmtStr(item["name"]),
				fmtFloat(item["chapterId"]),
				fmtStr(item["difficulty"]),
				fmtFloat(item["estimatedMinutes"]),
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
	kpCreateDescription      string
	kpCreateChapterID        uint
	kpCreateDifficulty       string
	kpCreateEstimatedMinutes int
)

var kpCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new knowledge point (requires --chapter-id and --name)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if kpCreateName == "" {
			return fmt.Errorf("--name is required")
		}
		if kpCreateChapterID == 0 {
			return fmt.Errorf("--chapter-id is required")
		}
		if kpCreateDifficulty == "" {
			kpCreateDifficulty = "medium"
		}
		c := client.NewClient()
		body := map[string]interface{}{
			"name":             kpCreateName,
			"description":      kpCreateDescription,
			"chapterId":        kpCreateChapterID,
			"difficulty":       kpCreateDifficulty,
			"estimatedMinutes": kpCreateEstimatedMinutes,
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

// ---- generate ----

var kpGenerateChapterID uint

var kpGenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "AI-generate knowledge points for a chapter (requires --chapter-id)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if kpGenerateChapterID == 0 {
			return fmt.Errorf("--chapter-id is required")
		}
		c := client.NewClient()
		var result interface{}
		if err := c.PostAndDecode("/v1/chapter/generate-keypoints", map[string]interface{}{"chapterId": kpGenerateChapterID}, &result); err != nil {
			return err
		}
		fmt.Println("Knowledge points generated successfully.")
		prettyPrint(result)
		return nil
	},
}

// ---- by-chapter ----

var kpByChapterID uint

var kpByChapterCmd = &cobra.Command{
	Use:   "by-chapter",
	Short: "List knowledge points for a specific chapter (requires --chapter-id)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if kpByChapterID == 0 {
			return fmt.Errorf("--chapter-id is required")
		}
		c := client.NewClient()
		var result struct {
			List  []map[string]interface{} `json:"list"`
			Total int                      `json:"total"`
		}
		if err := c.PostAndDecode("/v1/knowledge-point/byChapter", map[string]interface{}{"chapterId": kpByChapterID}, &result); err != nil {
			return err
		}
		fmt.Printf("Total: %d\n\n", result.Total)
		headers := []string{"ID", "Name", "Difficulty", "EstimatedMinutes", "Description"}
		rows := make([][]string, 0, len(result.List))
		for _, item := range result.List {
			rows = append(rows, []string{
				fmtFloat(item["id"]),
				fmtStr(item["name"]),
				fmtStr(item["difficulty"]),
				fmtFloat(item["estimatedMinutes"]),
				fmtStr(item["description"]),
			})
		}
		printTable(headers, rows)
		return nil
	},
}

func init() {
	kpListCmd.Flags().IntVar(&kpListPage, "page", 1, "Page number")
	kpListCmd.Flags().IntVar(&kpListPageSize, "page-size", 20, "Page size")
	kpListCmd.Flags().UintVar(&kpListSyllabusID, "syllabus-id", 0, "Filter by syllabus ID")
	kpListCmd.Flags().UintVar(&kpListChapterID, "chapter-id", 0, "Filter by chapter ID")

	kpCreateCmd.Flags().StringVar(&kpCreateName, "name", "", "Knowledge point name (required)")
	kpCreateCmd.Flags().StringVar(&kpCreateDescription, "description", "", "Description")
	kpCreateCmd.Flags().UintVar(&kpCreateChapterID, "chapter-id", 0, "Chapter ID (required)")
	kpCreateCmd.Flags().StringVar(&kpCreateDifficulty, "difficulty", "medium", "Difficulty: basic, medium, hard")
	kpCreateCmd.Flags().IntVar(&kpCreateEstimatedMinutes, "minutes", 30, "Estimated study minutes")

	kpEditCmd.Flags().UintVar(&kpEditID, "id", 0, "Knowledge point ID (required)")
	kpEditCmd.Flags().StringVar(&kpEditName, "name", "", "New name")
	kpEditCmd.Flags().StringVar(&kpEditDescription, "description", "", "New description")
	kpEditCmd.Flags().StringVar(&kpEditDifficulty, "difficulty", "", "New difficulty: basic, medium, hard")
	kpEditCmd.Flags().IntVar(&kpEditEstimatedMinutes, "minutes", 0, "New estimated study minutes")

	kpGenerateCmd.Flags().UintVar(&kpGenerateChapterID, "chapter-id", 0, "Chapter ID (required)")

	kpByChapterCmd.Flags().UintVar(&kpByChapterID, "chapter-id", 0, "Chapter ID (required)")

	knowledgePointCmd.AddCommand(kpListCmd)
	knowledgePointCmd.AddCommand(kpGetCmd)
	knowledgePointCmd.AddCommand(kpCreateCmd)
	knowledgePointCmd.AddCommand(kpEditCmd)
	knowledgePointCmd.AddCommand(kpDeleteCmd)
	knowledgePointCmd.AddCommand(kpGenerateCmd)
	knowledgePointCmd.AddCommand(kpByChapterCmd)
}
