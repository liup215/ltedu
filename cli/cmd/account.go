package cmd

import (
	"edu/cli/client"
	"fmt"

	"github.com/spf13/cobra"
)

var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Manage your personal account (个人账户管理)",
}

// ---- info ----

var accountInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Display current account information",
	RunE: func(cmd *cobra.Command, args []string) error {
		c := client.NewClient()
		var result map[string]interface{}
		if err := c.GetAndDecode("/v1/user", &result); err != nil {
			return err
		}

		fmt.Printf("ID       : %s\n", fmtFloat(result["id"]))
		fmt.Printf("Username : %s\n", fmtStr(result["username"]))
		fmt.Printf("Realname : %s\n", fmtStr(result["realname"]))
		fmt.Printf("Nickname : %s\n", fmtStr(result["nickname"]))
		fmt.Printf("Engname  : %s\n", fmtStr(result["engname"]))
		fmt.Printf("Email    : %s\n", fmtStr(result["email"]))
		fmt.Printf("Mobile   : %s\n", fmtStr(result["mobile"]))
		fmt.Printf("Sex      : %s\n", fmtSex(result["sex"]))
		fmt.Printf("Status   : %s\n", fmtUserStatus(result["status"]))
		fmt.Printf("IsAdmin  : %s\n", fmtBool(result["isAdmin"]))
		fmt.Printf("IsTeacher: %s\n", fmtBool(result["isTeacher"]))
		if result["vipExpireAt"] != nil && fmtStr(result["vipExpireAt"]) != "" {
			fmt.Printf("VIP Exp  : %s\n", fmtStr(result["vipExpireAt"]))
		}
		return nil
	},
}

// ---- update ----

var (
	accountUpdateRealname string
	accountUpdateNickname string
	accountUpdateEngname  string
	accountUpdateSex      uint
	accountUpdateMobile   string
)

var accountUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update personal account details (realname, nickname, engname, sex, mobile)",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Require at least one flag to be explicitly set
		if !cmd.Flags().Changed("realname") &&
			!cmd.Flags().Changed("nickname") &&
			!cmd.Flags().Changed("engname") &&
			!cmd.Flags().Changed("sex") &&
			!cmd.Flags().Changed("mobile") {
			return fmt.Errorf("at least one field flag must be specified (--realname, --nickname, --engname, --sex, --mobile)")
		}

		c := client.NewClient()

		// Fetch current profile to pre-populate unchanged fields
		var current map[string]interface{}
		if err := c.GetAndDecode("/v1/user", &current); err != nil {
			return fmt.Errorf("failed to fetch current account info: %w", err)
		}

		realname := fmtStr(current["realname"])
		nickname := fmtStr(current["nickname"])
		engname := fmtStr(current["engname"])
		mobile := fmtStr(current["mobile"])
		sex := uint(0)
		if s, ok := current["sex"].(float64); ok {
			sex = uint(s)
		}

		if cmd.Flags().Changed("realname") {
			realname = accountUpdateRealname
		}
		if cmd.Flags().Changed("nickname") {
			nickname = accountUpdateNickname
		}
		if cmd.Flags().Changed("engname") {
			engname = accountUpdateEngname
		}
		if cmd.Flags().Changed("sex") {
			sex = accountUpdateSex
		}
		if cmd.Flags().Changed("mobile") {
			mobile = accountUpdateMobile
		}

		body := map[string]interface{}{
			"realname": realname,
			"nickname": nickname,
			"engname":  engname,
			"sex":      sex,
			"mobile":   mobile,
		}

		if err := c.PostAndDecode("/v1/account/update", body, nil); err != nil {
			return err
		}
		fmt.Println("Account updated successfully.")
		return nil
	},
}

// ---- change-password ----

var (
	accountOldPassword string
	accountNewPassword string
)

var accountChangePasswordCmd = &cobra.Command{
	Use:   "change-password",
	Short: "Change account password",
	RunE: func(cmd *cobra.Command, args []string) error {
		if accountOldPassword == "" {
			return fmt.Errorf("--old-password is required")
		}
		if accountNewPassword == "" {
			return fmt.Errorf("--new-password is required")
		}
		c := client.NewClient()
		body := map[string]interface{}{
			"oldPassword": accountOldPassword,
			"newPassword": accountNewPassword,
		}
		if err := c.PostAndDecode("/v1/change-password", body, nil); err != nil {
			return err
		}
		fmt.Println("Password changed successfully. Please log in again with your new password.")
		return nil
	},
}

// fmtSex formats a sex value (1=Male, 2=Female).
func fmtSex(v interface{}) string {
	if f, ok := v.(float64); ok {
		switch int(f) {
		case 1:
			return "Male"
		case 2:
			return "Female"
		}
	}
	return "Unknown"
}

// fmtUserStatus formats a user status value.
func fmtUserStatus(v interface{}) string {
	if f, ok := v.(float64); ok {
		switch int(f) {
		case 1:
			return "Active"
		case 2:
			return "Pending Activation"
		case 3:
			return "Suspended"
		case 4:
			return "Banned"
		}
	}
	return fmtStr(v)
}

func init() {
	accountUpdateCmd.Flags().StringVar(&accountUpdateRealname, "realname", "", "Real name")
	accountUpdateCmd.Flags().StringVar(&accountUpdateNickname, "nickname", "", "Nickname")
	accountUpdateCmd.Flags().StringVar(&accountUpdateEngname, "engname", "", "English name")
	accountUpdateCmd.Flags().UintVar(&accountUpdateSex, "sex", 0, "Sex (1=Male, 2=Female)")
	accountUpdateCmd.Flags().StringVar(&accountUpdateMobile, "mobile", "", "Mobile number")

	accountChangePasswordCmd.Flags().StringVar(&accountOldPassword, "old-password", "", "Current password (required)")
	accountChangePasswordCmd.Flags().StringVar(&accountNewPassword, "new-password", "", "New password (min 6 characters, required)")

	accountCmd.AddCommand(accountInfoCmd)
	accountCmd.AddCommand(accountUpdateCmd)
	accountCmd.AddCommand(accountChangePasswordCmd)
}
