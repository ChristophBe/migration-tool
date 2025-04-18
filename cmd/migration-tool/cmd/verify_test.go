package cmd

import (
	"crypto/rand"
	"errors"
	"github.com/stretchr/testify/suite"
	"testing"
)

type VerifyCommandTestSuite struct {
	suite.Suite
	actionsMock *ActionsMock
}

func (s *VerifyCommandTestSuite) SetupTest() {
	s.actionsMock = NewActionsMock(s.T())
	acts = s.actionsMock

	rootCmd.SetArgs([]string{"verify"})
}

func (s *VerifyCommandTestSuite) TestSuccessfulRun() {

	expectedFolder := rand.Text()
	s.actionsMock.EXPECT().Verify(expectedFolder).Return(true, nil)
	baseFolder = expectedFolder

	err := rootCmd.Execute()

	s.ErrorIs(err, ChangeDetectedError)
}

func (s *VerifyCommandTestSuite) TestChangesDetectedRun() {

	expectedFolder := rand.Text()
	s.actionsMock.EXPECT().Verify(expectedFolder).Return(false, nil)
	baseFolder = expectedFolder

	err := rootCmd.Execute()

	s.NoError(err)
}

func (s *VerifyCommandTestSuite) TestRunError() {
	expectedError := errors.New(rand.Text())
	expectedFolder := rand.Text()
	s.actionsMock.EXPECT().Verify(expectedFolder).Return(false, expectedError)
	baseFolder = expectedFolder

	err := rootCmd.Execute()

	s.ErrorIs(err, expectedError)
}

func TestVerifyCommand(t *testing.T) {
	suite.Run(t, new(VerifyCommandTestSuite))
}
