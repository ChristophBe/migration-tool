package actions

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
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

func CalculateHash(filename string, prevHash string) (string, error) {
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("error reading file %s: %w", filename, err)
	}
	hasher := sha256.New()
	hasher.Write(fileContent)
	hasher.Write([]byte(prevHash))
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash, nil
}
