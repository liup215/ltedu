package cmd

import (
	"edu/cli/client"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var classCmd = &cobra.Command{
	Use:   "class",
	Short: "Manage classes and students (班级 & 学生管理)",
}

// ---- list ----

var (
	classListPage     int
	classListPageSize int
)

var classListCmd = &cobra.Command{
	Use:   "list",
	Short: "List classes (班级列表)",
	RunE: func(cmd *cobra.Command, args []string) error {
		c := client.NewClient()
		var result struct {
			List  []map[string]interface{} `json:"list"`
			Total int                      `json:"total"`
		}
		body := map[string]interface{}{
			"pageIndex": classListPage,
			"pageSize":  classListPageSize,
		}
		if err := c.PostAndDecode("/v1/school/class/list", body, &result); err != nil {
			return err
		}
		fmt.Printf("Total: %d\n\n", result.Total)
		headers := []string{"ID", "Name", "Type", "InviteCode", "AdminUserId"}
		rows := make([][]string, 0, len(result.List))
		for _, item := range result.List {
			classTypeStr := "教学班"
			if t, ok := item["classType"].(float64); ok && int(t) == 2 {
				classTypeStr = "行政班"
			}
			rows = append(rows, []string{
				fmtFloat(item["id"]),
				fmtStr(item["name"]),
				classTypeStr,
				fmtStr(item["inviteCode"]),
				fmtFloat(item["adminUserId"]),
			})
		}
		printTable(headers, rows)
		return nil
	},
}

// ---- get ----

var classGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get a class by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid id: %s", args[0])
		}
		c := client.NewClient()
		var result interface{}
		if err := c.PostAndDecode("/v1/school/class/byId", map[string]interface{}{"id": id}, &result); err != nil {
			return err
		}
		prettyPrint(result)
		return nil
	},
}

// ---- create ----

var (
	classCreateName string
	classCreateType int
)

var classCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new class (教师权限，自动生成邀请码)",
	Long:  "Create a new class. --type 1 = 教学班 (default), --type 2 = 行政班 (each user may only belong to one)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if classCreateName == "" {
			return fmt.Errorf("--name is required")
		}
		if classCreateType != 1 && classCreateType != 2 {
			classCreateType = 1
		}
		c := client.NewClient()
		body := map[string]interface{}{
			"name":      classCreateName,
			"classType": classCreateType,
		}
		var result interface{}
		if err := c.PostAndDecode("/v1/school/class/create", body, &result); err != nil {
			return err
		}
		fmt.Println("Class created successfully.")
		prettyPrint(result)
		return nil
	},
}

// ---- edit ----

var (
	classEditID   uint
	classEditName string
)

var classEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a class",
	RunE: func(cmd *cobra.Command, args []string) error {
		if classEditID == 0 {
			return fmt.Errorf("--id is required")
		}
		c := client.NewClient()
		body := map[string]interface{}{
			"id":   classEditID,
			"name": classEditName,
		}
		if err := c.PostAndDecode("/v1/school/class/edit", body, nil); err != nil {
			return err
		}
		fmt.Println("Class updated successfully.")
		return nil
	},
}

// ---- delete ----

var classDeleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete a class by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid id: %s", args[0])
		}
		c := client.NewClient()
		if err := c.PostAndDecode("/v1/school/class/delete", map[string]interface{}{"id": id}, nil); err != nil {
			return err
		}
		fmt.Printf("Class %d deleted successfully.\n", id)
		return nil
	},
}

// ---- student subcommand ----

var studentCmd = &cobra.Command{
	Use:   "student",
	Short: "Manage students in a class (学生管理)",
}

var (
	studentListClassID uint
)

