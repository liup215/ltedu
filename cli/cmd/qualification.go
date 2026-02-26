package cmd

import (
	"edu/cli/client"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var qualificationCmd = &cobra.Command{
	Use:   "qualification",
	Short: "Manage qualifications (考试管理)",
}

// ---- list ----

var (
	qualListPage     int
	qualListPageSize int
	qualListOrgID    uint
)

var qualListCmd = &cobra.Command{
	Use:   "list",
	Short: "List qualifications",
	RunE: func(cmd *cobra.Command, args []string) error {
		c := client.NewClient()
		var result struct {
			List  []map[string]interface{} `json:"list"`
			Total int                      `json:"total"`
		}
		body := map[string]interface{}{
			"pageIndex":      qualListPage,
			"pageSize":       qualListPageSize,
			"organisationId": qualListOrgID,
		}
		if err := c.PostAndDecode("/v1/qualification/list", body, &result); err != nil {
			return err
		}
		fmt.Printf("Total: %d\n\n", result.Total)
		headers := []string{"ID", "Name", "OrganisationId"}
		rows := make([][]string, 0, len(result.List))
		for _, item := range result.List {
			rows = append(rows, []string{
				fmtFloat(item["id"]),
				fmtStr(item["name"]),
				fmtFloat(item["organisationId"]),
			})
		}
		printTable(headers, rows)
		return nil
	},
}

// ---- get ----

var qualGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get a qualification by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid id: %s", args[0])
		}
		c := client.NewClient()
		var result interface{}
		if err := c.PostAndDecode("/v1/qualification/byId", map[string]interface{}{"id": id}, &result); err != nil {
			return err
		}
		prettyPrint(result)
		return nil
	},
}

// ---- create ----

var (
	qualCreateName  string
	qualCreateOrgID uint
)

var qualCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new qualification",
	RunE: func(cmd *cobra.Command, args []string) error {
		if qualCreateName == "" {
			return fmt.Errorf("--name is required")
		}
		if qualCreateOrgID == 0 {
			return fmt.Errorf("--organisation-id is required")
		}
		c := client.NewClient()
		body := map[string]interface{}{
			"name":           qualCreateName,
			"organisationId": qualCreateOrgID,
		}
		var result interface{}
		if err := c.PostAndDecode("/v1/qualification/create", body, &result); err != nil {
			return err
		}
		fmt.Println("Qualification created successfully.")
		prettyPrint(result)
		return nil
	},
}

// ---- edit ----

var (
	qualEditID    uint
	qualEditName  string
	qualEditOrgID uint
)

var qualEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a qualification",
	RunE: func(cmd *cobra.Command, args []string) error {
		if qualEditID == 0 {
			return fmt.Errorf("--id is required")
		}
		c := client.NewClient()
		body := map[string]interface{}{
			"id":             qualEditID,
			"name":           qualEditName,
			"organisationId": qualEditOrgID,
		}
		if err := c.PostAndDecode("/v1/qualification/edit", body, nil); err != nil {
			return err
		}
		fmt.Println("Qualification updated successfully.")
		return nil
	},
}

// ---- delete ----

var qualDeleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete a qualification by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid id: %s", args[0])
		}
		c := client.NewClient()
		if err := c.PostAndDecode("/v1/qualification/delete", map[string]interface{}{"id": id}, nil); err != nil {
			return err
		}
		fmt.Printf("Qualification %d deleted successfully.\n", id)
		return nil
	},
}

func init() {
	qualListCmd.Flags().IntVar(&qualListPage, "page", 1, "Page number")
	qualListCmd.Flags().IntVar(&qualListPageSize, "page-size", 20, "Page size")
	qualListCmd.Flags().UintVar(&qualListOrgID, "organisation-id", 0, "Filter by organisation ID")

	qualCreateCmd.Flags().StringVar(&qualCreateName, "name", "", "Qualification name (required)")
	qualCreateCmd.Flags().UintVar(&qualCreateOrgID, "organisation-id", 0, "Organisation ID (required)")

	qualEditCmd.Flags().UintVar(&qualEditID, "id", 0, "Qualification ID (required)")
	qualEditCmd.Flags().StringVar(&qualEditName, "name", "", "New name")
	qualEditCmd.Flags().UintVar(&qualEditOrgID, "organisation-id", 0, "New organisation ID")

	qualificationCmd.AddCommand(qualListCmd)
	qualificationCmd.AddCommand(qualGetCmd)
	qualificationCmd.AddCommand(qualCreateCmd)
	qualificationCmd.AddCommand(qualEditCmd)
	qualificationCmd.AddCommand(qualDeleteCmd)
}
