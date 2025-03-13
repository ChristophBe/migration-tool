package actions

import (
	"crypto/rand"
	"fmt"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"os"
	"path"
	"strconv"
	"testing"
)

type ActionRunTestSuite struct {
	suite.Suite
	inputFolder                string
	outFolder                  string
	runOutputFile              string
	actions                    *Actions
	executionLoggerMock        *ExecutionLoggerMock
	definitionReaderWriterMock *MigrationDefinitionReaderWriterMock
}

func (s *ActionRunTestSuite) SetupTest() {

	s.executionLoggerMock = NewExecutionLoggerMock(s.T())
	s.definitionReaderWriterMock = NewMigrationDefinitionReaderWriterMock(s.T())
	s.actions = New(s.executionLoggerMock, s.definitionReaderWriterMock)

	var err error
	s.inputFolder, err = os.MkdirTemp("", "test-input")
	s.NoError(err)

	s.outFolder, err = os.MkdirTemp("", "test-output")
	s.NoError(err)

	s.runOutputFile = path.Join(s.outFolder, "run-output.txt")

}
func (s *ActionRunTestSuite) addMirgrationConfig() (MigrationDefinition, string) {
	migrationDefinition := MigrationDefinition{}
	var expectedOutput string

	lastHash := ""
	for i := 0; i < 2; i++ {
		filename := rand.Text() + ".sh"

		runOutput := fmt.Sprintf("step-%d-%s", i, rand.Text())
		expectedOutput = fmt.Sprintf("%s%s\n", expectedOutput, runOutput)
		content := `
#!/bin/bash
echo "` + runOutput + `" >> ` + s.runOutputFile + `
`
		scriptFilepath := path.Join(s.inputFolder, filename)
		err := os.WriteFile(scriptFilepath, []byte(content), 0644)
		s.NoError(err)
		lastHash, err = CalculateHash(scriptFilepath, lastHash)
		migrationDefinition.Steps = append(migrationDefinition.Steps, MigrationStep{
			Filename:    filename,
			Description: "step " + strconv.Itoa(i),
			Hash:        lastHash,
		})

	}

	return migrationDefinition, expectedOutput
}

func (s *ActionRunTestSuite) TestRun() {
	migrationDefinition, expectedOutput := s.addMirgrationConfig()

	s.definitionReaderWriterMock.EXPECT().Read(path.Join(s.inputFolder, migrationFileName)).Return(migrationDefinition, nil)
	s.executionLoggerMock.EXPECT().LoadExecutionLog().Return(ExecutionLogs{}, nil)

	s.executionLoggerMock.EXPECT().LogExecution(mock.Anything).Run(func(results []StepResult) {
		for i, result := range results {
			s.NotZero(result.Timestamp)
			s.Equal(migrationDefinition.Steps[i].Hash, result.Hash)
		}
	}).Return(nil)
	err := s.actions.Run(s.inputFolder)
	s.NoError(err)

	s.FileExists(s.runOutputFile)
	runOutputs, err := os.ReadFile(s.runOutputFile)
	s.NoError(err)
	s.Equal(expectedOutput, string(runOutputs))
}

func (s *ActionRunTestSuite) TearDownTest() {
	s.NoError(os.RemoveAll(s.inputFolder))
	s.NoError(os.RemoveAll(s.outFolder))
}

func TestActions_Run(t *testing.T) {
	suite.Run(t, new(ActionRunTestSuite))
}
