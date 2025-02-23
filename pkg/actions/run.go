package actions

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

const (
	outputFileName    = "output.yaml"
	migrationFileName = "migrations.yaml"
)

func (a *Actions) Run(folder, outputFolder string) error {
	definition, err := loadMigrationDefinition(folder)
	if err != nil {
		return fmt.Errorf("error read step definition: %w", err)
	}
	changed, err := verifyDefinition(folder, definition)
	if err != nil {
		return fmt.Errorf("error verifying migrations before execution: %w", err)
	}
	if changed {
		return fmt.Errorf("aborting execution: One or more step files have changed")
	}

	results, err := loadResults(outputFolder)
	if err != nil {
		return err
	}

	index := len(results.Steps)

	if index > 0 {
		lastStepIndex := index - 1

		lastExecutedStep := results.Steps[lastStepIndex]
		lastExecutedStepDefinition := definition.Steps[lastStepIndex]

		if lastExecutedStep.Hash != lastExecutedStepDefinition.Hash {
			return fmt.Errorf("hash of last executed step and its definition does not match")

		}
	}

	var executedSteps []StepResult
	for i := index; i < len(definition.Steps); i++ {
		step := definition.Steps[i]
		var res StepResult
		if res, err = executeMigrationStep(folder, step); err != nil {
			return fmt.Errorf("execution of step step failed: %w", err)
		}
		executedSteps = append(executedSteps, res)
	}

	err = updateResults(outputFolder, results, executedSteps)
	return nil
}

func executeMigrationStep(folder string, step MigrationStep) (res StepResult, err error) {
	scriptPath := filepath.Join(folder, step.Filename)
	fmt.Println("Executing:", scriptPath)
	scriptContent, err := os.ReadFile(scriptPath)
	if err != nil {
		err = fmt.Errorf("error reading script file: %w", err)
		return
	}

	tmpFile, err := os.CreateTemp("", "temp_script_*.sh")
	if err != nil {
		return
	}
	defer tmpFile.Close()
	defer os.Remove(tmpFile.Name())

	if _, err = io.WriteString(tmpFile, string(scriptContent)); err != nil {
		return

	}
	err = tmpFile.Chmod(0755)
	if err != nil {
		return
	}

	cmd := exec.Command("bash", tmpFile.Name())
	output, err := cmd.CombinedOutput()
	if err != nil {
		err = fmt.Errorf("error executing script: %w", err)
		return
	}

	fmt.Println(string(output))
	res.Hash = step.Hash
	res.Timestamp = time.Now()
	return
}
