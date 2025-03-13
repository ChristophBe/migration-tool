package execution_loggers

import (
	"errors"
	"github.com/ChristophBe/migration-tool/pkg/actions"
	"io/fs"
)

type ExecutionLogs = actions.ExecutionLogs
type OutputFileReaderWriter interface {
	Read(filename string) (ExecutionLogs, error)
	Write(filename string, logs ExecutionLogs) error
}

type FileExecutionLogger struct {
	outputFilename         string
	outputFileReaderWriter OutputFileReaderWriter
}

func NewFileExecutionLogger(outputFilename string, outputFileReaderWriter OutputFileReaderWriter) *FileExecutionLogger {
	return &FileExecutionLogger{
		outputFilename:         outputFilename,
		outputFileReaderWriter: outputFileReaderWriter,
	}
}

func (f *FileExecutionLogger) LoadExecutionLog() (actions.ExecutionLogs, error) {
	res, err := f.outputFileReaderWriter.Read(f.outputFilename)

	if errors.Is(err, fs.ErrNotExist) {
		return actions.ExecutionLogs{}, nil
	}

	return res, err
}

func (f *FileExecutionLogger) LogExecution(steps []actions.StepResult) error {
	results, err := f.LoadExecutionLog()
	if err != nil {
		return err
	}
	results.Steps = append(results.Steps, steps...)
	return f.outputFileReaderWriter.Write(f.outputFilename, results)
}
