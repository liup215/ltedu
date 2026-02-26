package cmd

import (
	"edu/cli/client"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var organisationCmd = &cobra.Command{
	Use:   "organisation",
	Short: "Manage organisations (机构管理)",
}

// ---- list ----

var (
	orgListPage     int
	orgListPageSize int
	orgListName     string
)

var orgListCmd = &cobra.Command{
	Use:   "list",
	Short: "List organisations",
	RunE: func(cmd *cobra.Command, args []string) error {
		c := client.NewClient()
		var result struct {
			List  []map[string]interface{} `json:"list"`
			Total int                      `json:"total"`
		}
		body := map[string]interface{}{
			"pageIndex": orgListPage,
			"pageSize":  orgListPageSize,
			"name":      orgListName,
		}
		if err := c.PostAndDecode("/v1/organisation/list", body, &result); err != nil {
			return err
		}
		fmt.Printf("Total: %d\n\n", result.Total)
		headers := []string{"ID", "Name"}
		rows := make([][]string, 0, len(result.List))
		for _, item := range result.List {
			rows = append(rows, []string{
				fmtFloat(item["id"]),
				fmtStr(item["name"]),
			})
		}
		printTable(headers, rows)
		return nil
	},
}

// ---- get ----

var orgGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get an organisation by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid id: %s", args[0])
		}
		c := client.NewClient()
		var result interface{}
		if err := c.PostAndDecode("/v1/organisation/byId", map[string]interface{}{"id": id}, &result); err != nil {
			return err
		}
		prettyPrint(result)
		return nil
	},
}

// ---- create ----

var orgCreateName string

var orgCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new organisation",
	RunE: func(cmd *cobra.Command, args []string) error {
		if orgCreateName == "" {
			return fmt.Errorf("--name is required")
		}
		c := client.NewClient()
		body := map[string]interface{}{
			"name": orgCreateName,
		}
		var result interface{}
		if err := c.PostAndDecode("/v1/organisation/create", body, &result); err != nil {
			return err
		}
		fmt.Println("Organisation created successfully.")
		prettyPrint(result)
		return nil
	},
}

// ---- edit ----

var (
	orgEditID   uint
	orgEditName string
)

var orgEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit an organisation",
	RunE: func(cmd *cobra.Command, args []string) error {
		if orgEditID == 0 {
			return fmt.Errorf("--id is required")
		}
		c := client.NewClient()
		body := map[string]interface{}{
			"id":   orgEditID,
			"name": orgEditName,
		}
		if err := c.PostAndDecode("/v1/organisation/edit", body, nil); err != nil {
			return err
		}
		fmt.Println("Organisation updated successfully.")
		return nil
	},
}

// ---- delete ----

var orgDeleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete an organisation by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid id: %s", args[0])
		}
		c := client.NewClient()
		if err := c.PostAndDecode("/v1/organisation/delete", map[string]interface{}{"id": id}, nil); err != nil {
			return err
		}
		fmt.Printf("Organisation %d deleted successfully.\n", id)
		return nil
	},
}

func init() {
	orgListCmd.Flags().IntVar(&orgListPage, "page", 1, "Page number")
	orgListCmd.Flags().IntVar(&orgListPageSize, "page-size", 20, "Page size")
	orgListCmd.Flags().StringVar(&orgListName, "name", "", "Filter by name")

	orgCreateCmd.Flags().StringVar(&orgCreateName, "name", "", "Organisation name (required)")

	orgEditCmd.Flags().UintVar(&orgEditID, "id", 0, "Organisation ID (required)")
	orgEditCmd.Flags().StringVar(&orgEditName, "name", "", "New name")

	organisationCmd.AddCommand(orgListCmd)
	organisationCmd.AddCommand(orgGetCmd)
	organisationCmd.AddCommand(orgCreateCmd)
	organisationCmd.AddCommand(orgEditCmd)
	organisationCmd.AddCommand(orgDeleteCmd)
}
