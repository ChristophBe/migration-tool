package main

import (
	"flag"
	"fmt"
	"github.com/ChristophBe/migration-tool/internal/utils"
	"github.com/ChristophBe/migration-tool/pkg/actions"
	"github.com/ChristophBe/migration-tool/pkg/execution_loggers"
	"log"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Usage: migration-tool [options] <command>")
		fmt.Println("Commands:")
		fmt.Println("  recalculate-hashes   Recalculate migration hashes")
		fmt.Println("  verify               Verify if migration files have changed")
		fmt.Println("  run                  Run the migrations")
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	folder := flag.String("folder", "migrations", "Folder where migrations.yaml and scripts are located")
	outputFolder := flag.String("execution-filename", "execution-log.yaml", "File where the the executes steps are logged.")
	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	command := flag.Arg(0)

	ouputfileReaderWriter := utils.NewYamlReaderWriter[execution_loggers.ExecutionLogs]()
	fileExecutionLogger := execution_loggers.NewFileExecutionLogger(*outputFolder, ouputfileReaderWriter)
	definitionWriterReader := utils.NewYamlReaderWriter[actions.MigrationDefinition]()

	err := RunCommands(actions.New(fileExecutionLogger, definitionWriterReader), command, *folder)
	if err != nil {
		log.Fatal(err)
	}
}
func RunCommands(act Actions, command, folder string) (err error) {

	switch command {
	case "recalculate-hashes":
		if err := act.RecalculateHashes(folder); err != nil {
			return fmt.Errorf("error recalculating hashes: %w", err)
		}
	case "verify":
		var changesDetected bool
		if changesDetected, err = act.Verify(folder); err != nil {
			return fmt.Errorf("error verifying migrations: %w", err)
		}

		if changesDetected {
			return fmt.Errorf("changes were detected during verification")
		}
	case "run":
		if err = act.Run(folder); err != nil {
			return fmt.Errorf("error running migrations: %w", err)
		}
	case "help":
		flag.Usage()

	default:
		fmt.Println("Unknown command:", command)
		flag.Usage()
	}
	return nil
}
