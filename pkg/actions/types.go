package actions

import "time"

type MigrationStep struct {
	Filename    string `yaml:"filename"`
	Description string `yaml:"description"`
	Hash        string `yaml:"hash"`
}

type MigrationDefinition struct {
	Steps []MigrationStep `yaml:"steps"`
}

type Results struct {
	Steps []StepResult `yaml:"steps"`
}
type StepResult struct {
	Timestamp time.Time `yaml:"timestamp"`
	Hash      string    `yaml:"hash"`
}
