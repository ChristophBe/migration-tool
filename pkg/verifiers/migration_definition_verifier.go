package verifiers

import (
	"fmt"
	"github.com/ChristophBe/migration-tool/pkg/actions"
	"path/filepath"
)

type MigrationDefinition = actions.MigrationDefinition
type MigrationStep = actions.MigrationStep

const initialHash = ""

type ModelDefinitionVerifier struct {
	hashFunction HashFunction
}

func NewModelDefinitionVerifier(hashFunction HashFunction) *ModelDefinitionVerifier {
	return &ModelDefinitionVerifier{
		hashFunction: hashFunction,
	}
}

func (v *ModelDefinitionVerifier) Verify(folder string, migrationDefinition MigrationDefinition) (bool, error) {
	prevHash := initialHash
	ok := true
	for _, migration := range migrationDefinition.Steps {
		scriptPath := filepath.Join(folder, migration.Filename)
		hash, err := v.hashFunction.CalculateHash(scriptPath, prevHash)
		if err != nil {
			return false, err
		}
		if hash != migration.Hash {
			fmt.Printf("Warning: %s has changed!\n", migration.Filename)
			ok = false
		} else {
			fmt.Printf("%s is unchanged.\n", migration.Filename)
		}
		prevHash = hash
	}
	return ok, nil
}
