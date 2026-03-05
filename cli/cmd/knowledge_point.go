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

// ---- generate ----

var knowledgePointGenerateChapterID uint

var knowledgePointGenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "AI-generate knowledge points for a chapter (requires --chapter-id)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if knowledgePointGenerateChapterID == 0 {
			return fmt.Errorf("--chapter-id is required")
		}
		c := client.NewClient()
		var result struct {
			Keypoints []map[string]interface{} `json:"keypoints"`
			Count     int                      `json:"count"`
		}
		body := map[string]interface{}{
			"chapterId": knowledgePointGenerateChapterID,
		}
		if err := c.PostAndDecode("/v1/knowledge-points/generate", body, &result); err != nil {
			return err
		}
		fmt.Printf("Generated %d knowledge point(s) for chapter %d:\n\n", result.Count, knowledgePointGenerateChapterID)
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

// ---- list ----

var (
	knowledgePointListChapterID  uint
	knowledgePointListSyllabusID uint
	knowledgePointListPage       int
	knowledgePointListPageSize   int
)

var knowledgePointListCmd = &cobra.Command{
	Use:   "list",
	Short: "List knowledge points (requires --chapter-id or --syllabus-id)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if knowledgePointListChapterID == 0 && knowledgePointListSyllabusID == 0 {
			return fmt.Errorf("--chapter-id or --syllabus-id is required")
		}
		c := client.NewClient()
		var result struct {
			List  []map[string]interface{} `json:"list"`
			Total int                      `json:"total"`
		}
		body := map[string]interface{}{
			"chapterId":  knowledgePointListChapterID,
			"syllabusId": knowledgePointListSyllabusID,
			"pageIndex":  knowledgePointListPage,
			"pageSize":   knowledgePointListPageSize,
		}
		if err := c.PostAndDecode("/v1/knowledge-point/list", body, &result); err != nil {
			return err
		}
		fmt.Printf("Total: %d\n\n", result.Total)
		headers := []string{"ID", "ChapterId", "Name", "Difficulty", "Minutes"}
		rows := make([][]string, 0, len(result.List))
		for _, kp := range result.List {
			rows = append(rows, []string{
				fmtFloat(kp["id"]),
				fmtFloat(kp["chapterId"]),
				fmtStr(kp["name"]),
				fmtStr(kp["difficulty"]),
				fmtFloat(kp["estimatedMinutes"]),
			})
		}
		printTable(headers, rows)
		return nil
	},
}

// ---- get ----

var knowledgePointGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get a knowledge point by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil || id == 0 {
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
	knowledgePointGenerateCmd.Flags().UintVar(&knowledgePointGenerateChapterID, "chapter-id", 0, "Chapter ID to generate knowledge points for (required)")

	knowledgePointListCmd.Flags().UintVar(&knowledgePointListChapterID, "chapter-id", 0, "Filter by chapter ID")
	knowledgePointListCmd.Flags().UintVar(&knowledgePointListSyllabusID, "syllabus-id", 0, "Filter by syllabus ID")
	knowledgePointListCmd.Flags().IntVar(&knowledgePointListPage, "page", 1, "Page number")
	knowledgePointListCmd.Flags().IntVar(&knowledgePointListPageSize, "page-size", 20, "Page size")

	knowledgePointCmd.AddCommand(knowledgePointGenerateCmd)
	knowledgePointCmd.AddCommand(knowledgePointListCmd)
	knowledgePointCmd.AddCommand(knowledgePointGetCmd)
}