var studentListCmd = &cobra.Command{
	Use:   "list",
	Short: "List students in a class",
	RunE: func(cmd *cobra.Command, args []string) error {
		if studentListClassID == 0 {
			return fmt.Errorf("--class-id is required")
		}
		c := client.NewClient()
		var result struct {
			List  []map[string]interface{} `json:"list"`
			Total int                      `json:"total"`
		}
		body := map[string]interface{}{"id": studentListClassID}
		if err := c.PostAndDecode("/v1/school/class/studentList", body, &result); err != nil {
			return err
		}
		fmt.Printf("Total: %d\n\n", result.Total)
		headers := []string{"ID", "Username", "Realname", "Email"}
		rows := make([][]string, 0, len(result.List))
		for _, item := range result.List {
			rows = append(rows, []string{
				fmtFloat(item["id"]),
				fmtStr(item["username"]),
				fmtStr(item["realname"]),
				fmtStr(item["email"]),
			})
		}
		printTable(headers, rows)
		return nil
	},
}

var (
	studentAddClassID uint
	studentAddUserID  uint
)

var studentAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Directly add a student to a class (超级管理员专用)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if studentAddClassID == 0 || studentAddUserID == 0 {
			return fmt.Errorf("--class-id and --user-id are required")
		}
		c := client.NewClient()
		body := map[string]interface{}{
			"classId": studentAddClassID,
			"userId":  studentAddUserID,
		}
		if err := c.PostAndDecode("/v1/school/class/addStudent", body, nil); err != nil {
			return err
		}
		fmt.Printf("User %d added to class %d successfully.\n", studentAddUserID, studentAddClassID)
		return nil
	},
}

var (
	studentRemoveClassID uint
	studentRemoveUserID  uint
)

var studentRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a student from a class (管理员操作)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if studentRemoveClassID == 0 || studentRemoveUserID == 0 {
			return fmt.Errorf("--class-id and --user-id are required")
		}
		c := client.NewClient()
		body := map[string]interface{}{
			"classId": studentRemoveClassID,
			"userId":  studentRemoveUserID,
		}
		if err := c.PostAndDecode("/v1/school/class/removeStudent", body, nil); err != nil {
			return err
		}
		fmt.Printf("User %d removed from class %d successfully.\n", studentRemoveUserID, studentRemoveClassID)
		return nil
	},
}

// ---- apply subcommand (学生申请加入班级) ----

var (
	applyInviteCode string
	applyMessage    string
)

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply to join a class using an invite code (学生使用邀请码申请加入班级)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if applyInviteCode == "" {
			return fmt.Errorf("--invite-code is required")
		}
		c := client.NewClient()
		body := map[string]interface{}{
			"inviteCode": applyInviteCode,
			"message":    applyMessage,
		}
		var result interface{}
		if err := c.PostAndDecode("/v1/school/class/apply", body, &result); err != nil {
			return err
		}
		fmt.Println("Application submitted successfully. Waiting for admin approval.")
		prettyPrint(result)
		return nil
	},
}

// ---- request subcommand (管理员审核申请) ----

var requestCmd = &cobra.Command{
	Use:   "request",
	Short: "Manage join requests (管理员审核入班申请)",
}

var (
	requestListClassID uint
	requestListPage    int
	requestListSize    int
)

var requestListCmd = &cobra.Command{
	Use:   "list",
	Short: "List join requests for a class",
	RunE: func(cmd *cobra.Command, args []string) error {
		if requestListClassID == 0 {
			return fmt.Errorf("--class-id is required")
		}
		c := client.NewClient()
		var result struct {
			List  []map[string]interface{} `json:"list"`
			Total int                      `json:"total"`
		}
		pending := 0
		body := map[string]interface{}{
			"classId":   requestListClassID,
			"status":    &pending,
			"pageIndex": requestListPage,
			"pageSize":  requestListSize,
		}
		if err := c.PostAndDecode("/v1/school/class/joinRequest/list", body, &result); err != nil {
			return err
		}
		fmt.Printf("Total: %d\n\n", result.Total)
		headers := []string{"ID", "UserId", "Username", "Status", "Message"}
		rows := make([][]string, 0, len(result.List))
		for _, item := range result.List {
			statusStr := "pending"
			if s, ok := item["status"].(float64); ok {
				switch int(s) {
				case 1:
					statusStr = "approved"
				case 2:
					statusStr = "rejected"
				}
			}
			username := ""
			if u, ok := item["user"].(map[string]interface{}); ok {
				username = fmtStr(u["username"])
			}
			rows = append(rows, []string{
				fmtFloat(item["id"]),
				fmtFloat(item["userId"]),
				username,
				statusStr,
				fmtStr(item["message"]),
			})
		}
		printTable(headers, rows)
		return nil
	},
}

