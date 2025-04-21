package verifiers

type HashFunction interface {
	CalculateHash(filename, previousHash string) (string, error)
}
