/*
Copyright © 2025 Christoph Becker <post@christopb.de>
*/
package cmd

import (
	"github.com/ChristophBe/migration-tool/internal/utils"
	"github.com/ChristophBe/migration-tool/pkg/actions"
	"github.com/ChristophBe/migration-tool/pkg/execution_loggers"
	"github.com/spf13/cobra/doc"
	"os"

	"github.com/spf13/cobra"
)

var baseFolder string
var executionLogFile string

var acts Actions

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

	outfileReaderWriter := utils.NewYamlReaderWriter[execution_loggers.ExecutionLogs]()
	fileExecutionLogger := execution_loggers.NewFileExecutionLogger(executionLogFile, outfileReaderWriter)
	definitionWriterReader := utils.NewYamlReaderWriter[actions.MigrationDefinition]()

	acts = actions.New(fileExecutionLogger, definitionWriterReader)
}

func GenerateDoc(folder string) error {
	rootCmd.DisableAutoGenTag = true
	return doc.GenMarkdownTree(rootCmd, folder)
}
