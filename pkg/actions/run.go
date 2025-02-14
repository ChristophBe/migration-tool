package actions

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

func Run(folder string) error {
	definition, err := loadMigrationDefinition(folder)
	if err != nil {
		return fmt.Errorf("error read migration definition: %w", err)
	}
	changed, err := verifyDefinition(folder, definition)
	if err != nil {
		return fmt.Errorf("error verifying migrations before execution: %w", err)
	}
	if changed {
		return fmt.Errorf("aborting execution: One or more migration files have changed")
	}
	for _, migration := range definition.Migrations {
		if err := executeMigration(folder, migration); err != nil {
			return fmt.Errorf("execution of migration step failed: %w", err)
		}
	}
	return nil
}

func executeMigration(folder string, migration Migration) error {
	scriptPath := filepath.Join(folder, migration.Filename)
	fmt.Println("Executing:", scriptPath)
	scriptContent, err := os.ReadFile(scriptPath)
	if err != nil {
		return fmt.Errorf("error reading script file: %w", err)
	}

	tmpFile, err := os.CreateTemp("", "temp_script_*.sh")
	if err != nil {
		return err
	}
	defer tmpFile.Close()
	defer os.Remove(tmpFile.Name())

	if _, err := io.WriteString(tmpFile, string(scriptContent)); err != nil {
		return err

	}
	err = tmpFile.Chmod(0755)
	if err != nil {
		return err
	}

	cmd := exec.Command("bash", tmpFile.Name())
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error executing script: %w", err)
	}

	fmt.Println(string(output))
	return nil
}
