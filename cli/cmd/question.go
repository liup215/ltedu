package cmd

import (
	"edu/cli/client"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var questionCmd = &cobra.Command{
	Use:   "question",
	Short: "Manage questions (题目管理)",
}

// ---- list ----

var (
	questionListPage       int
	questionListPageSize   int
	questionListSyllabusID uint
	questionListStem       string
	questionListDifficult  int
	questionListStatus     int
	questionListPastPaperID uint
)

var questionListCmd = &cobra.Command{
	Use:   "list",
	Short: "List questions (requires --syllabus-id)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if questionListSyllabusID == 0 {
			return fmt.Errorf("--syllabus-id is required")
		}
		c := client.NewClient()
		var result struct {
			List  []map[string]interface{} `json:"list"`
			Total int                      `json:"total"`
		}
		body := map[string]interface{}{
			"pageIndex":   questionListPage,
			"pageSize":    questionListPageSize,
			"syllabusId":  questionListSyllabusID,
			"stem":        questionListStem,
			"difficult":   questionListDifficult,
			"Status":      questionListStatus,
			"pastPaperId": questionListPastPaperID,
		}
		if err := c.PostAndDecode("/v1/question/list", body, &result); err != nil {
			return err
		}
		fmt.Printf("Total: %d\n\n", result.Total)
		headers := []string{"ID", "SyllabusId", "Difficult", "Status", "PastPaperId", "TotalScore"}
		rows := make([][]string, 0, len(result.List))
		for _, item := range result.List {
			rows = append(rows, []string{
				fmtFloat(item["id"]),
				fmtFloat(item["syllabusId"]),
				fmtFloat(item["difficult"]),
				fmtFloat(item["status"]),
				fmtFloat(item["pastPaperId"]),
				fmtFloat(item["totalScore"]),
			})
		}
		printTable(headers, rows)
		return nil
	},
}

// ---- get ----

var questionGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get a question by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid id: %s", args[0])
		}
		c := client.NewClient()
		var result interface{}
		if err := c.PostAndDecode("/v1/question/byId", map[string]interface{}{"id": id}, &result); err != nil {
			return err
		}
		prettyPrint(result)
		return nil
	},
}

// ---- delete ----

var questionDeleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete a question by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid id: %s", args[0])
		}
		c := client.NewClient()
		if err := c.PostAndDecode("/v1/question/delete", map[string]interface{}{"id": id}, nil); err != nil {
			return err
		}
		fmt.Printf("Question %d deleted successfully.\n", id)
		return nil
	},
}

func init() {
	questionListCmd.Flags().IntVar(&questionListPage, "page", 1, "Page number")
	questionListCmd.Flags().IntVar(&questionListPageSize, "page-size", 20, "Page size")
	questionListCmd.Flags().UintVar(&questionListSyllabusID, "syllabus-id", 0, "Syllabus ID (required)")
	questionListCmd.Flags().StringVar(&questionListStem, "stem", "", "Filter by stem text")
	questionListCmd.Flags().IntVar(&questionListDifficult, "difficult", 0, "Filter by difficulty (1-5)")
	questionListCmd.Flags().IntVar(&questionListStatus, "status", 0, "Filter by status (1=Normal, 2=Forbidden, 3=Deleted)")
	questionListCmd.Flags().UintVar(&questionListPastPaperID, "past-paper-id", 0, "Filter by past paper ID")

	questionCmd.AddCommand(questionListCmd)
	questionCmd.AddCommand(questionGetCmd)
	questionCmd.AddCommand(questionDeleteCmd)
}
