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
	classListGradeID  uint
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
			"gradeId":   classListGradeID,
		}
		if err := c.PostAndDecode("/v1/school/class/list", body, &result); err != nil {
			return err
		}
		fmt.Printf("Total: %d\n\n", result.Total)
		headers := []string{"ID", "Name", "GradeId"}
		rows := make([][]string, 0, len(result.List))
		for _, item := range result.List {
			rows = append(rows, []string{
				fmtFloat(item["id"]),
				fmtStr(item["name"]),
				fmtFloat(item["gradeId"]),
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
	classCreateName    string
	classCreateGradeID uint
)

var classCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new class",
	RunE: func(cmd *cobra.Command, args []string) error {
		if classCreateName == "" {
			return fmt.Errorf("--name is required")
		}
		c := client.NewClient()
		body := map[string]interface{}{
			"name":    classCreateName,
			"gradeId": classCreateGradeID,
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
	classEditID      uint
	classEditName    string
	classEditGradeID uint
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
			"id":      classEditID,
			"name":    classEditName,
			"gradeId": classEditGradeID,
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
	Short: "Add a student to a class",
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
	Short: "Remove a student from a class",
	RunE: func(cmd *cobra.Command, args []string) error {
		if studentRemoveClassID == 0 || studentRemoveUserID == 0 {
			return fmt.Errorf("--class-id and --user-id are required")
		}
		c := client.NewClient()
		body := map[string]interface{}{
			"classId": studentRemoveClassID,
			"userId":  studentRemoveUserID,
		}
		if err := c.PostAndDecode("/v1/school/class/deleteStudent", body, nil); err != nil {
			return err
		}
		fmt.Printf("User %d removed from class %d successfully.\n", studentRemoveUserID, studentRemoveClassID)
		return nil
	},
}

func init() {
	classListCmd.Flags().IntVar(&classListPage, "page", 1, "Page number")
	classListCmd.Flags().IntVar(&classListPageSize, "page-size", 20, "Page size")
	classListCmd.Flags().UintVar(&classListGradeID, "grade-id", 0, "Filter by grade ID")

	classCreateCmd.Flags().StringVar(&classCreateName, "name", "", "Class name (required)")
	classCreateCmd.Flags().UintVar(&classCreateGradeID, "grade-id", 0, "Grade ID")

	classEditCmd.Flags().UintVar(&classEditID, "id", 0, "Class ID (required)")
	classEditCmd.Flags().StringVar(&classEditName, "name", "", "Class name")
	classEditCmd.Flags().UintVar(&classEditGradeID, "grade-id", 0, "Grade ID")

	studentListCmd.Flags().UintVar(&studentListClassID, "class-id", 0, "Class ID (required)")

	studentAddCmd.Flags().UintVar(&studentAddClassID, "class-id", 0, "Class ID (required)")
	studentAddCmd.Flags().UintVar(&studentAddUserID, "user-id", 0, "User ID (required)")

	studentRemoveCmd.Flags().UintVar(&studentRemoveClassID, "class-id", 0, "Class ID (required)")
	studentRemoveCmd.Flags().UintVar(&studentRemoveUserID, "user-id", 0, "User ID (required)")

	studentCmd.AddCommand(studentListCmd)
	studentCmd.AddCommand(studentAddCmd)
	studentCmd.AddCommand(studentRemoveCmd)

	classCmd.AddCommand(classListCmd)
	classCmd.AddCommand(classGetCmd)
	classCmd.AddCommand(classCreateCmd)
	classCmd.AddCommand(classEditCmd)
	classCmd.AddCommand(classDeleteCmd)
	classCmd.AddCommand(studentCmd)
}
