/*
Copyright © 2025 Christoph Becker <post@christopb.de>
*/
package cmd

import (
	"github.com/spf13/cobra/doc"
	"os"

	"github.com/spf13/cobra"
)

var baseFolder string
var executionLogFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "migration-tool",
	Short: "migration-tool is a CLI that orchestrates the execution of scripts.",
	Long: `migration-tool is a CLI that orchestrates the execution of bash scripts organized as steps.

It makes sure that the scripts are executed in the correct order and that the scripts are only executed if the script have not run before.
To ensure consistency, the scripts are checked for unexecuted changes.
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&baseFolder, "folder", "", "Folder where the scripts and configurations file are located.")
	rootCmd.PersistentFlags().StringVarP(&executionLogFile, "execution-log-file", "o", "execution-log.yaml", "File where the execution log is written to.")
}

func GenerateDoc(folder string) error {
	return doc.GenMarkdownTree(rootCmd, folder)
}
