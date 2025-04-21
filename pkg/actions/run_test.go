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
	hashFunctionMock           *HashFunctionMock
	definitionVerifierMock     *MigrationDefinitionVerifierMock
}

func (s *ActionRunTestSuite) SetupTest() {

	s.executionLoggerMock = NewExecutionLoggerMock(s.T())
	s.definitionReaderWriterMock = NewMigrationDefinitionReaderWriterMock(s.T())
	s.hashFunctionMock = NewHashFunctionMock(s.T())
	s.definitionVerifierMock = NewMigrationDefinitionVerifierMock(s.T())
	s.actions = New(s.executionLoggerMock, s.definitionReaderWriterMock, s.definitionVerifierMock, s.hashFunctionMock)

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

	for i := 0; i < 2; i++ {
		filename := fmt.Sprintf("step-%d-%s.sh", i, rand.Text())

		runOutput := fmt.Sprintf("step-%d-%s", i, rand.Text())
		expectedOutput = fmt.Sprintf("%s%s\n", expectedOutput, runOutput)
		content := `
#!/bin/bash
echo "` + runOutput + `" >> ` + s.runOutputFile + `
`
		scriptFilepath := path.Join(s.inputFolder, filename)
		err := os.WriteFile(scriptFilepath, []byte(content), 0644)
		s.Require().NoError(err)

		migrationDefinition.Steps = append(migrationDefinition.Steps, MigrationStep{
			Filename:    filename,
			Description: "step " + strconv.Itoa(i),
			Hash:        rand.Text(),
		})
	}

	return migrationDefinition, expectedOutput
}

func (s *ActionRunTestSuite) TestRun() {
	migrationDefinition, expectedOutput := s.addMirgrationConfig()

	s.definitionReaderWriterMock.EXPECT().Read(path.Join(s.inputFolder, migrationFileName)).Return(migrationDefinition, nil)

	s.definitionVerifierMock.EXPECT().Verify(s.inputFolder, migrationDefinition).Return(true, nil)
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
