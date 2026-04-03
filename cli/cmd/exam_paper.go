package cmd

import (
	"edu/cli/client"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var examPaperCmd = &cobra.Command{
	Use:   "exam",
	Short: "Manage exam papers (试卷管理)",
}

// ---- list ----

var (
	examPaperListPage       int
	examPaperListPageSize   int
	examPaperListSyllabusID uint
	examPaperListUserID     uint
)

var examPaperListCmd = &cobra.Command{
	Use:   "list",
	Short: "List exam papers (列出试卷)",
	RunE: func(cmd *cobra.Command, args []string) error {
		c := client.NewClient()
		var result struct {
			List  []map[string]interface{} `json:"list"`
			Total int                      `json:"total"`
		}
		body := map[string]interface{}{
			"pageIndex":  examPaperListPage,
			"pageSize":   examPaperListPageSize,
			"syllabusId": examPaperListSyllabusID,
			"userId":     examPaperListUserID,
		}
		if err := c.PostAndDecode("/v1/paper/exam/list", body, &result); err != nil {
			return err
		}
		fmt.Printf("Total: %d\n\n", result.Total)
		headers := []string{"ID", "Name", "SyllabusId", "UserId"}
		rows := make([][]string, 0, len(result.List))
		for _, item := range result.List {
			rows = append(rows, []string{
				fmtFloat(item["id"]),
				fmtStr(item["name"]),
				fmtFloat(item["syllabusId"]),
				fmtFloat(item["userId"]),
			})
		}
		printTable(headers, rows)
		return nil
	},
}

// ---- get ----

var examPaperGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get an exam paper by ID (根据ID获取试卷)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid id: %s", args[0])
		}
		c := client.NewClient()
		var result interface{}
		if err := c.PostAndDecode("/v1/paper/exam/byId", map[string]interface{}{"id": id}, &result); err != nil {
			return err
		}
		prettyPrint(result)
		return nil
	},
}

// ---- create ----

var (
	examPaperCreateName        string
	examPaperCreateSyllabusID  uint
	examPaperCreateQuestionIDs string
)

var examPaperCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new exam paper (创建试卷)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if examPaperCreateName == "" {
			return fmt.Errorf("--name is required")
		}
		if examPaperCreateSyllabusID == 0 {
			return fmt.Errorf("--syllabus-id is required")
		}
		questionIDs, err := parseUintList(examPaperCreateQuestionIDs)
		if err != nil {
			return fmt.Errorf("invalid --question-ids: %w", err)
		}
		c := client.NewClient()
		body := map[string]interface{}{
			"name":        examPaperCreateName,
			"syllabusId":  examPaperCreateSyllabusID,
			"questionIds": questionIDs,
		}
		if err := c.PostAndDecode("/v1/paper/exam/create", body, nil); err != nil {
			return err
		}
		fmt.Println("Exam paper created successfully.")
		return nil
	},
}

// ---- update ----

var (
	examPaperUpdateID          uint
	examPaperUpdateName        string
	examPaperUpdateSyllabusID  uint
	examPaperUpdateQuestionIDs string
)

var examPaperUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an exam paper (修改试卷)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if examPaperUpdateID == 0 {
			return fmt.Errorf("--id is required")
		}
		body := map[string]interface{}{
			"id": examPaperUpdateID,
		}
		if examPaperUpdateName != "" {
			body["name"] = examPaperUpdateName
		}
		if examPaperUpdateSyllabusID != 0 {
			body["syllabusId"] = examPaperUpdateSyllabusID
		}
		if examPaperUpdateQuestionIDs != "" {
			questionIDs, err := parseUintList(examPaperUpdateQuestionIDs)
			if err != nil {
				return fmt.Errorf("invalid --question-ids: %w", err)
			}
			body["questionIds"] = questionIDs
		}
		c := client.NewClient()
		if err := c.PostAndDecode("/v1/paper/exam/edit", body, nil); err != nil {
			return err
		}
		fmt.Println("Exam paper updated successfully.")
		return nil
	},
}

// ---- delete ----

var examPaperDeleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete an exam paper by ID (删除试卷)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid id: %s", args[0])
		}
		c := client.NewClient()
		if err := c.PostAndDecode("/v1/paper/exam/delete", map[string]interface{}{"id": id}, nil); err != nil {
			return err
		}
		fmt.Printf("Exam paper %d deleted successfully.\n", id)
		return nil
	},
}

// ---- export ----

var (
	examPaperExportFormat string
	examPaperExportOutput string
)

