package actions

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ChristophBe/migration-tool/internal/utils"
	"io/fs"
	"os"
	"path/filepath"
)

func saveMigrationDefinition(folder string, migrationDefinition *MigrationDefinition) error {
	file := filepath.Join(folder, migrationFileName)
	return utils.SaveYaml(file, migrationDefinition)
}
func loadMigrationDefinition(folder string) (*MigrationDefinition, error) {
	file := filepath.Join(folder, migrationFileName)
	return utils.LoadYaml[MigrationDefinition](file)
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

func updateResults(folder string, results *ExecutionLogs, steps []StepResult) error {
	results.Steps = append(results.Steps, steps...)
	return utils.SaveYaml(filepath.Join(folder, outputFileName), results)
}

func loadResults(folder string) (*ExecutionLogs, error) {
	res, err := utils.LoadYaml[ExecutionLogs](filepath.Join(folder, outputFileName))

	if errors.Is(err, fs.ErrNotExist) {
		return &ExecutionLogs{}, nil
	}

	return res, err
}
