package main

import (
	"log"
	"os/exec"
	"strings"
)

//go:generate go run gen/generate_readme.go

// Checks all additional installations.
func main() {
	version, err := exec.Command("golangci-lint", "--version").Output()
	if err != nil {
		log.Fatal(err)
	}

	if !strings.Contains(string(version), "golangci-lint has version") {
		log.Fatal("please, install golangci-lint")
	}

	version, err = exec.Command("git", "version").Output()
	if err != nil {
		log.Fatal(err)
	}

	if !strings.Contains(string(version), "git version") {
		log.Fatal("please, install git")
	}

	log.Println("Everything installed correct")
}
