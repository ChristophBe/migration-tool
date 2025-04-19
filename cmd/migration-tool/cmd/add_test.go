package cmd

import (
	"crypto/rand"
	"github.com/stretchr/testify/suite"
	"testing"
)

type AddCommandTestSuite struct {
	suite.Suite
	actionsMock *ActionsMock
}

func (s *AddCommandTestSuite) SetupTest() {
	s.actionsMock = NewActionsMock(s.T())
	acts = s.actionsMock
}

func (s *AddCommandTestSuite) TestSuccessfulRun() {

	expectedFolder := rand.Text()
	expectedFileName := rand.Text()

	rootCmd.SetArgs([]string{"add", expectedFileName})

	s.actionsMock.EXPECT().AddStepFile(expectedFolder, expectedFileName).Return(nil)
	baseFolder = expectedFolder

	err := rootCmd.Execute()

	s.NoError(err)
}

func (s *AddCommandTestSuite) TestAddError() {

	expectedFolder := rand.Text()
	expectedFileName := rand.Text()

	rootCmd.SetArgs([]string{"add", expectedFileName})

	s.actionsMock.EXPECT().AddStepFile(expectedFolder, expectedFileName).Return(nil)
	baseFolder = expectedFolder

	err := rootCmd.Execute()

	s.NoError(err)
}
func (s *AddCommandTestSuite) TestAddMissingArgument() {

	expectedFolder := rand.Text()

	rootCmd.SetArgs([]string{"add"})

	baseFolder = expectedFolder

	err := rootCmd.Execute()

	s.Error(err)
}

func TestAddCommand(t *testing.T) {
	suite.Run(t, new(AddCommandTestSuite))
}
