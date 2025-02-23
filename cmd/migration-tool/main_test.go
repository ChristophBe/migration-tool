package main

import (
	"crypto/rand"
	"github.com/ChristophBe/migration-tool/internal/mocks/main_mocks"
	"github.com/stretchr/testify/suite"
	"testing"
)

type RunCommandTestSuite struct {
	suite.Suite
	mockAction *main_mocks.Actions
}

func (s *RunCommandTestSuite) SetupTest() {
	s.mockAction = main_mocks.NewActions(s.T())
}

func (s *RunCommandTestSuite) TestRun() {
	expectedFolder := rand.Text()
	expectedOutPutFolder := rand.Text()
	s.mockAction.EXPECT().Run(expectedFolder, expectedOutPutFolder).Return(nil)
	RunCommands(s.mockAction, "run", expectedFolder, expectedOutPutFolder)
}

func TestRunCommand(t *testing.T) {
	suite.Run(t, new(RunCommandTestSuite))
}
