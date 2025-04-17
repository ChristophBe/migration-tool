package actions

import "path"

func (a *Actions) AddStepFile(folder, filename string) error {

	definitionFile := path.Join(folder, migrationFileName)
	migrationDefinition, err := a.definitionReaderWriter.Read(definitionFile)
	if err != nil {
		return err
	}

	ok, err := a.definitionVerifier.Verify(folder, migrationDefinition)
	if err != nil {
		return err
	}
	if !ok {
		return NewInvalidModelDefinitionError()
	}

	lastHash := migrationDefinition.Steps[len(migrationDefinition.Steps)-1].Hash

	newHash, err := a.hashFunction.CalculateHash(path.Join(folder, filename), lastHash)

	if err != nil {
		return err
	}

	migrationDefinition.Steps = append(migrationDefinition.Steps, MigrationStep{
		Filename: filename,
		Hash:     newHash,
	})

	err = a.definitionReaderWriter.Write(definitionFile, migrationDefinition)
	if err != nil {
		return err
	}

	return nil
}
