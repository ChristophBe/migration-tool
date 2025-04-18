package cmd

import (
	"crypto/rand"
	"errors"
	"github.com/stretchr/testify/suite"
	"testing"
)

type RunCommandTestSuite struct {
	suite.Suite
	actionsMock *ActionsMock
}

func (s *RunCommandTestSuite) SetupTest() {
	s.actionsMock = NewActionsMock(s.T())
	acts = s.actionsMock

	rootCmd.SetArgs([]string{"run"})
}

func (s *RunCommandTestSuite) TestSuccessfulRun() {

	expectedFolder := rand.Text()
	s.actionsMock.EXPECT().Run(expectedFolder).Return(nil)
	baseFolder = expectedFolder

	err := rootCmd.Execute()

	s.NoError(err)
}

func (s *RunCommandTestSuite) TestRunError() {
	expectedError := errors.New(rand.Text())
	expectedFolder := rand.Text()
	s.actionsMock.EXPECT().Run(expectedFolder).Return(expectedError)
	baseFolder = expectedFolder

	err := rootCmd.Execute()

	s.ErrorIs(err, expectedError)
}

func TestRunCommand(t *testing.T) {
	suite.Run(t, new(RunCommandTestSuite))
}
