package actions

import (
	"crypto/rand"
	"fmt"
	"github.com/stretchr/testify/suite"
	"os"
	"path"
	"strconv"
	"testing"
)

type ActionRunTestSuite struct {
	suite.Suite
	inputFolder   string
	outFolder     string
	runOutputFile string
	actions       *Actions
}

func (s *ActionRunTestSuite) SetupTest() {
	s.actions = new(Actions)

	var err error
	s.inputFolder, err = os.MkdirTemp("", "test-input")
	s.NoError(err)

	s.outFolder, err = os.MkdirTemp("", "test-output")
	s.NoError(err)

	s.runOutputFile = path.Join(s.outFolder, "run-output.txt")

}
func (s *ActionRunTestSuite) addMirgrationConfig() (MigrationDefinition, string) {
	migrationDefinition := MigrationDefinition{}
	var expectedOutput string
	for i := 0; i < 2; i++ {
		filename := rand.Text() + ".sh"

		runOutput := fmt.Sprintf("step-%d-%s", i, rand.Text())
		expectedOutput = fmt.Sprintf("%s%s\n", expectedOutput, runOutput)
		content := `
#!/bin/bash
echo "` + runOutput + `" >> ` + s.runOutputFile + `
`

		s.NoError(os.WriteFile(path.Join(s.inputFolder, filename), []byte(content), 0644))

		migrationDefinition.Steps = append(migrationDefinition.Steps, MigrationStep{
			Filename:    filename,
			Description: "step " + strconv.Itoa(i),
		})

	}
	s.NoError(saveMigrationDefinition(s.inputFolder, &migrationDefinition))
	s.NoError(s.actions.RecalculateHashes(s.inputFolder))
	return migrationDefinition, expectedOutput
}

func (s *ActionRunTestSuite) TestRun() {
	_, expectedOutput := s.addMirgrationConfig()

	err := s.actions.Run(s.inputFolder)
	s.NoError(err)
	s.FileExists(path.Join(s.outFolder, outputFileName))
	s.FileExists(s.runOutputFile)

	runOutputs, err := os.ReadFile(s.runOutputFile)
	s.NoError(err)
	s.Equal(expectedOutput, string(runOutputs))
}

func (s *ActionRunTestSuite) TearDownTest() {
	s.NoError(os.RemoveAll(s.inputFolder))
	s.NoError(os.RemoveAll(s.outFolder))
}

func TestActions_Run(t *testing.T) {
	suite.Run(t, new(ActionRunTestSuite))
}
