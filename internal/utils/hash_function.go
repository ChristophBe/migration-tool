package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

type HashFunction struct {
}

func NewHashFunction() *HashFunction {
	return new(HashFunction)
}

func (h HashFunction) CalculateHash(filename, previousHash string) (string, error) {
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("error reading file %s: %w", filename, err)
	}
	hasher := sha256.New()
	hasher.Write(fileContent)
	hasher.Write([]byte(previousHash))
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash, nil
}
