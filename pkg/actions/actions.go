package actions

type Actions struct {
	executionLogger        ExecutionLogger
	definitionReaderWriter MigrationDefinitionReaderWriter
	definitionVerifier     MigrationDefinitionVerifier
	hashFunction           HashFunction
}

func New(executionLogger ExecutionLogger, definitionReaderWriter MigrationDefinitionReaderWriter, definitionVerifier MigrationDefinitionVerifier, hashFunction HashFunction) *Actions {
	return &Actions{
		executionLogger:        executionLogger,
		definitionReaderWriter: definitionReaderWriter,
		hashFunction:           hashFunction,
		definitionVerifier:     definitionVerifier,
	}
}
