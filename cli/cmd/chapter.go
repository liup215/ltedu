package cmd

import (
	"edu/cli/client"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var chapterCmd = &cobra.Command{
	Use:   "chapter",
	Short: "Manage chapters (章节管理)",
}

// ---- list ----

var (
	chapterListPage       int
	chapterListPageSize   int
	chapterListSyllabusID uint
	chapterListParentID   uint
)

var chapterListCmd = &cobra.Command{
	Use:   "list",
	Short: "List chapters (requires --syllabus-id)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if chapterListSyllabusID == 0 {
			return fmt.Errorf("--syllabus-id is required")
		}
		c := client.NewClient()
		var result struct {
			List  []map[string]interface{} `json:"list"`
			Total int                      `json:"total"`
		}
		body := map[string]interface{}{
			"pageIndex":  chapterListPage,
			"pageSize":   chapterListPageSize,
			"syllabusId": chapterListSyllabusID,
			"parentId":   chapterListParentID,
		}
		if err := c.PostAndDecode("/v1/chapter/list", body, &result); err != nil {
			return err
		}
		fmt.Printf("Total: %d\n\n", result.Total)
		headers := []string{"ID", "Name", "SyllabusId", "ParentId", "IsLeaf"}
		rows := make([][]string, 0, len(result.List))
		for _, item := range result.List {
			rows = append(rows, []string{
				fmtFloat(item["id"]),
				fmtStr(item["name"]),
				fmtFloat(item["syllabusId"]),
				fmtFloat(item["parentId"]),
				fmtFloat(item["isLeaf"]),
			})
		}
		printTable(headers, rows)
		return nil
	},
}

// ---- tree ----

var chapterTreeSyllabusID uint

var chapterTreeCmd = &cobra.Command{
	Use:   "tree",
	Short: "Get chapter tree for a syllabus (requires --syllabus-id)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if chapterTreeSyllabusID == 0 {
			return fmt.Errorf("--syllabus-id is required")
		}
		c := client.NewClient()
		var result interface{}
		if err := c.PostAndDecode("/v1/chapter/tree", map[string]interface{}{"syllabusId": chapterTreeSyllabusID}, &result); err != nil {
			return err
		}
		prettyPrint(result)
		return nil
	},
}

// ---- get ----

var chapterGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get a chapter by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid id: %s", args[0])
		}
		c := client.NewClient()
		var result interface{}
		if err := c.PostAndDecode("/v1/chapter/byId", map[string]interface{}{"id": id}, &result); err != nil {
			return err
		}
		prettyPrint(result)
		return nil
	},
}

// ---- create ----

var (
	chapterCreateName       string
	chapterCreateSyllabusID uint
	chapterCreateParentID   uint
)

var chapterCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new chapter (requires --syllabus-id)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if chapterCreateName == "" {
			return fmt.Errorf("--name is required")
		}
		if chapterCreateSyllabusID == 0 {
			return fmt.Errorf("--syllabus-id is required")
		}
		c := client.NewClient()
		body := map[string]interface{}{
			"name":       chapterCreateName,
			"syllabusId": chapterCreateSyllabusID,
			"parentId":   chapterCreateParentID,
		}
		var result interface{}
		if err := c.PostAndDecode("/v1/chapter/create", body, &result); err != nil {
			return err
		}
		fmt.Println("Chapter created successfully.")
		prettyPrint(result)
		return nil
	},
}

// ---- edit ----

var (
	chapterEditID    uint
	chapterEditName  string
)

var chapterEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a chapter",
	RunE: func(cmd *cobra.Command, args []string) error {
		if chapterEditID == 0 {
			return fmt.Errorf("--id is required")
		}
		c := client.NewClient()
		body := map[string]interface{}{
			"id":   chapterEditID,
			"name": chapterEditName,
		}
		if err := c.PostAndDecode("/v1/chapter/edit", body, nil); err != nil {
			return err
		}
		fmt.Println("Chapter updated successfully.")
		return nil
	},
}

// ---- delete ----

var chapterDeleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete a chapter by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid id: %s", args[0])
		}
		c := client.NewClient()
		if err := c.PostAndDecode("/v1/chapter/delete", map[string]interface{}{"id": id}, nil); err != nil {
			return err
		}
		fmt.Printf("Chapter %d deleted successfully.\n", id)
		return nil
	},
}

func init() {
	chapterListCmd.Flags().IntVar(&chapterListPage, "page", 1, "Page number")
	chapterListCmd.Flags().IntVar(&chapterListPageSize, "page-size", 20, "Page size")
	chapterListCmd.Flags().UintVar(&chapterListSyllabusID, "syllabus-id", 0, "Syllabus ID (required)")
	chapterListCmd.Flags().UintVar(&chapterListParentID, "parent-id", 0, "Filter by parent chapter ID")

	chapterTreeCmd.Flags().UintVar(&chapterTreeSyllabusID, "syllabus-id", 0, "Syllabus ID (required)")

	chapterCreateCmd.Flags().StringVar(&chapterCreateName, "name", "", "Chapter name (required)")
	chapterCreateCmd.Flags().UintVar(&chapterCreateSyllabusID, "syllabus-id", 0, "Syllabus ID (required)")
	chapterCreateCmd.Flags().UintVar(&chapterCreateParentID, "parent-id", 0, "Parent chapter ID (0 = root)")

	chapterEditCmd.Flags().UintVar(&chapterEditID, "id", 0, "Chapter ID (required)")
	chapterEditCmd.Flags().StringVar(&chapterEditName, "name", "", "New name")

	chapterCmd.AddCommand(chapterListCmd)
	chapterCmd.AddCommand(chapterTreeCmd)
	chapterCmd.AddCommand(chapterGetCmd)
	chapterCmd.AddCommand(chapterCreateCmd)
	chapterCmd.AddCommand(chapterEditCmd)
	chapterCmd.AddCommand(chapterDeleteCmd)
}
