package actions

type Actions struct {
	executionLogger        ExecutionLogger
	definitionReaderWriter MigrationDefinitionReaderWriter
	hashFunction           HashFunction
}

func New(executionLogger ExecutionLogger, definitionReaderWriter MigrationDefinitionReaderWriter, hashFunction HashFunction) *Actions {
	return &Actions{
		executionLogger:        executionLogger,
		definitionReaderWriter: definitionReaderWriter,
		hashFunction:           hashFunction,
	}
}
