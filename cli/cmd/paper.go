package cmd

import (
	"edu/cli/client"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var paperCmd = &cobra.Command{
	Use:   "paper",
	Short: "Manage papers (试卷 & 真题管理)",
}

// ---- past paper subcommand ----

var pastPaperCmd = &cobra.Command{
	Use:   "past",
	Short: "Manage past papers (真题管理)",
}

var (
	pastPaperListPage       int
	pastPaperListPageSize   int
	pastPaperListSyllabusID uint
	pastPaperListYear       int
	pastPaperListCodeID     uint
	pastPaperListSeriesID   uint
)

var pastPaperListCmd = &cobra.Command{
	Use:   "list",
	Short: "List past papers",
	RunE: func(cmd *cobra.Command, args []string) error {
		c := client.NewClient()
		var result struct {
			List  []map[string]interface{} `json:"list"`
			Total int                      `json:"total"`
		}
		body := map[string]interface{}{
			"pageIndex":     pastPaperListPage,
			"pageSize":      pastPaperListPageSize,
			"syllabusId":    pastPaperListSyllabusID,
			"year":          pastPaperListYear,
			"paperCodeId":   pastPaperListCodeID,
			"paperSeriesId": pastPaperListSeriesID,
		}
		if err := c.PostAndDecode("/v1/paper/past/list", body, &result); err != nil {
			return err
		}
		fmt.Printf("Total: %d\n\n", result.Total)
		headers := []string{"ID", "Name", "Year", "SyllabusId", "PaperCodeId", "SeriesId"}
		rows := make([][]string, 0, len(result.List))
		for _, item := range result.List {
			rows = append(rows, []string{
				fmtFloat(item["id"]),
				fmtStr(item["name"]),
				fmtFloat(item["year"]),
				fmtFloat(item["syllabusId"]),
				fmtFloat(item["paperCodeId"]),
				fmtFloat(item["paperSeriesId"]),
			})
		}
		printTable(headers, rows)
		return nil
	},
}

var pastPaperGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get a past paper by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid id: %s", args[0])
		}
		c := client.NewClient()
		var result interface{}
		if err := c.PostAndDecode("/v1/paper/past/getById", map[string]interface{}{"id": id}, &result); err != nil {
			return err
		}
		prettyPrint(result)
		return nil
	},
}

// ---- paper code subcommand ----

var paperCodeCmd = &cobra.Command{
	Use:   "code",
	Short: "Manage paper codes",
}

var (
	paperCodeListPage       int
	paperCodeListPageSize   int
	paperCodeListSyllabusID uint
)

var paperCodeListCmd = &cobra.Command{
	Use:   "list",
	Short: "List paper codes",
	RunE: func(cmd *cobra.Command, args []string) error {
		c := client.NewClient()
		var result struct {
			List  []map[string]interface{} `json:"list"`
			Total int                      `json:"total"`
		}
		body := map[string]interface{}{
			"pageIndex":  paperCodeListPage,
			"pageSize":   paperCodeListPageSize,
			"syllabusId": paperCodeListSyllabusID,
		}
		if err := c.PostAndDecode("/v1/pastPaper/code/list", body, &result); err != nil {
			return err
		}
		fmt.Printf("Total: %d\n\n", result.Total)
		headers := []string{"ID", "Name", "ExamNodeId", "SyllabusId"}
		rows := make([][]string, 0, len(result.List))
		for _, item := range result.List {
			rows = append(rows, []string{
				fmtFloat(item["id"]),
				fmtStr(item["name"]),
				fmtFloat(item["examNodeId"]),
				fmtFloat(item["syllabusId"]),
			})
		}
		printTable(headers, rows)
		return nil
	},
}

var paperCodeGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get a paper code by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid id: %s", args[0])
		}
		c := client.NewClient()
		var result interface{}
		if err := c.PostAndDecode("/v1/pastPaper/code/getById", map[string]interface{}{"id": id}, &result); err != nil {
			return err
		}
		prettyPrint(result)
		return nil
	},
}

// ---- paper series subcommand ----

var paperSeriesCmd = &cobra.Command{
	Use:   "series",
	Short: "Manage paper series",
}

var (
	paperSeriesListPage       int
	paperSeriesListPageSize   int
	paperSeriesListSyllabusID uint
)

var paperSeriesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List paper series",
	RunE: func(cmd *cobra.Command, args []string) error {
		c := client.NewClient()
		var result struct {
			List  []map[string]interface{} `json:"list"`
			Total int                      `json:"total"`
		}
		body := map[string]interface{}{
			"pageIndex":  paperSeriesListPage,
			"pageSize":   paperSeriesListPageSize,
			"syllabusId": paperSeriesListSyllabusID,
		}
		if err := c.PostAndDecode("/v1/pastPaper/series/list", body, &result); err != nil {
			return err
		}
		fmt.Printf("Total: %d\n\n", result.Total)
		headers := []string{"ID", "Name", "SyllabusId"}
		rows := make([][]string, 0, len(result.List))
		for _, item := range result.List {
			rows = append(rows, []string{
				fmtFloat(item["id"]),
				fmtStr(item["name"]),
				fmtFloat(item["syllabusId"]),
			})
		}
		printTable(headers, rows)
		return nil
	},
}

func init() {
	// Past paper flags
	pastPaperListCmd.Flags().IntVar(&pastPaperListPage, "page", 1, "Page number")
	pastPaperListCmd.Flags().IntVar(&pastPaperListPageSize, "page-size", 20, "Page size")
	pastPaperListCmd.Flags().UintVar(&pastPaperListSyllabusID, "syllabus-id", 0, "Filter by syllabus ID")
	pastPaperListCmd.Flags().IntVar(&pastPaperListYear, "year", 0, "Filter by year")
	pastPaperListCmd.Flags().UintVar(&pastPaperListCodeID, "code-id", 0, "Filter by paper code ID")
	pastPaperListCmd.Flags().UintVar(&pastPaperListSeriesID, "series-id", 0, "Filter by paper series ID")

	pastPaperCmd.AddCommand(pastPaperListCmd)
	pastPaperCmd.AddCommand(pastPaperGetCmd)

	// Paper code flags
	paperCodeListCmd.Flags().IntVar(&paperCodeListPage, "page", 1, "Page number")
	paperCodeListCmd.Flags().IntVar(&paperCodeListPageSize, "page-size", 20, "Page size")
	paperCodeListCmd.Flags().UintVar(&paperCodeListSyllabusID, "syllabus-id", 0, "Filter by syllabus ID")

	paperCodeCmd.AddCommand(paperCodeListCmd)
	paperCodeCmd.AddCommand(paperCodeGetCmd)

	// Paper series flags
	paperSeriesListCmd.Flags().IntVar(&paperSeriesListPage, "page", 1, "Page number")
	paperSeriesListCmd.Flags().IntVar(&paperSeriesListPageSize, "page-size", 20, "Page size")
	paperSeriesListCmd.Flags().UintVar(&paperSeriesListSyllabusID, "syllabus-id", 0, "Filter by syllabus ID")

	paperSeriesCmd.AddCommand(paperSeriesListCmd)

	// Assemble paper command
	paperCmd.AddCommand(pastPaperCmd)
	paperCmd.AddCommand(paperCodeCmd)
	paperCmd.AddCommand(paperSeriesCmd)
}
