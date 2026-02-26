package cmd

import (
	"edu/cli/client"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Manage users (用户管理)",
}

// ---- list ----

var (
	userListPage     int
	userListPageSize int
	userListUsername string
	userListRealname string
	userListStatus   int
	userListClassID  uint
)

var userListCmd = &cobra.Command{
	Use:   "list",
	Short: "List users",
	RunE: func(cmd *cobra.Command, args []string) error {
		c := client.NewClient()
		var result struct {
			List  []map[string]interface{} `json:"list"`
			Total int                      `json:"total"`
		}
		body := map[string]interface{}{
			"pageIndex": userListPage,
			"pageSize":  userListPageSize,
			"username":  userListUsername,
			"realname":  userListRealname,
			"status":    userListStatus,
			"classId":   userListClassID,
		}
		if err := c.PostAndDecode("/v1/user/list", body, &result); err != nil {
			return err
		}
		fmt.Printf("Total: %d\n\n", result.Total)
		headers := []string{"ID", "Username", "Realname", "Email", "Mobile", "Status", "IsAdmin"}
		rows := make([][]string, 0, len(result.List))
		for _, item := range result.List {
			rows = append(rows, []string{
				fmtFloat(item["id"]),
				fmtStr(item["username"]),
				fmtStr(item["realname"]),
				fmtStr(item["email"]),
				fmtStr(item["mobile"]),
				fmtFloat(item["status"]),
				fmtBool(item["isAdmin"]),
			})
		}
		printTable(headers, rows)
		return nil
	},
}

// ---- get ----

var userGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get a user by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid id: %s", args[0])
		}
		c := client.NewClient()
		var result interface{}
		if err := c.PostAndDecode("/v1/user/byId", map[string]interface{}{"id": id}, &result); err != nil {
			return err
		}
		prettyPrint(result)
		return nil
	},
}

// ---- create ----

var (
	userCreateUsername string
	userCreateRealname string
	userCreateEmail    string
	userCreateMobile   string
)

var userCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new user (default password: 123456)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if userCreateUsername == "" {
			return fmt.Errorf("--username is required")
		}
		c := client.NewClient()
		body := map[string]interface{}{
			"username": userCreateUsername,
			"realname": userCreateRealname,
			"email":    userCreateEmail,
			"mobile":   userCreateMobile,
		}
		var result interface{}
		if err := c.PostAndDecode("/v1/user/create", body, &result); err != nil {
			return err
		}
		fmt.Println("User created successfully.")
		prettyPrint(result)
		return nil
	},
}

// ---- edit ----

var (
	userEditID       uint
	userEditRealname string
	userEditNickname string
	userEditEngname  string
	userEditSex      uint
	userEditStatus   int
)

var userEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a user",
	RunE: func(cmd *cobra.Command, args []string) error {
		if userEditID == 0 {
			return fmt.Errorf("--id is required")
		}
		c := client.NewClient()
		body := map[string]interface{}{
			"id":       userEditID,
			"realname": userEditRealname,
			"nickname": userEditNickname,
			"engname":  userEditEngname,
			"sex":      userEditSex,
			"status":   userEditStatus,
		}
		if err := c.PostAndDecode("/v1/user/edit", body, nil); err != nil {
			return err
		}
		fmt.Println("User updated successfully.")
		return nil
	},
}

// ---- delete ----

var userDeleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete a user by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid id: %s", args[0])
		}
		c := client.NewClient()
		if err := c.PostAndDecode("/v1/user/delete", map[string]interface{}{"id": id}, nil); err != nil {
			return err
		}
		fmt.Printf("User %d deleted successfully.\n", id)
		return nil
	},
}

func init() {
	userListCmd.Flags().IntVar(&userListPage, "page", 1, "Page number")
	userListCmd.Flags().IntVar(&userListPageSize, "page-size", 20, "Page size")
	userListCmd.Flags().StringVar(&userListUsername, "username", "", "Filter by username")
	userListCmd.Flags().StringVar(&userListRealname, "realname", "", "Filter by realname")
	userListCmd.Flags().IntVar(&userListStatus, "status", 0, "Filter by status")
	userListCmd.Flags().UintVar(&userListClassID, "class-id", 0, "Filter by class ID")

	userCreateCmd.Flags().StringVar(&userCreateUsername, "username", "", "Username (required)")
	userCreateCmd.Flags().StringVar(&userCreateRealname, "realname", "", "Real name")
	userCreateCmd.Flags().StringVar(&userCreateEmail, "email", "", "Email address")
	userCreateCmd.Flags().StringVar(&userCreateMobile, "mobile", "", "Mobile number")

	userEditCmd.Flags().UintVar(&userEditID, "id", 0, "User ID (required)")
	userEditCmd.Flags().StringVar(&userEditRealname, "realname", "", "Real name")
	userEditCmd.Flags().StringVar(&userEditNickname, "nickname", "", "Nickname")
	userEditCmd.Flags().StringVar(&userEditEngname, "engname", "", "English name")
	userEditCmd.Flags().UintVar(&userEditSex, "sex", 0, "Sex (1=Male, 2=Female)")
	userEditCmd.Flags().IntVar(&userEditStatus, "status", 0, "Status (1=Active, 2=Inactive, 3=Suspended, 4=Banned)")

	userCmd.AddCommand(userListCmd)
	userCmd.AddCommand(userGetCmd)
	userCmd.AddCommand(userCreateCmd)
	userCmd.AddCommand(userEditCmd)
	userCmd.AddCommand(userDeleteCmd)
}
