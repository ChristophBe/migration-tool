package execution_loggers

import (
	"errors"
	"github.com/ChristophBe/migration-tool/internal/utils"
	"github.com/ChristophBe/migration-tool/pkg/actions"
	"io/fs"
)

type FileExecutionLogger struct {
	outputFilename string
}

func NewFileExecutionLogger(outputFilename string) *FileExecutionLogger {
	return &FileExecutionLogger{
		outputFilename: outputFilename,
	}
}

func (f *FileExecutionLogger) LoadExecutionLog() (actions.ExecutionLogs, error) {
	res, err := utils.LoadYaml[actions.ExecutionLogs](f.outputFilename)

	if errors.Is(err, fs.ErrNotExist) {
		return actions.ExecutionLogs{}, nil
	}

	return *res, err
}

func (f *FileExecutionLogger) LogExecution(steps []actions.StepResult) error {
	results, err := f.LoadExecutionLog()
	if err != nil {
		return err
	}
	results.Steps = append(results.Steps, steps...)
	return utils.SaveYaml[actions.ExecutionLogs](f.outputFilename, results)
}
