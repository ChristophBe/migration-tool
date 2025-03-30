package actions

type HashFunction interface {
	CalculateHash(filename, previousHash string) (string, error)
}