var examPaperExportCmd = &cobra.Command{
	Use:   "export <id>",
	Short: "Export an exam paper to a file (导出试卷)",
	Long: `Export an exam paper to a file.

Supported formats:
  json  - Export as JSON (default)
  text  - Export as plain text`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid id: %s", args[0])
		}

		c := client.NewClient()
		var paper map[string]interface{}
		if err := c.PostAndDecode("/v1/paper/exam/byId", map[string]interface{}{"id": id}, &paper); err != nil {
			return err
		}

		format := strings.ToLower(examPaperExportFormat)
		if format == "" {
			format = "json"
		}

		var content []byte
		var defaultExt string

		switch format {
		case "json":
			content, err = json.MarshalIndent(paper, "", "  ")
			if err != nil {
				return fmt.Errorf("failed to marshal paper: %w", err)
			}
			defaultExt = ".json"
		case "text", "txt":
			content = []byte(formatExamPaperAsText(paper))
			defaultExt = ".txt"
		default:
			return fmt.Errorf("unsupported format %q; supported formats: json, text", format)
		}

		outFile := examPaperExportOutput
		if outFile == "" {
			outFile = fmt.Sprintf("exam_paper_%d%s", id, defaultExt)
		}

		if err := os.WriteFile(outFile, content, 0600); err != nil {
			return fmt.Errorf("failed to write file: %w", err)
		}
		fmt.Printf("Exam paper exported to %s\n", outFile)
		return nil
	},
}

// formatExamPaperAsText renders an exam paper as human-readable plain text.
func formatExamPaperAsText(paper map[string]interface{}) string {
	var sb strings.Builder

	sb.WriteString("========================================\n")
	name := fmtStr(paper["name"])
	if name == "" {
		name = "(untitled)"
	}
	sb.WriteString(fmt.Sprintf("Exam Paper: %s\n", name))
	sb.WriteString(fmt.Sprintf("ID:         %s\n", fmtFloat(paper["id"])))
	sb.WriteString(fmt.Sprintf("SyllabusId: %s\n", fmtFloat(paper["syllabusId"])))
	if syllabus, ok := paper["syllabus"].(map[string]interface{}); ok {
		sb.WriteString(fmt.Sprintf("Syllabus:   %s\n", fmtStr(syllabus["name"])))
	}
	sb.WriteString("========================================\n\n")

	questions, _ := paper["questions"].([]interface{})
	if len(questions) == 0 {
		sb.WriteString("(No questions)\n")
		return sb.String()
	}

	sb.WriteString(fmt.Sprintf("Questions (%d total):\n\n", len(questions)))
	for i, q := range questions {
		qMap, ok := q.(map[string]interface{})
		if !ok {
			continue
		}
		sb.WriteString(fmt.Sprintf("Q%d. [ID: %s]\n", i+1, fmtFloat(qMap["id"])))
		if stem := fmtStr(qMap["stem"]); stem != "" {
			sb.WriteString(fmt.Sprintf("    %s\n", stem))
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

// parseUintList parses a comma-separated list of unsigned integers.
func parseUintList(s string) ([]uint, error) {
	if s == "" {
		return nil, nil
	}
	parts := strings.Split(s, ",")
	result := make([]uint, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		n, err := strconv.ParseUint(p, 10, 0)
		if err != nil {
			return nil, fmt.Errorf("invalid integer %q", p)
		}
		result = append(result, uint(n))
	}
	return result, nil
}

func init() {
	// list flags
	examPaperListCmd.Flags().IntVar(&examPaperListPage, "page", 1, "Page number")
	examPaperListCmd.Flags().IntVar(&examPaperListPageSize, "page-size", 20, "Page size")
	examPaperListCmd.Flags().UintVar(&examPaperListSyllabusID, "syllabus-id", 0, "Filter by syllabus ID")
	examPaperListCmd.Flags().UintVar(&examPaperListUserID, "user-id", 0, "Filter by user ID")

	// create flags
	examPaperCreateCmd.Flags().StringVar(&examPaperCreateName, "name", "", "Exam paper name (required)")
	examPaperCreateCmd.Flags().UintVar(&examPaperCreateSyllabusID, "syllabus-id", 0, "Syllabus ID (required)")
	examPaperCreateCmd.Flags().StringVar(&examPaperCreateQuestionIDs, "question-ids", "", "Comma-separated list of question IDs (e.g. 1,2,3)")

	// update flags
	examPaperUpdateCmd.Flags().UintVar(&examPaperUpdateID, "id", 0, "Exam paper ID (required)")
	examPaperUpdateCmd.Flags().StringVar(&examPaperUpdateName, "name", "", "New name")
	examPaperUpdateCmd.Flags().UintVar(&examPaperUpdateSyllabusID, "syllabus-id", 0, "New syllabus ID")
	examPaperUpdateCmd.Flags().StringVar(&examPaperUpdateQuestionIDs, "question-ids", "", "New comma-separated list of question IDs (replaces existing)")

	// export flags
	examPaperExportCmd.Flags().StringVar(&examPaperExportFormat, "format", "json", "Export format: json, text")
	examPaperExportCmd.Flags().StringVar(&examPaperExportOutput, "output", "", "Output file path (default: exam_paper_<id>.<ext>)")

	examPaperCmd.AddCommand(examPaperListCmd)
	examPaperCmd.AddCommand(examPaperGetCmd)
	examPaperCmd.AddCommand(examPaperCreateCmd)
	examPaperCmd.AddCommand(examPaperUpdateCmd)
	examPaperCmd.AddCommand(examPaperDeleteCmd)
	examPaperCmd.AddCommand(examPaperExportCmd)
}
