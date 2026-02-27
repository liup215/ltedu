package cmd

import (
	"edu/cli/client"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var teacherCmd = &cobra.Command{
	Use:   "teacher",
	Short: "Manage teacher-class relationships (教师班级管理)",
}

// ---- apply ----

var (
	teacherApplyClassID uint
	teacherApplyMessage string
)

var teacherApplyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply to join a class as a teacher (教师申请加入班级)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if teacherApplyClassID == 0 {
			return fmt.Errorf("--class-id is required")
		}
		c := client.NewClient()
		body := map[string]interface{}{
			"classId": teacherApplyClassID,
			"message": teacherApplyMessage,
		}
		var result interface{}
		if err := c.PostAndDecode("/v1/school/class/teacher/apply", body, &result); err != nil {
			return err
		}
		fmt.Println("Teacher application submitted successfully. Waiting for admin approval.")
		prettyPrint(result)
		return nil
	},
}

// ---- approve/reject ----

var teacherApproveCmd = &cobra.Command{
	Use:   "approve <application-id>",
	Short: "Approve a teacher join application (审核通过教师加入申请)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid application id: %s", args[0])
		}
		c := client.NewClient()
		if err := c.PostAndDecode("/v1/school/class/teacher/approve", map[string]interface{}{"id": id}, nil); err != nil {
			return err
		}
		fmt.Printf("Teacher application %d approved successfully.\n", id)
		return nil
	},
}

var teacherRejectCmd = &cobra.Command{
	Use:   "reject <application-id>",
	Short: "Reject a teacher join application (审核拒绝教师加入申请)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid application id: %s", args[0])
		}
		c := client.NewClient()
		if err := c.PostAndDecode("/v1/school/class/teacher/reject", map[string]interface{}{"id": id}, nil); err != nil {
			return err
		}
		fmt.Printf("Teacher application %d rejected.\n", id)
		return nil
	},
}

// ---- list applications ----

var (
	teacherListAppClassID uint
	teacherListAppPage    int
	teacherListAppSize    int
)

var teacherListApplicationsCmd = &cobra.Command{
	Use:   "list-applications",
	Short: "List teacher join applications for a class (班级教师加入申请列表)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if teacherListAppClassID == 0 {
			return fmt.Errorf("--class-id is required")
		}
		c := client.NewClient()
		var result struct {
			List  []map[string]interface{} `json:"list"`
			Total int                      `json:"total"`
		}
		pending := 0
		body := map[string]interface{}{
			"classId":   teacherListAppClassID,
			"status":    &pending,
			"pageIndex": teacherListAppPage,
			"pageSize":  teacherListAppSize,
		}
		if err := c.PostAndDecode("/v1/school/class/teacher/applications", body, &result); err != nil {
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

func init() {
	teacherApplyCmd.Flags().UintVar(&teacherApplyClassID, "class-id", 0, "Class ID (required)")
	teacherApplyCmd.Flags().StringVar(&teacherApplyMessage, "message", "", "Optional message to the admin")

	teacherListApplicationsCmd.Flags().UintVar(&teacherListAppClassID, "class-id", 0, "Class ID (required)")
	teacherListApplicationsCmd.Flags().IntVar(&teacherListAppPage, "page", 1, "Page number")
	teacherListApplicationsCmd.Flags().IntVar(&teacherListAppSize, "page-size", 20, "Page size")

	teacherCmd.AddCommand(teacherApplyCmd)
	teacherCmd.AddCommand(teacherApproveCmd)
	teacherCmd.AddCommand(teacherRejectCmd)
	teacherCmd.AddCommand(teacherListApplicationsCmd)
}
