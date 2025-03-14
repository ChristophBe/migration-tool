package main

import (
	"crypto/rand"
	"errors"
	"github.com/stretchr/testify/suite"
	"testing"
)

type RunCommandTestSuite struct {
	suite.Suite
	mockAction *ActionsMock
}

func (s *RunCommandTestSuite) SetupTest() {
	s.mockAction = NewActionsMock(s.T())
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
			s.mockAction.EXPECT().Run(expectedFolder).Return(tc.err)
			err := RunCommands(s.mockAction, "run", expectedFolder)
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
			err := RunCommands(s.mockAction, "verify", expectedFolder)
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
			err := RunCommands(s.mockAction, "recalculate-hashes", expectedFolder)

			s.ErrorIs(err, tc.err)

		})
	}
}

func TestRunCommand(t *testing.T) {
	suite.Run(t, new(RunCommandTestSuite))
}
