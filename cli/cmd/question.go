package cmd

import (
	"edu/cli/client"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var questionCmd = &cobra.Command{
	Use:   "question",
	Short: "Manage questions (题目管理)",
}

// ---- list ----

var (
	questionListPage              int
	questionListPageSize          int
	questionListSyllabusID        uint
	questionListStem              string
	questionListDifficult         int
	questionListStatus            int
	questionListPastPaperID       uint
	questionListChapterID         uint
	questionListKnowledgePointID  uint
	questionListKnowledgePointIDs string
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
			"chapterId":   questionListChapterID,
		}
		// Build knowledge point IDs list from --knowledge-point-id or --knowledge-point-ids
		kpIDs, err := parseUintList(questionListKnowledgePointIDs)
		if err != nil {
			return fmt.Errorf("invalid --knowledge-point-ids: %w", err)
		}
		if questionListKnowledgePointID != 0 {
			kpIDs = append(kpIDs, questionListKnowledgePointID)
		}
		if len(kpIDs) > 0 {
			body["knowledgePointIds"] = kpIDs
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

var questionGetWithKP bool

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
		var result map[string]interface{}
		if err := c.PostAndDecode("/v1/question/byId", map[string]interface{}{"id": id}, &result); err != nil {
			return err
		}
		if questionGetWithKP {
			fmt.Printf("Question ID: %s\n\n", fmtFloat(result["id"]))
			kps, _ := result["knowledgePoints"].([]interface{})
			if len(kps) == 0 {
				fmt.Println("No knowledge points linked to this question.")
			} else {
				fmt.Printf("Linked Knowledge Points (%d):\n\n", len(kps))
				headers := []string{"ID", "ChapterId", "Name", "Difficulty"}
				rows := make([][]string, 0, len(kps))
				for _, kp := range kps {
					if m, ok := kp.(map[string]interface{}); ok {
						rows = append(rows, []string{
							fmtFloat(m["id"]),
							fmtFloat(m["chapterId"]),
							fmtStr(m["name"]),
							fmtStr(m["difficulty"]),
						})
					}
				}
				printTable(headers, rows)
			}
		} else {
			prettyPrint(result)
		}
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

// ---- link-knowledge-point ----

var (
	questionLinkKPQuestionID        uint
	questionLinkKPKnowledgePointID  uint
	questionLinkKPKnowledgePointIDs string
)

var questionLinkKPCmd = &cobra.Command{
	Use:   "link-knowledge-point",
	Short: "Link knowledge point(s) to a question (requires --question-id and --knowledge-point-id or --knowledge-point-ids)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if questionLinkKPQuestionID == 0 {
			return fmt.Errorf("--question-id is required")
		}
		kpIDs, err := parseUintList(questionLinkKPKnowledgePointIDs)
		if err != nil {
			return fmt.Errorf("invalid --knowledge-point-ids: %w", err)
		}
		if questionLinkKPKnowledgePointID != 0 {
			kpIDs = append(kpIDs, questionLinkKPKnowledgePointID)
		}
		if len(kpIDs) == 0 {
			return fmt.Errorf("--knowledge-point-id or --knowledge-point-ids is required")
		}
		c := client.NewClient()
		for _, kpID := range kpIDs {
			body := map[string]interface{}{
				"questionId":       questionLinkKPQuestionID,
				"knowledgePointId": kpID,
			}
			if err := c.PostAndDecode("/v1/question/link-knowledge-point", body, nil); err != nil {
				return fmt.Errorf("failed to link knowledge point %d: %w", kpID, err)
			}
			fmt.Printf("Knowledge point %d linked to question %d successfully.\n", kpID, questionLinkKPQuestionID)
		}
		return nil
	},
}

// ---- unlink-knowledge-point ----

var (
	questionUnlinkKPQuestionID       uint
	questionUnlinkKPKnowledgePointID uint
)

var questionUnlinkKPCmd = &cobra.Command{
	Use:   "unlink-knowledge-point",
	Short: "Unlink a knowledge point from a question (requires --question-id and --knowledge-point-id)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if questionUnlinkKPQuestionID == 0 {
			return fmt.Errorf("--question-id is required")
		}
		if questionUnlinkKPKnowledgePointID == 0 {
			return fmt.Errorf("--knowledge-point-id is required")
		}
		c := client.NewClient()
		body := map[string]interface{}{
			"questionId":       questionUnlinkKPQuestionID,
			"knowledgePointId": questionUnlinkKPKnowledgePointID,
		}
		if err := c.PostAndDecode("/v1/question/unlink-knowledge-point", body, nil); err != nil {
			return err
		}
		fmt.Printf("Knowledge point %d unlinked from question %d successfully.\n", questionUnlinkKPKnowledgePointID, questionUnlinkKPQuestionID)
		return nil
	},
}

// ---- create ----

var (
	questionCreateSyllabusID        uint
	questionCreatePastPaperID       uint
	questionCreateIndexInPastPaper  int
	questionCreateStem              string
	questionCreateDifficult         int
	questionCreateContentsFile      string
	questionCreateKnowledgePointIDs string
)

var questionCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new question (requires --syllabus-id, --stem, --past-paper-id, --index-in-past-paper, and --contents-file)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if questionCreateSyllabusID == 0 {
			return fmt.Errorf("--syllabus-id is required")
		}
		if questionCreateStem == "" {
			return fmt.Errorf("--stem is required")
		}
		if questionCreatePastPaperID == 0 {
			return fmt.Errorf("--past-paper-id is required")
		}
		if questionCreateIndexInPastPaper == 0 {
			return fmt.Errorf("--index-in-past-paper is required")
		}
		if questionCreateContentsFile == "" {
			return fmt.Errorf("--contents-file is required")
		}

		// Read the contents file
		data, err := os.ReadFile(questionCreateContentsFile)
		if err != nil {
			return fmt.Errorf("failed to read contents file %q: %w", questionCreateContentsFile, err)
		}

		// Parse contents: accept either {"parts": [...]} or a plain array
		var questionContents []interface{}
		var wrapper struct {
			Parts []interface{} `json:"parts"`
		}
		if json.Unmarshal(data, &wrapper) == nil && wrapper.Parts != nil {
			questionContents = wrapper.Parts
		} else if err := json.Unmarshal(data, &questionContents); err != nil {
			return fmt.Errorf("contents file must be a JSON object with a \"parts\" array or a JSON array: %w", err)
		}

		difficult := questionCreateDifficult
		if difficult == 0 {
			difficult = 1
		}

		c := client.NewClient()

		// Build the question body
		q := map[string]interface{}{
			"syllabusId":       questionCreateSyllabusID,
			"stem":             questionCreateStem,
			"difficult":        difficult,
			"pastPaperId":      questionCreatePastPaperID,
			"indexInPastPaper": questionCreateIndexInPastPaper,
			"questionContents": questionContents,
		}

		var createResult struct {
			ID float64 `json:"id"`
		}
		if err := c.PostAndDecode("/v1/question/create", q, &createResult); err != nil {
			return fmt.Errorf("failed to create question: %w", err)
		}

		questionID := uint(createResult.ID)
		fmt.Printf("Question created successfully (ID: %d).\n", questionID)

		// Link knowledge points if provided
		kpIDs, err := parseUintList(questionCreateKnowledgePointIDs)
		if err != nil {
			return fmt.Errorf("invalid --knowledge-point-ids: %w", err)
		}
		for _, kpID := range kpIDs {
			body := map[string]interface{}{
				"questionId":       questionID,
				"knowledgePointId": kpID,
			}
			if err := c.PostAndDecode("/v1/question/link-knowledge-point", body, nil); err != nil {
				fmt.Fprintf(os.Stderr, "Warning: failed to link knowledge point %d: %v\n", kpID, err)
			} else {
				fmt.Printf("Knowledge point %d linked to question %d.\n", kpID, questionID)
			}
		}
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
	questionListCmd.Flags().UintVar(&questionListChapterID, "chapter-id", 0, "Filter by chapter ID")
	questionListCmd.Flags().UintVar(&questionListKnowledgePointID, "knowledge-point-id", 0, "Filter by a single knowledge point ID")
	questionListCmd.Flags().StringVar(&questionListKnowledgePointIDs, "knowledge-point-ids", "", "Filter by knowledge point IDs (comma-separated)")

	questionGetCmd.Flags().BoolVar(&questionGetWithKP, "with-knowledge-points", false, "Show linked knowledge points")

	questionLinkKPCmd.Flags().UintVar(&questionLinkKPQuestionID, "question-id", 0, "Question ID (required)")
	questionLinkKPCmd.Flags().UintVar(&questionLinkKPKnowledgePointID, "knowledge-point-id", 0, "Knowledge point ID to link")
	questionLinkKPCmd.Flags().StringVar(&questionLinkKPKnowledgePointIDs, "knowledge-point-ids", "", "Knowledge point IDs to link (comma-separated)")

	questionUnlinkKPCmd.Flags().UintVar(&questionUnlinkKPQuestionID, "question-id", 0, "Question ID (required)")
	questionUnlinkKPCmd.Flags().UintVar(&questionUnlinkKPKnowledgePointID, "knowledge-point-id", 0, "Knowledge point ID to unlink (required)")

	questionCreateCmd.Flags().UintVar(&questionCreateSyllabusID, "syllabus-id", 0, "Syllabus ID (required)")
	questionCreateCmd.Flags().UintVar(&questionCreatePastPaperID, "past-paper-id", 0, "Past paper ID (required)")
	questionCreateCmd.Flags().IntVar(&questionCreateIndexInPastPaper, "index-in-past-paper", 0, "Index of question within the past paper (required)")
	questionCreateCmd.Flags().StringVar(&questionCreateStem, "stem", "", "Question stem text (required)")
	questionCreateCmd.Flags().IntVar(&questionCreateDifficult, "difficult", 1, "Difficulty level 1-5 (default 1)")
	questionCreateCmd.Flags().StringVar(&questionCreateContentsFile, "contents-file", "", "Path to question contents JSON file (required)")
	questionCreateCmd.Flags().StringVar(&questionCreateKnowledgePointIDs, "knowledge-point-ids", "", "Knowledge point IDs to link (comma-separated, optional)")

	questionCmd.AddCommand(questionListCmd)
	questionCmd.AddCommand(questionGetCmd)
	questionCmd.AddCommand(questionDeleteCmd)
	questionCmd.AddCommand(questionLinkKPCmd)
	questionCmd.AddCommand(questionUnlinkKPCmd)
	questionCmd.AddCommand(questionCreateCmd)
}

