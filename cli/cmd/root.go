package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "edu-cli",
	Short: "edu-cli is a command-line client for managing the LTEdu platform",
	Long: `edu-cli lets you manage syllabuses, papers, past papers, users,
classes, and students through the LTEdu backend API.

Configuration:
  Set EDU_BASE_URL to point to the backend (e.g. https://api.example.com)
  Set EDU_TOKEN   to authenticate with a JWT token obtained from the web UI

  Alternatively, run 'edu-cli config' commands to persist settings locally.`,
}

// Execute runs the root command and exits on error.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(accountCmd)
	rootCmd.AddCommand(organisationCmd)
	rootCmd.AddCommand(qualificationCmd)
	rootCmd.AddCommand(syllabusCmd)
	rootCmd.AddCommand(chapterCmd)
	rootCmd.AddCommand(questionCmd)
	rootCmd.AddCommand(paperCmd)
	rootCmd.AddCommand(userCmd)
	rootCmd.AddCommand(classCmd)
	rootCmd.AddCommand(learningPlanCmd)
	rootCmd.AddCommand(examNodeCmd)
	rootCmd.AddCommand(phasePlanCmd)
	rootCmd.AddCommand(teacherCmd)
}

// prettyPrint prints any value as indented JSON.
func prettyPrint(v interface{}) {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(v); err != nil {
		fmt.Fprintf(os.Stderr, "error formatting output: %v\n", err)
	}
}

// fmtStr safely converts an interface{} to string.
func fmtStr(v interface{}) string {
	if v == nil {
		return ""
	}
	return fmt.Sprintf("%v", v)
}

// fmtFloat safely converts a JSON number (float64) to a string integer representation.
func fmtFloat(v interface{}) string {
	if v == nil {
		return ""
	}
	if f, ok := v.(float64); ok {
		return strconv.Itoa(int(f))
	}
	return fmt.Sprintf("%v", v)
}

// fmtBool formats a boolean interface{} value.
func fmtBool(v interface{}) string {
	if v == nil {
		return "false"
	}
	if b, ok := v.(bool); ok {
		if b {
			return "true"
		}
		return "false"
	}
	return fmt.Sprintf("%v", v)
}

// printTable prints a simple table with headers and rows.
func printTable(headers []string, rows [][]string) {
	// Calculate column widths
	widths := make([]int, len(headers))
	for i, h := range headers {
		widths[i] = len(h)
	}
	for _, row := range rows {
		for i, cell := range row {
			if i < len(widths) && len(cell) > widths[i] {
				widths[i] = len(cell)
			}
		}
	}

	// Print header
	for i, h := range headers {
		fmt.Printf("%-*s  ", widths[i], h)
	}
	fmt.Println()

	// Print separator
	for _, w := range widths {
		for j := 0; j < w+2; j++ {
			fmt.Print("-")
		}
	}
	fmt.Println()

	// Print rows
	for _, row := range rows {
		for i, cell := range row {
			if i < len(widths) {
				fmt.Printf("%-*s  ", widths[i], cell)
			}
		}
		fmt.Println()
	}
}
