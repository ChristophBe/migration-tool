package cmd

import (
	"crypto/rand"
	"errors"
	"github.com/stretchr/testify/suite"
	"testing"
)

type RegenerateHashesCommandTestSuite struct {
	suite.Suite
	actionsMock *ActionsMock
}

func (s *RegenerateHashesCommandTestSuite) SetupTest() {
	s.actionsMock = NewActionsMock(s.T())
	acts = s.actionsMock

	rootCmd.SetArgs([]string{"regenerate-hashes"})
}

func (s *RegenerateHashesCommandTestSuite) TestSuccessfulRun() {

	expectedFolder := rand.Text()
	s.actionsMock.EXPECT().RecalculateHashes(expectedFolder).Return(nil)
	baseFolder = expectedFolder

	err := rootCmd.Execute()

	s.NoError(err)
}

func (s *RegenerateHashesCommandTestSuite) TestWithError() {
	expectedError := errors.New(rand.Text())
	expectedFolder := rand.Text()
	s.actionsMock.EXPECT().RecalculateHashes(expectedFolder).Return(expectedError)
	baseFolder = expectedFolder

	err := rootCmd.Execute()

	s.ErrorIs(err, expectedError)
}

func TestRegenerateHashesCommand(t *testing.T) {
	suite.Run(t, new(RegenerateHashesCommandTestSuite))
}
