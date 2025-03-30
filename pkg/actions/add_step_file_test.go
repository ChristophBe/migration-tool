package actions

import (
	"crypto/rand"
	"github.com/stretchr/testify/suite"
	"io"
	"os"
	"path"
	"testing"
)

type AddStepFileTestSuite struct {
	suite.Suite
	actions                    *Actions
	definitionReaderWriterMock *MigrationDefinitionReaderWriterMock
	executionLoggerMock        *ExecutionLoggerMock
	hashFunctionMock           *HashFunctionMock
}

func (s *AddStepFileTestSuite) SetupTest() {
	s.executionLoggerMock = NewExecutionLoggerMock(s.T())
	s.definitionReaderWriterMock = NewMigrationDefinitionReaderWriterMock(s.T())
	s.hashFunctionMock = NewHashFunctionMock(s.T())
	s.actions = New(s.executionLoggerMock, s.definitionReaderWriterMock, s.hashFunctionMock)
}

func (s *AddStepFileTestSuite) TestAddStepFile() {
	tmpFile, err := os.CreateTemp("", "test_file_*.sh")
	s.Require().NoError(err)

	_, err = io.WriteString(tmpFile, rand.Text())
	s.Require().NoError(err)

	folder, expectedFileName := path.Split(tmpFile.Name())

	initialDefinition := MigrationDefinition{
		Steps: []MigrationStep{
			{
				Filename:    rand.Text() + ".sh",
				Description: rand.Text(),
				Hash:        rand.Text(),
			},
		},
	}

	expectedDefinitionFileName := path.Join(folder, migrationFileName)

	s.definitionReaderWriterMock.EXPECT().Read(expectedDefinitionFileName).Return(initialDefinition, nil)

	expectedResultDefinition := MigrationDefinition{}

	expectedHash := rand.Text()
	s.hashFunctionMock.EXPECT().CalculateHash(tmpFile.Name(), initialDefinition.Steps[0].Hash).Return(expectedHash, nil)
	s.NoError(err)
	expectedResultDefinition.Steps = append(initialDefinition.Steps, MigrationStep{
		Filename: expectedFileName,
		Hash:     expectedHash,
	})

	s.definitionReaderWriterMock.EXPECT().Write(expectedDefinitionFileName, expectedResultDefinition).Return(nil)

	err = s.actions.AddStepFile(folder, expectedFileName)
	s.NoError(err)
}

func TestActions_AddStepFile(t *testing.T) {
	suite.Run(t, new(AddStepFileTestSuite))
}
