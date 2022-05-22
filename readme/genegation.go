package readme

import (
	"bytes"
	_ "embed"
	"log"
	"text/template"
)

//go:embed template.md
var readme string

func MakeContent(packages []Package) string {
	readmeTemplate := template.Must(template.New("readme").Parse(readme))

	result := &bytes.Buffer{}

	err := readmeTemplate.Execute(result, packages)
	if err != nil {
		log.Fatalln("template execution:", err)
	}

	return result.String()
}