var requestApproveCmd = &cobra.Command{
	Use:   "approve <request-id>",
	Short: "Approve a join request",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid request id: %s", args[0])
		}
		c := client.NewClient()
		if err := c.PostAndDecode("/v1/school/class/joinRequest/approve", map[string]interface{}{"id": id}, nil); err != nil {
			return err
		}
		fmt.Printf("Request %d approved successfully.\n", id)
		return nil
	},
}

var requestRejectCmd = &cobra.Command{
	Use:   "reject <request-id>",
	Short: "Reject a join request",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid request id: %s", args[0])
		}
		c := client.NewClient()
		if err := c.PostAndDecode("/v1/school/class/joinRequest/reject", map[string]interface{}{"id": id}, nil); err != nil {
			return err
		}
		fmt.Printf("Request %d rejected successfully.\n", id)
		return nil
	},
}

func init() {
	classListCmd.Flags().IntVar(&classListPage, "page", 1, "Page number")
	classListCmd.Flags().IntVar(&classListPageSize, "page-size", 20, "Page size")

	classCreateCmd.Flags().StringVar(&classCreateName, "name", "", "Class name (required)")
	classCreateCmd.Flags().IntVar(&classCreateType, "type", 1, "Class type: 1=教学班 (default), 2=行政班")

	classEditCmd.Flags().UintVar(&classEditID, "id", 0, "Class ID (required)")
	classEditCmd.Flags().StringVar(&classEditName, "name", "", "Class name")

	studentListCmd.Flags().UintVar(&studentListClassID, "class-id", 0, "Class ID (required)")

	studentAddCmd.Flags().UintVar(&studentAddClassID, "class-id", 0, "Class ID (required)")
	studentAddCmd.Flags().UintVar(&studentAddUserID, "user-id", 0, "User ID (required)")

	studentRemoveCmd.Flags().UintVar(&studentRemoveClassID, "class-id", 0, "Class ID (required)")
	studentRemoveCmd.Flags().UintVar(&studentRemoveUserID, "user-id", 0, "User ID (required)")

	applyCmd.Flags().StringVar(&applyInviteCode, "invite-code", "", "Invite code (required)")
	applyCmd.Flags().StringVar(&applyMessage, "message", "", "Optional message to the admin")

	requestListCmd.Flags().UintVar(&requestListClassID, "class-id", 0, "Class ID (required)")
	requestListCmd.Flags().IntVar(&requestListPage, "page", 1, "Page number")
	requestListCmd.Flags().IntVar(&requestListSize, "page-size", 20, "Page size")

	studentCmd.AddCommand(studentListCmd)
	studentCmd.AddCommand(studentAddCmd)
	studentCmd.AddCommand(studentRemoveCmd)

	requestCmd.AddCommand(requestListCmd)
	requestCmd.AddCommand(requestApproveCmd)
	requestCmd.AddCommand(requestRejectCmd)

	classCmd.AddCommand(classListCmd)
	classCmd.AddCommand(classGetCmd)
	classCmd.AddCommand(classCreateCmd)
	classCmd.AddCommand(classEditCmd)
	classCmd.AddCommand(classDeleteCmd)
	classCmd.AddCommand(studentCmd)
	classCmd.AddCommand(applyCmd)
	classCmd.AddCommand(requestCmd)
}

