package utils

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
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

func SaveYaml[T any](filepath string, content T) error {
	data, err := yaml.Marshal(content)
	if err != nil {
		return fmt.Errorf("error marshalling YAML: %w", err)
	}
	return os.WriteFile(filepath, data, 0644)
}
