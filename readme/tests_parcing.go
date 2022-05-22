package readme

import (
	"encoding/json"
	"log"
	"strings"
)

type outputRow struct {
	Package string
	Output  string
}

func GetPackages(commandRunner func(cmd string, args ...string) string) (result []Package) {
	testsOutput := commandRunner("go", "test", "-cover", "-json", "./...")

	lines := strings.Split(strings.TrimSpace(testsOutput), "\n")

	for _, line := range lines {
		row := outputRow{}

		err := json.Unmarshal([]byte(line), &row)
		if err != nil {
			log.Fatalln("test parsing:", err)
		}

		path := strings.TrimPrefix(row.Package, "github.com/golang-learning/")

		// skip generator
		if strings.HasPrefix(path, "readme") {
			continue
		}

		// skip not coverage
		if !strings.HasPrefix(row.Output, "coverage: ") {
			continue
		}

		coverage := strings.TrimPrefix(row.Output, "coverage: ")
		coverage = strings.TrimSuffix(coverage, " of statements\n")

		packageDoc, functions := docs(commandRunner, path)

		result = append(result, Package{
			Package:      path,
			TestCoverage: coverage,
			Description:  packageDoc,
			Functions:    functions,
		})
	}

	return result
}

func docs(commandRunner func(cmd string, args ...string) string, path string) (string, []Function) {
	doc := commandRunner("go", "doc", "-all", "-short", path)
	packageDoc, funcsDoc, _ := strings.Cut(doc, "FUNCTIONS")

	funcsDoc = strings.TrimSpace(funcsDoc)

	var funcs []Function

	for _, line := range strings.Split(funcsDoc, "\n") {
		if line = strings.TrimSpace(line); line == "" {
			continue
		}

		if strings.HasPrefix(line, "func ") {
			funcs = append(funcs, Function{
				Interface: line,
			})

			continue
		}

		funcs[len(funcs)-1].DocLines = append(funcs[len(funcs)-1].DocLines, line)
	}

	return strings.TrimSpace(packageDoc), funcs
}
