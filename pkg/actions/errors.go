package actions

import "errors"

type InvalidModelDefinitionError error

func NewInvalidModelDefinitionError() InvalidModelDefinitionError {
	return errors.New("invalid migration definition")
}
