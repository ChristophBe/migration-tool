package actions

import (
	"fmt"
	"path/filepath"
)

func (a *Actions) Verify(folder string) (bool, error) {

	migrationDefinition, err := loadMigrationDefinition(folder)
	if err != nil {
		return false, fmt.Errorf("error read migration definition: %w", err)
	}
	return verifyDefinition(folder, migrationDefinition)
}

func verifyDefinition(folder string, migrationDefinition *MigrationDefinition) (bool, error) {
	prevHash := ""
	changed := false
	for _, migration := range migrationDefinition.Steps {
		scriptPath := filepath.Join(folder, migration.Filename)
		hash, err := CalculateHash(scriptPath, prevHash)
		if err != nil {
			return false, err
		}
		if hash != migration.Hash {
			fmt.Printf("Warning: %s has changed!\n", migration.Filename)
			changed = true
		} else {
			fmt.Printf("%s is unchanged.\n", migration.Filename)
		}
		prevHash = hash
	}
	return changed, nil
}
