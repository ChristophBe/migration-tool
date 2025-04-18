package cmd

import (
	"github.com/ChristophBe/migration-tool/internal/utils"
	"github.com/ChristophBe/migration-tool/pkg/actions"
	"github.com/ChristophBe/migration-tool/pkg/execution_loggers"
)

func initActions(outputFolder string) Actions {

	outfileReaderWriter := utils.NewYamlReaderWriter[execution_loggers.ExecutionLogs]()
	fileExecutionLogger := execution_loggers.NewFileExecutionLogger(outputFolder, outfileReaderWriter)
	definitionWriterReader := utils.NewYamlReaderWriter[actions.MigrationDefinition]()

	return actions.New(fileExecutionLogger, definitionWriterReader)
}
