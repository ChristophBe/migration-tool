package actions

import (
	"path/filepath"
)

func (a *Actions) saveMigrationDefinition(folder string, migrationDefinition MigrationDefinition) error {
	file := filepath.Join(folder, migrationFileName)
	return a.definitionReaderWriter.Write(file, migrationDefinition)
}
func (a *Actions) loadMigrationDefinition(folder string) (MigrationDefinition, error) {
	file := filepath.Join(folder, migrationFileName)

	return a.definitionReaderWriter.Read(file)
}
