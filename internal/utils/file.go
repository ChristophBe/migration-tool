package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func FileInBaseFolderCheck(base, filename string) error {
	absPath, err := filepath.Abs(filename)
	if err != nil {
		return fmt.Errorf("failed to resolve absolute path: %w", err)
	}

	if _, err := os.Stat(absPath); err != nil {
		return fmt.Errorf("file does not exist: %w", err)
	}

	baseFolderAbs, err := filepath.Abs(base)
	if err != nil {
		return fmt.Errorf("failed to resolve base folder path: %w", err)
	}

	if !strings.HasPrefix(absPath, baseFolderAbs) {
		return fmt.Errorf("file must be inside the base folder")
	}
	return nil
}

func FileExistsCheck(filename string) error {
	_, err := os.Stat(filename)
	if err != nil && errors.Is(err, os.ErrNotExist) {
		return err
	} else if err != nil {
		return fmt.Errorf("failed to stat file: %w", err)
	}
	return nil
}
