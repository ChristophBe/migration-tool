package cmd

import (
	"crypto/rand"
	"github.com/stretchr/testify/suite"
	"os"
	"path"
	"testing"
)

type AddCommandTestSuite struct {
	suite.Suite
	actionsMock    *ActionsMock
	expectedFolder string
	folders        []string
}

func (s *AddCommandTestSuite) SetupTest() {
	s.actionsMock = NewActionsMock(s.T())
	acts = s.actionsMock

	s.expectedFolder = s.testDir()
	baseFolder = s.expectedFolder
}

func (s *AddCommandTestSuite) TearDownTest() {

	for _, folder := range s.folders {

		s.NoError(os.RemoveAll(folder))
	}
}

func (s *AddCommandTestSuite) TestSuccessfulRun() {

	filePath := s.addTestFile(s.expectedFolder)

	rootCmd.SetArgs([]string{"add", filePath})

	s.actionsMock.EXPECT().AddStepFile(s.expectedFolder, filePath).Return(nil)

	err := rootCmd.Execute()

	s.NoError(err)
}

func (s *AddCommandTestSuite) addTestFile(dir string) string {
	fileName := rand.Text() + ".sh"
	filePath := path.Join(dir, fileName)

	err := os.WriteFile(filePath, []byte(rand.Text()), 0644)
	s.Require().NoError(err)
	return filePath
}

func (s *AddCommandTestSuite) TestAddError() {

	expectedFile := s.addTestFile(s.expectedFolder)
	rootCmd.SetArgs([]string{"add", expectedFile})

	s.actionsMock.EXPECT().AddStepFile(s.expectedFolder, expectedFile).Return(nil)

	err := rootCmd.Execute()

	s.NoError(err)
}
func (s *AddCommandTestSuite) TestAddMissingArgument() {

	rootCmd.SetArgs([]string{"add"})

	err := rootCmd.Execute()

	s.Error(err)
}

func (s *AddCommandTestSuite) TestAddInvalidFileArgument() {

	otherDir := s.testDir()

	fileInOtherDir := s.addTestFile(otherDir)

	tt := []struct {
		name     string
		filename string
	}{
		{
			name:     "file not exists",
			filename: path.Join(s.expectedFolder, rand.Text()),
		},
		{
			name:     "file in other directory",
			filename: fileInOtherDir,
		},
	}

	for _, tc := range tt {
		s.Run(tc.name, func() {
			rootCmd.SetArgs([]string{"add", tc.filename})
			err := rootCmd.Execute()
			s.Require().Error(err)
		})
	}

}

func (s *AddCommandTestSuite) testDir() string {
	dir, err := os.MkdirTemp(os.TempDir(), "migration-tool-test-add-*")
	s.Require().NoError(err)

	s.folders = append(s.folders, dir)
	return dir

}

func TestAddCommand(t *testing.T) {
	suite.Run(t, new(AddCommandTestSuite))
}
