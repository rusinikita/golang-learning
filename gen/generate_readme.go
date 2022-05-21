package main

import (
	"log"
	"os"
)

func main() {
	s := `# kgd-golang-learning
Project template for golang learning with tests
`

	err := os.WriteFile("README.md", []byte(s), os.ModeExclusive)
	if err != nil {
		log.Panicln(err)
	}

	log.Println("README.md generated")
}
