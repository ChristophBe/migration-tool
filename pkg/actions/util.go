package actions

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/fs"
	"os"
	"path/filepath"
)

func LoadYaml[T any](filepath string) (*T, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("error reading YAML file: %w", err)
	}

	var content T
	err = yaml.Unmarshal(data, &content)
	if err != nil {
		return nil, fmt.Errorf("error parsing YAML file: %w", err)
	}

	return &content, nil
}

func SaveYaml[T any](filepath string, content *T) error {
	data, err := yaml.Marshal(content)
	if err != nil {
		return fmt.Errorf("error marshalling YAML: %w", err)
	}
	return os.WriteFile(filepath, data, 0644)
}

func saveMigrationDefinition(folder string, migrationDefinition *MigrationDefinition) error {
	file := filepath.Join(folder, migrationFileName)
	return SaveYaml(file, migrationDefinition)
}
func loadMigrationDefinition(folder string) (*MigrationDefinition, error) {
	file := filepath.Join(folder, migrationFileName)
	return LoadYaml[MigrationDefinition](file)
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

func updateResults(folder string, results *Results, steps []StepResult) error {
	results.Steps = append(results.Steps, steps...)
	return SaveYaml(filepath.Join(folder, outputFileName), results)
}

func loadResults(folder string) (*Results, error) {
	res, err := LoadYaml[Results](filepath.Join(folder, outputFileName))

	if errors.Is(err, fs.ErrNotExist) {
		return &Results{}, nil
	}

	return res, err
}
