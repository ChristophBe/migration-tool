package actions

import (
	"fmt"
	"path/filepath"
)

func (a *Actions) RecalculateHashes(folder string) error {

	config, err := a.loadMigrationDefinition(folder)
	if err != nil {
		return fmt.Errorf("error read migration definition: %w", err)
	}
	prevHash := ""
	for i, migration := range config.Steps {
		scriptPath := filepath.Join(folder, migration.Filename)
		hash, err := a.hashFunction.CalculateHash(scriptPath, prevHash)
		if err != nil {
			return err
		}
		config.Steps[i].Hash = hash
		prevHash = hash
	}
	return a.saveMigrationDefinition(folder, config)
}
