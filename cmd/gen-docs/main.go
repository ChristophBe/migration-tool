package main

import (
	"github.com/ChristophBe/migration-tool/cmd/migration-tool/cmd"
	"os"
)

func main() {
	docPath := "./docs/cli/"
	if len(os.Args) > 1 {
		docPath = os.Args[1]
	}

	err := cmd.GenerateDoc(docPath)
	if err != nil {
		panic(err)
	}
}
