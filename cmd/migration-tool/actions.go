package main

type Actions interface {
	Run(folder string, outputFolder string) error
	Verify(folder string) (bool, error)
	RecalculateHashes(folder string) error
}
