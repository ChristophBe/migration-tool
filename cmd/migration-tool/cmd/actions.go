package cmd

type Actions interface {
	Run(folder string) error
	Verify(folder string) (bool, error)
	RecalculateHashes(folder string) error
	AddStepFile(folder string, filename string, description string) error
}
