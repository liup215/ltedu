package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var genDocsDir string

var genDocsCmd = &cobra.Command{
	Use:   "generate-docs",
	Short: "Generate Markdown documentation for all commands",
	Long: `Generate Markdown documentation files for edu-cli and all its subcommands.

Each command gets its own .md file in the specified output directory.
The generated files can be used as standalone documentation or integrated
into a project documentation site.

Example:
  edu-cli generate-docs --output ./docs/cli`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := os.MkdirAll(genDocsDir, 0755); err != nil {
			return fmt.Errorf("failed to create output directory %q: %w", genDocsDir, err)
		}
		if err := doc.GenMarkdownTree(rootCmd, genDocsDir); err != nil {
			return fmt.Errorf("failed to generate docs: %w", err)
		}
		fmt.Printf("Documentation generated in: %s\n", genDocsDir)
		return nil
	},
}

func init() {
	genDocsCmd.Flags().StringVarP(&genDocsDir, "output", "o", "./docs/cli", "Output directory for generated Markdown files")
	rootCmd.AddCommand(genDocsCmd)
}
