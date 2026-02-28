package cmd

import (
	"edu/cli/client"
	"fmt"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage edu-cli configuration",
}

var configSetURLCmd = &cobra.Command{
	Use:   "set-url <url>",
	Short: "Set the backend base URL",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := client.LoadConfig()
		cfg.BaseURL = args[0]
		if err := client.SaveConfig(cfg); err != nil {
			return fmt.Errorf("failed to save config: %w", err)
		}
		fmt.Printf("Base URL set to: %s\n", args[0])
		return nil
	},
}

var configSetTokenCmd = &cobra.Command{
	Use:   "set-token <token>",
	Short: "Set the authentication token (MCP token obtained from the web UI)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := client.LoadConfig()
		cfg.Token = args[0]
		if err := client.SaveConfig(cfg); err != nil {
			return fmt.Errorf("failed to save config: %w", err)
		}
		fmt.Println("Token saved successfully.")
		return nil
	},
}

var configShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the current configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := client.LoadConfig()
		fmt.Printf("Base URL : %s\n", cfg.BaseURL)
		if cfg.Token != "" {
			masked := cfg.Token
			if len(masked) > 12 {
				masked = masked[:8] + "..." + masked[len(masked)-4:]
			}
			fmt.Printf("Token    : %s\n", masked)
		} else {
			fmt.Println("Token    : (not set)")
		}
		return nil
	},
}

func init() {
	configCmd.AddCommand(configSetURLCmd)
	configCmd.AddCommand(configSetTokenCmd)
	configCmd.AddCommand(configShowCmd)
}
