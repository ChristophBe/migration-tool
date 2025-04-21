package actions

import (
	"fmt"
)

func (a *Actions) Verify(folder string) (bool, error) {
	migrationDefinition, err := a.loadMigrationDefinition(folder)
	if err != nil {
		return false, fmt.Errorf("error read migration definition: %w", err)
	}
	return a.definitionVerifier.Verify(folder, migrationDefinition)
}
