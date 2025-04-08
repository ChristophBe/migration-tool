package actions

type MigrationDefinitionVerifier interface {
	Verify(folder string, definition MigrationDefinition) (bool, error)
}
