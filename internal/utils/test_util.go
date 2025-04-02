package utils

func GetNotNilError(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}

	}
	return nil
}
