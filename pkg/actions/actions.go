package actions

type Actions struct {
	executionLogger        ExecutionLogger
	definitionReaderWriter MigrationDefinitionReaderWriter
}

func New(ExecutionLogger ExecutionLogger, DefinitionReaderWriter MigrationDefinitionReaderWriter) *Actions {
	return &Actions{
		executionLogger:        ExecutionLogger,
		definitionReaderWriter: DefinitionReaderWriter,
	}
}
