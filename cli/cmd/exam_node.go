package cmd

import (
	"edu/cli/client"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var examNodeCmd = &cobra.Command{
	Use:   "exam-node",
	Short: "Manage syllabus exam nodes (考试节点管理)",
}

// ---- list ----

var examNodeListSyllabusID uint

var examNodeListCmd = &cobra.Command{
	Use:   "list",
	Short: "List exam nodes for a syllabus (列出考纲的考试节点)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if examNodeListSyllabusID == 0 {
			return fmt.Errorf("--syllabus-id is required")
		}
		c := client.NewClient()
		var result []map[string]interface{}
		if err := c.PostAndDecode("/v1/syllabus/examNode/list", map[string]interface{}{"syllabusId": examNodeListSyllabusID}, &result); err != nil {
			return err
		}
		fmt.Printf("Total: %d\n\n", len(result))
		headers := []string{"ID", "Name", "SortOrder", "Chapters", "PaperCodes"}
		rows := make([][]string, 0, len(result))
		for _, item := range result {
			chapCount := 0
			if chs, ok := item["chapters"].([]interface{}); ok {
				chapCount = len(chs)
			}
			pcCount := 0
			if pcs, ok := item["paperCodes"].([]interface{}); ok {
				pcCount = len(pcs)
			}
			rows = append(rows, []string{
				fmtFloat(item["id"]),
				fmtStr(item["name"]),
				fmtFloat(item["sortOrder"]),
				strconv.Itoa(chapCount),
				strconv.Itoa(pcCount),
			})
		}
		printTable(headers, rows)
		return nil
	},
}

// ---- get ----

var examNodeGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get an exam node by ID (获取考试节点详情)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid id: %s", args[0])
		}
		c := client.NewClient()
		var result interface{}
		if err := c.PostAndDecode("/v1/syllabus/examNode/byId", map[string]interface{}{"id": id}, &result); err != nil {
			return err
		}
		prettyPrint(result)
		return nil
	},
}

// ---- create ----

var (
	examNodeCreateSyllabusID  uint
	examNodeCreateName        string
	examNodeCreateDescription string
	examNodeCreateSortOrder   int
)

var examNodeCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an exam node for a syllabus (为考纲创建考试节点)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if examNodeCreateSyllabusID == 0 || examNodeCreateName == "" {
			return fmt.Errorf("--syllabus-id and --name are required")
		}
		c := client.NewClient()
		body := map[string]interface{}{
			"syllabusId":  examNodeCreateSyllabusID,
			"name":        examNodeCreateName,
			"description": examNodeCreateDescription,
			"sortOrder":   examNodeCreateSortOrder,
		}
		var result interface{}
		if err := c.PostAndDecode("/v1/syllabus/examNode/create", body, &result); err != nil {
			return err
		}
		fmt.Println("Exam node created successfully.")
		prettyPrint(result)
		return nil
	},
}

// ---- edit ----

var (
	examNodeEditID          uint
	examNodeEditName        string
	examNodeEditDescription string
	examNodeEditSortOrder   int
)

var examNodeEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Update an exam node (更新考试节点信息)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if examNodeEditID == 0 {
			return fmt.Errorf("--id is required")
		}
		c := client.NewClient()
		body := map[string]interface{}{
			"id":          examNodeEditID,
			"name":        examNodeEditName,
			"description": examNodeEditDescription,
			"sortOrder":   examNodeEditSortOrder,
		}
		var result interface{}
		if err := c.PostAndDecode("/v1/syllabus/examNode/edit", body, &result); err != nil {
			return err
		}
		fmt.Println("Exam node updated successfully.")
		prettyPrint(result)
		return nil
	},
}

// ---- delete ----

var examNodeDeleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete an exam node by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid id: %s", args[0])
		}
		c := client.NewClient()
		if err := c.PostAndDecode("/v1/syllabus/examNode/delete", map[string]interface{}{"id": id}, nil); err != nil {
			return err
		}
		fmt.Printf("Exam node %d deleted successfully.\n", id)
		return nil
	},
}

// ---- add-chapter ----

var (
	addChapterExamNodeID uint
	addChapterChapterID  uint
)

var examNodeAddChapterCmd = &cobra.Command{
	Use:   "add-chapter",
	Short: "Add a chapter (and all its sub-chapters) to an exam node (添加章节及其所有子章节到考试节点)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if addChapterExamNodeID == 0 || addChapterChapterID == 0 {
			return fmt.Errorf("--exam-node-id and --chapter-id are required")
		}
		c := client.NewClient()
		body := map[string]interface{}{
			"examNodeId": addChapterExamNodeID,
			"chapterId":  addChapterChapterID,
		}
		if err := c.PostAndDecode("/v1/syllabus/examNode/chapter/add", body, nil); err != nil {
			return err
		}
		fmt.Printf("Chapter %d (and sub-chapters) added to exam node %d successfully.\n", addChapterChapterID, addChapterExamNodeID)
		return nil
	},
}

// ---- remove-chapter ----

var (
	removeChapterExamNodeID uint
	removeChapterChapterID  uint
)

