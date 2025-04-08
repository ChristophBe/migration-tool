package actions

import (
	"crypto/rand"
	"errors"
	"github.com/ChristophBe/migration-tool/internal/utils"
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
	s.actions = New(s.executionLoggerMock, s.definitionReaderWriterMock, nil, s.hashFunctionMock)
}

func (s *AddStepFileTestSuite) TestAddStepFile() {

	tt := []struct {
		name                           string
		readMigrationError             error
		hashCalculationError           error
		writeMirgrationDefinitionError error
	}{
		{
			name: "no_error",
		},
		{
			name:               "read_migration_failed",
			readMigrationError: errors.New(rand.Text()),
		},
		{
			name:                 "hash_calculation_failed",
			hashCalculationError: errors.New(rand.Text()),
		},
		{
			name:                           "write_migration_failed",
			writeMirgrationDefinitionError: errors.New(rand.Text()),
		},
	}

	for _, tc := range tt {
		s.Run(tc.name, func() {

			tmpFile, err := os.CreateTemp("", "test_file_*.sh")
			s.Require().NoError(err)

			_, err = io.WriteString(tmpFile, rand.Text())
			s.Require().NoError(err)

			folder, expectedFileName := path.Split(tmpFile.Name())

			previousHash := rand.Text()

			initialDefinition := MigrationDefinition{
				Steps: []MigrationStep{
					{
						Filename:    rand.Text() + ".sh",
						Description: rand.Text(),
						Hash:        previousHash,
					},
				},
			}

			expectedError := utils.GetNotNilError(tc.readMigrationError, tc.hashCalculationError, tc.writeMirgrationDefinitionError)
			isErrorTestCase := expectedError != nil

			expectedDefinitionFileName := path.Join(folder, migrationFileName)

			readCall := s.definitionReaderWriterMock.EXPECT().Read(expectedDefinitionFileName).Return(initialDefinition, tc.readMigrationError)
			defer readCall.Unset()

			expectedHash := rand.Text()
			calculateHashCall := s.hashFunctionMock.EXPECT().CalculateHash(tmpFile.Name(), initialDefinition.Steps[0].Hash).Return(expectedHash, tc.hashCalculationError)

			if isErrorTestCase {
				calculateHashCall.Maybe()
			}

			expectedResultDefinition := MigrationDefinition{}
			expectedResultDefinition.Steps = append(initialDefinition.Steps, MigrationStep{
				Filename: expectedFileName,
				Hash:     expectedHash,
			})

			writeCall := s.definitionReaderWriterMock.EXPECT().Write(expectedDefinitionFileName, expectedResultDefinition).Return(tc.writeMirgrationDefinitionError)
			if isErrorTestCase {
				writeCall.Maybe()
			}

			err = s.actions.AddStepFile(folder, expectedFileName)

			if isErrorTestCase {
				s.Require().ErrorIs(err, expectedError)
			} else {
				s.NoError(err)
			}
		})
	}
}

func TestActions_AddStepFile(t *testing.T) {
	suite.Run(t, new(AddStepFileTestSuite))
}
