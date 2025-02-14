package actions

import (
	"fmt"
	"path/filepath"
)

func RecalculateHashes(folder string) error {

	config, err := loadMigrationDefinition(folder)
	if err != nil {
		return fmt.Errorf("error read migration definition: %w", err)
	}
	prevHash := ""
	for i, migration := range config.Migrations {
		scriptPath := filepath.Join(folder, migration.Filename)
		hash, err := CalculateHash(scriptPath, prevHash)
		if err != nil {
			return err
		}
		config.Migrations[i].Hash = hash
		prevHash = hash
	}
	return saveMigrationDefinition(folder, config)
}
