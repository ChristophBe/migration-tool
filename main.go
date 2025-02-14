package main

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

type Migration struct {
	Filename    string `yaml:"filename"`
	Description string `yaml:"description"`
	Hash        string `yaml:"hash"`
}

type Config struct {
	Migrations []Migration `yaml:"migrations"`
}

func loadConfig(folder string) (*Config, error) {
	yamlPath := filepath.Join(folder, "migrations.yaml")
	data, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		return nil, fmt.Errorf("error reading YAML file: %w", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("error parsing YAML file: %w", err)
	}

	return &config, nil
}

func saveConfig(folder string, config *Config) error {
	yamlPath := filepath.Join(folder, "migrations.yaml")
	data, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("error marshalling YAML: %w", err)
	}
	return ioutil.WriteFile(yamlPath, data, 0644)
}

func calculateHash(filename string, prevHash string) (string, error) {
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("error reading file %s: %w", filename, err)
	}
	hasher := sha256.New()
	hasher.Write(fileContent)
	hasher.Write([]byte(prevHash))
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash, nil
}

func recalculateHashes(folder string, config *Config) error {
	prevHash := ""
	for i, migration := range config.Migrations {
		scriptPath := filepath.Join(folder, migration.Filename)
		hash, err := calculateHash(scriptPath, prevHash)
		if err != nil {
			return err
		}
		config.Migrations[i].Hash = hash
		prevHash = hash
	}
	return saveConfig(folder, config)
}

func verifyMigrations(folder string, config *Config) (bool, error) {
	prevHash := ""
	changed := false
	for _, migration := range config.Migrations {
		scriptPath := filepath.Join(folder, migration.Filename)
		hash, err := calculateHash(scriptPath, prevHash)
		if err != nil {
			return false, err
		}
		if hash != migration.Hash {
			fmt.Printf("Warning: %s has changed!\n", migration.Filename)
			changed = true
		} else {
			fmt.Printf("%s is unchanged.\n", migration.Filename)
		}
		prevHash = hash
	}
	return changed, nil
}

func executeMigration(folder string, migration Migration) {
	scriptPath := filepath.Join(folder, migration.Filename)
	fmt.Println("Executing:", scriptPath)
	scriptContent, err := ioutil.ReadFile(scriptPath)
	if err != nil {
		fmt.Println("Error reading script file:", err)
		return
	}

	tmpFile, err := os.CreateTemp("", "temp_script_*.sh")
	if err != nil {
		fmt.Println("Error creating temporary file:", err)
		return
	}
	defer os.Remove(tmpFile.Name())

	if _, err := io.WriteString(tmpFile, string(scriptContent)); err != nil {
		fmt.Println("Error writing script file:", err)
		return
	}
	tmpFile.Chmod(0755)
	tmpFile.Close()

	cmd := exec.Command("bash", tmpFile.Name())
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing script:", err)
	}

	fmt.Println(string(output))
}

func main() {
	flag.Usage = func() {
		fmt.Println("Usage: migration-tool <command> [options]")
		fmt.Println("Commands:")
		fmt.Println("  recalculate-hashes   Recalculate migration hashes")
		fmt.Println("  verify               Verify if migration files have changed")
		fmt.Println("  run                 Run the migrations")
		flag.PrintDefaults()
	}

	folder := flag.String("folder", "migrations", "Folder where migrations.yaml and scripts are located")
	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	command := flag.Arg(0)

	config, err := loadConfig(*folder)
	if err != nil {
		fmt.Println(err)
		return
	}

	switch command {
	case "recalculate-hashes":
		if err := recalculateHashes(*folder, config); err != nil {
			fmt.Println("Error recalculating hashes:", err)
		}
	case "verify":
		if _, err := verifyMigrations(*folder, config); err != nil {
			fmt.Println("Error verifying migrations:", err)
		}
	case "run":
		changed, err := verifyMigrations(*folder, config)
		if err != nil {
			fmt.Println("Error verifying migrations before execution:", err)
			return
		}
		if changed {
			fmt.Println("Aborting execution: One or more migration files have changed.")
			return
		}
		for _, migration := range config.Migrations {
			executeMigration(*folder, migration)
		}
	default:
		fmt.Println("Unknown command:", command)
		flag.Usage()
		os.Exit(1)
	}
}
