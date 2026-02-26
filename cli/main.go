// edu-cli is a command-line client for the LTEdu platform.
//
// Usage:
//
//	edu-cli <command> [subcommand] [flags]
//
// Configuration (environment variables):
//
//	EDU_BASE_URL  Backend base URL (e.g. https://api.example.com)
//	EDU_TOKEN     JWT token obtained from the web UI
package main

import "edu/cli/cmd"

func main() {
	cmd.Execute()
}
