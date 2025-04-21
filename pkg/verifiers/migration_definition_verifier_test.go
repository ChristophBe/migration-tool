package verifiers

import (
	"crypto/rand"
	"errors"
	"github.com/stretchr/testify/suite"
	"path"
	"testing"
)

type ModelDefinitionVerifierTestSuite struct {
	suite.Suite
	modelDefinitionVerifier *ModelDefinitionVerifier
	hashFunctionMock        *HashFunctionMock
}

func (s *ModelDefinitionVerifierTestSuite) SetupTest() {
	s.hashFunctionMock = NewHashFunctionMock(s.T())
	s.modelDefinitionVerifier = NewModelDefinitionVerifier(s.hashFunctionMock)
}

func (s *ModelDefinitionVerifierTestSuite) TestVerify() {

	folder := rand.Text()
	hashPrefix := rand.Text()

	tt := []struct {
		name              string
		definition        MigrationDefinition
		hashes            []string
		hashFunctionError error
		expectedResult    bool
	}{
		{
			name:           "emptyDefinition",
			definition:     MigrationDefinition{},
			expectedResult: true,
		},
		{
			name: "validDefinition",
			definition: MigrationDefinition{
				Steps: []MigrationStep{
					{
						Filename: rand.Text(),
						Hash:     hashPrefix + "0",
					},
					{
						Filename: rand.Text(),
						Hash:     hashPrefix + "1",
					},
					{
						Filename: rand.Text(),
						Hash:     hashPrefix + "2",
					},
				},
			},
			hashes:         []string{hashPrefix + "0", hashPrefix + "1", hashPrefix + "2"},
			expectedResult: true,
		},
		{
			name: "invalidDefinition",
			definition: MigrationDefinition{
				Steps: []MigrationStep{
					{
						Filename: rand.Text(),
						Hash:     hashPrefix + "0",
					},
					{
						Filename: rand.Text(),
						Hash:     hashPrefix + "1",
					},
					{
						Filename: rand.Text(),
						Hash:     hashPrefix + "2",
					},
				},
			},
			hashes:         []string{hashPrefix + "0", rand.Text(), hashPrefix + "2"},
			expectedResult: false,
		},
		{
			name: "hashFunctionError",
			definition: MigrationDefinition{
				Steps: []MigrationStep{
					{
						Filename: rand.Text(),
						Hash:     hashPrefix + "0",
					},
				},
			},
			hashes:            []string{""},
			hashFunctionError: errors.New(rand.Text()),
			expectedResult:    false,
		},
	}

	for _, tc := range tt {
		s.Run(tc.name, func() {

			expectedError := tc.hashFunctionError

			prevHash := initialHash

			for i, step := range tc.definition.Steps {
				expectedHash := tc.hashes[i]
				s.hashFunctionMock.EXPECT().CalculateHash(path.Join(folder, step.Filename), prevHash).Return(expectedHash, tc.hashFunctionError)
				prevHash = expectedHash
			}

			ok, err := s.modelDefinitionVerifier.Verify(folder, tc.definition)

			s.Require().ErrorIs(err, expectedError)
			s.Require().Equal(tc.expectedResult, ok)

		})
	}

}

func TestModelDefinitionVerifier(t *testing.T) {
	suite.Run(t, new(ModelDefinitionVerifierTestSuite))
}
