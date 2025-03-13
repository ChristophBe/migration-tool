package utils

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type YamlReaderWriter[T any] struct{}

func NewYamlReaderWriter[T any]() YamlReaderWriter[T] {
	return YamlReaderWriter[T]{}

}

func (y YamlReaderWriter[T]) Write(filepath string, value T) error {
	data, err := yaml.Marshal(value)
	if err != nil {
		return fmt.Errorf("error marshalling YAML: %w", err)
	}
	return os.WriteFile(filepath, data, 0644)
}

func (y YamlReaderWriter[T]) Read(filepath string) (T, error) {
	var result T
	data, err := os.ReadFile(filepath)
	if err != nil {
		return result, fmt.Errorf("error reading YAML file: %w", err)
	}

	err = yaml.Unmarshal(data, &result)
	if err != nil {
		return result, fmt.Errorf("error parsing YAML file: %w", err)
	}

	return result, nil
}
