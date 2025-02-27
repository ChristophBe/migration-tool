package main

import (
	"crypto/rand"
	"errors"
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

	tt := []struct {
		name string
		err  error
	}{
		{
			name: "no error",
		},
		{
			name: "with error",
			err:  errors.New(rand.Text()),
		},
	}
	for _, tc := range tt {
		s.Run(tc.name, func() {
			expectedFolder := rand.Text()
			expectedOutPutFolder := rand.Text()
			s.mockAction.EXPECT().Run(expectedFolder, expectedOutPutFolder).Return(tc.err)
			err := RunCommands(s.mockAction, "run", expectedFolder, expectedOutPutFolder)
			s.ErrorIs(err, tc.err)

		})
	}
}

func (s *RunCommandTestSuite) TestVerify() {

	tt := []struct {
		name            string
		err             error
		changesDetected bool
	}{
		{
			name: "no error",
		},
		{
			name: "with error",
			err:  errors.New(rand.Text()),
		},
		{
			name:            "with changes detected",
			changesDetected: true,
		},
		{
			name:            "with changes detected and error",
			changesDetected: true,
			err:             errors.New(rand.Text()),
		},
	}
	for _, tc := range tt {
		s.Run(tc.name, func() {
			expectedFolder := rand.Text()
			s.mockAction.EXPECT().Verify(expectedFolder).Return(tc.changesDetected, tc.err)
			err := RunCommands(s.mockAction, "verify", expectedFolder, rand.Text())
			if tc.changesDetected && tc.err == nil {
				s.Error(err)
			} else {
				s.ErrorIs(err, tc.err)
			}
		})
	}
}
func (s *RunCommandTestSuite) TestRecalculateHashes() {

	tt := []struct {
		name string
		err  error
	}{
		{
			name: "no error",
		},
		{
			name: "with error",
			err:  errors.New(rand.Text()),
		},
	}
	for _, tc := range tt {
		s.Run(tc.name, func() {
			expectedFolder := rand.Text()
			s.mockAction.EXPECT().RecalculateHashes(expectedFolder).Return(tc.err)
			err := RunCommands(s.mockAction, "recalculate-hashes", expectedFolder, rand.Text())

			s.ErrorIs(err, tc.err)

		})
	}
}

func TestRunCommand(t *testing.T) {
	suite.Run(t, new(RunCommandTestSuite))
}
