package actions

type MigrationDefinitionReaderWriter interface {
	Read(fileName string) (MigrationDefinition, error)
	Write(fileName string, vales MigrationDefinition) error
}