var examNodeRemoveChapterCmd = &cobra.Command{
	Use:   "remove-chapter",
	Short: "Remove a chapter from an exam node (从考试节点移除章节)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if removeChapterExamNodeID == 0 || removeChapterChapterID == 0 {
			return fmt.Errorf("--exam-node-id and --chapter-id are required")
		}
		c := client.NewClient()
		body := map[string]interface{}{
			"examNodeId": removeChapterExamNodeID,
			"chapterId":  removeChapterChapterID,
		}
		if err := c.PostAndDecode("/v1/syllabus/examNode/chapter/remove", body, nil); err != nil {
			return err
		}
		fmt.Printf("Chapter %d removed from exam node %d successfully.\n", removeChapterChapterID, removeChapterExamNodeID)
		return nil
	},
}

// ---- add-paper-code ----

var (
	addPaperCodeExamNodeID  uint
	addPaperCodePaperCodeID uint
)

var examNodeAddPaperCodeCmd = &cobra.Command{
	Use:   "add-paper-code",
	Short: "Add a paper code to an exam node (为考试节点添加试卷代码)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if addPaperCodeExamNodeID == 0 || addPaperCodePaperCodeID == 0 {
			return fmt.Errorf("--exam-node-id and --paper-code-id are required")
		}
		c := client.NewClient()
		body := map[string]interface{}{
			"examNodeId":  addPaperCodeExamNodeID,
			"paperCodeId": addPaperCodePaperCodeID,
		}
		if err := c.PostAndDecode("/v1/syllabus/examNode/paperCode/add", body, nil); err != nil {
			return err
		}
		fmt.Printf("Paper code %d added to exam node %d successfully.\n", addPaperCodePaperCodeID, addPaperCodeExamNodeID)
		return nil
	},
}

// ---- remove-paper-code ----

var (
	removePaperCodeExamNodeID  uint
	removePaperCodePaperCodeID uint
)

var examNodeRemovePaperCodeCmd = &cobra.Command{
	Use:   "remove-paper-code",
	Short: "Remove a paper code from an exam node (从考试节点移除试卷代码)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if removePaperCodeExamNodeID == 0 || removePaperCodePaperCodeID == 0 {
			return fmt.Errorf("--exam-node-id and --paper-code-id are required")
		}
		c := client.NewClient()
		body := map[string]interface{}{
			"examNodeId":  removePaperCodeExamNodeID,
			"paperCodeId": removePaperCodePaperCodeID,
		}
		if err := c.PostAndDecode("/v1/syllabus/examNode/paperCode/remove", body, nil); err != nil {
			return err
		}
		fmt.Printf("Paper code %d removed from exam node %d successfully.\n", removePaperCodePaperCodeID, removePaperCodeExamNodeID)
		return nil
	},
}

func init() {
	examNodeListCmd.Flags().UintVar(&examNodeListSyllabusID, "syllabus-id", 0, "Syllabus ID (required)")

	examNodeCreateCmd.Flags().UintVar(&examNodeCreateSyllabusID, "syllabus-id", 0, "Syllabus ID (required)")
	examNodeCreateCmd.Flags().StringVar(&examNodeCreateName, "name", "", "Exam node name (required)")
	examNodeCreateCmd.Flags().StringVar(&examNodeCreateDescription, "description", "", "Description")
	examNodeCreateCmd.Flags().IntVar(&examNodeCreateSortOrder, "sort-order", 0, "Sort order")

	examNodeEditCmd.Flags().UintVar(&examNodeEditID, "id", 0, "Exam node ID (required)")
	examNodeEditCmd.Flags().StringVar(&examNodeEditName, "name", "", "New name")
	examNodeEditCmd.Flags().StringVar(&examNodeEditDescription, "description", "", "New description")
	examNodeEditCmd.Flags().IntVar(&examNodeEditSortOrder, "sort-order", 0, "New sort order")

	examNodeAddChapterCmd.Flags().UintVar(&addChapterExamNodeID, "exam-node-id", 0, "Exam node ID (required)")
	examNodeAddChapterCmd.Flags().UintVar(&addChapterChapterID, "chapter-id", 0, "Chapter ID (required); all sub-chapters are included automatically")

	examNodeRemoveChapterCmd.Flags().UintVar(&removeChapterExamNodeID, "exam-node-id", 0, "Exam node ID (required)")
	examNodeRemoveChapterCmd.Flags().UintVar(&removeChapterChapterID, "chapter-id", 0, "Chapter ID (required)")

	examNodeAddPaperCodeCmd.Flags().UintVar(&addPaperCodeExamNodeID, "exam-node-id", 0, "Exam node ID (required)")
	examNodeAddPaperCodeCmd.Flags().UintVar(&addPaperCodePaperCodeID, "paper-code-id", 0, "Paper code ID (required)")

	examNodeRemovePaperCodeCmd.Flags().UintVar(&removePaperCodeExamNodeID, "exam-node-id", 0, "Exam node ID (required)")
	examNodeRemovePaperCodeCmd.Flags().UintVar(&removePaperCodePaperCodeID, "paper-code-id", 0, "Paper code ID (required)")

	examNodeCmd.AddCommand(examNodeListCmd)
	examNodeCmd.AddCommand(examNodeGetCmd)
	examNodeCmd.AddCommand(examNodeCreateCmd)
	examNodeCmd.AddCommand(examNodeEditCmd)
	examNodeCmd.AddCommand(examNodeDeleteCmd)
	examNodeCmd.AddCommand(examNodeAddChapterCmd)
	examNodeCmd.AddCommand(examNodeRemoveChapterCmd)
	examNodeCmd.AddCommand(examNodeAddPaperCodeCmd)
	examNodeCmd.AddCommand(examNodeRemovePaperCodeCmd)
}
