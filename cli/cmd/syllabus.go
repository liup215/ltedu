package cmd

import (
	"edu/cli/client"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var syllabusCmd = &cobra.Command{
	Use:   "syllabus",
	Short: "Manage syllabuses (考纲信息)",
}

// ---- list ----

var (
	syllabusListPage     int
	syllabusListPageSize int
	syllabusListOrgID    uint
)

var syllabusListCmd = &cobra.Command{
	Use:   "list",
	Short: "List syllabuses",
	RunE: func(cmd *cobra.Command, args []string) error {
		c := client.NewClient()
		var result struct {
			List  []map[string]interface{} `json:"list"`
			Total int                      `json:"total"`
		}
		body := map[string]interface{}{
			"pageIndex":      syllabusListPage,
			"pageSize":       syllabusListPageSize,
			"qualificationId": syllabusListOrgID,
		}
		if err := c.PostAndDecode("/v1/syllabus/list", body, &result); err != nil {
			return err
		}
		fmt.Printf("Total: %d\n\n", result.Total)
		headers := []string{"ID", "Name", "Code", "QualificationId"}
		rows := make([][]string, 0, len(result.List))
		for _, item := range result.List {
			rows = append(rows, []string{
				fmtFloat(item["id"]),
				fmtStr(item["name"]),
				fmtStr(item["code"]),
				fmtFloat(item["qualificationId"]),
			})
		}
		printTable(headers, rows)
		return nil
	},
}

// ---- get ----

var syllabusGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get a syllabus by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid id: %s", args[0])
		}
		c := client.NewClient()
		var result interface{}
		if err := c.PostAndDecode("/v1/syllabus/byId", map[string]interface{}{"id": id}, &result); err != nil {
			return err
		}
		prettyPrint(result)
		return nil
	},
}

func init() {
	syllabusListCmd.Flags().IntVar(&syllabusListPage, "page", 1, "Page number")
	syllabusListCmd.Flags().IntVar(&syllabusListPageSize, "page-size", 20, "Page size")
	syllabusListCmd.Flags().UintVar(&syllabusListOrgID, "qualification-id", 0, "Filter by qualification ID")

	syllabusCmd.AddCommand(syllabusListCmd)
	syllabusCmd.AddCommand(syllabusGetCmd)
}
