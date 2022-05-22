package main

import (
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/golang-learning/readme"
)

func main() {
	log.Println("Generation started")
	// get content
	packages := readme.GetPackages(func(cmd string, args ...string) string {
		output, err := exec.Command(cmd, args...).Output()
		if err != nil {
			log.Fatalln("command:", cmd, strings.Join(args, " "), err)
		}

		log.Println(string(output))

		return string(output)
	})

	content := readme.MakeContent(packages)

	// rewrite file
	err := os.WriteFile("README.md", []byte(content), os.ModeExclusive)
	if err != nil {
		log.Panicln(err)
	}

	log.Println("README.md generated")
}
