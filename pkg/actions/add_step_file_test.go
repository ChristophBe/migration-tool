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
	definitionVerifierMock     *MigrationDefinitionVerifierMock
}

func (s *AddStepFileTestSuite) SetupTest() {
	s.executionLoggerMock = NewExecutionLoggerMock(s.T())
	s.definitionReaderWriterMock = NewMigrationDefinitionReaderWriterMock(s.T())
	s.hashFunctionMock = NewHashFunctionMock(s.T())
	s.definitionVerifierMock = NewMigrationDefinitionVerifierMock(s.T())
	s.actions = New(s.executionLoggerMock, s.definitionReaderWriterMock, s.definitionVerifierMock, s.hashFunctionMock)
}

func (s *AddStepFileTestSuite) TestAddStepFile() {

	tt := []struct {
		name                           string
		readMigrationError             error
		hashCalculationError           error
		writeMirgrationDefinitionError error
		verificationResult             bool
		verificationError              error
	}{
		{
			name:               "no_error",
			verificationResult: true,
		},
		{
			name:               "read_migration_failed",
			readMigrationError: errors.New(rand.Text()),
			verificationResult: true,
		},
		{
			name:                 "hash_calculation_failed",
			hashCalculationError: errors.New(rand.Text()),
			verificationResult:   true,
		},
		{
			name:                           "write_migration_failed",
			writeMirgrationDefinitionError: errors.New(rand.Text()),
			verificationResult:             true,
		},
		{
			name:                           "write_migration_failed",
			writeMirgrationDefinitionError: errors.New(rand.Text()),
			verificationResult:             true,
		},
		{
			name:              "definition_verification_failed",
			verificationError: errors.New(rand.Text()),
		},
		{
			name:               "definition_invalid",
			verificationResult: false,
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

			expectedError := utils.GetNotNilError(tc.readMigrationError, tc.hashCalculationError, tc.writeMirgrationDefinitionError, tc.verificationError)

			isErrorTestCase := expectedError != nil || !tc.verificationResult

			expectedDefinitionFileName := path.Join(folder, migrationFileName)

			readCall := s.definitionReaderWriterMock.EXPECT().Read(expectedDefinitionFileName).Return(initialDefinition, tc.readMigrationError)
			defer readCall.Unset()

			verifierCall := s.definitionVerifierMock.EXPECT().Verify(folder, initialDefinition).Return(tc.verificationResult, tc.verificationError)
			defer verifierCall.Unset()
			if isErrorTestCase {
				verifierCall.Maybe()
			}

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

			err = s.actions.AddStepFile(folder, tmpFile.Name())

			if isErrorTestCase && tc.verificationResult {
				s.Require().ErrorIs(err, expectedError)
			} else if !tc.verificationResult {
				var invalidModelDefinitionError InvalidModelDefinitionError
				s.Require().ErrorAs(err, &invalidModelDefinitionError)
			} else {
				s.Require().NoError(err)
			}
		})
	}
}

func TestActions_AddStepFile(t *testing.T) {
	suite.Run(t, new(AddStepFileTestSuite))
}
