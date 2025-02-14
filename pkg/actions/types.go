package actions

type Migration struct {
	Filename    string `yaml:"filename"`
	Description string `yaml:"description"`
	Hash        string `yaml:"hash"`
}

type MigrationDefinition struct {
	Migrations []Migration `yaml:"migrations"`
}

type Output struct {
	ExecutedSteps []ExecutionStep `yaml:"executedSteps"`
}
type ExecutionStep struct {
	Timestamp string `yaml:"timestamp"`
	Hash      string `yaml:"hash"`
}
