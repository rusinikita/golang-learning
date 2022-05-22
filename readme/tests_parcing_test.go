package readme_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/golang-learning/readme"
)

func Test_parseResult(t *testing.T) {
	t.Parallel()

	testOutput := `
{"Time":"2022-05-21T18:18:20.580369+02:00","Action":"output","Package":"github.com/golang-learning","Output":"?   \tgithub.com/golang-learning\t[no test files]\n"}
{"Time":"2022-05-21T18:18:20.58051+02:00","Action":"skip","Package":"github.com/golang-learning","Elapsed":0}
{"Time":"2022-05-21T18:18:20.580517+02:00","Action":"output","Package":"github.com/golang-learning/gen","Output":"?   \tgithub.com/golang-learning/gen\t[no test files]\n"}
{"Time":"2022-05-21T18:18:20.58052+02:00","Action":"skip","Package":"github.com/golang-learning/gen","Elapsed":0}
{"Time":"2022-05-21T18:18:20.642424+02:00","Action":"run","Package":"github.com/golang-learning/hellogo","Test":"TestHello"}
{"Time":"2022-05-21T18:18:20.642455+02:00","Action":"output","Package":"github.com/golang-learning/hellogo","Test":"TestHello","Output":"=== RUN   TestHello\n"}
{"Time":"2022-05-21T18:18:20.642465+02:00","Action":"output","Package":"github.com/golang-learning/hellogo","Test":"TestHello","Output":"=== PAUSE TestHello\n"}
{"Time":"2022-05-21T18:18:20.642467+02:00","Action":"pause","Package":"github.com/golang-learning/hellogo","Test":"TestHello"}
{"Time":"2022-05-21T18:18:20.642469+02:00","Action":"cont","Package":"github.com/golang-learning/hellogo","Test":"TestHello"}
{"Time":"2022-05-21T18:18:20.642471+02:00","Action":"output","Package":"github.com/golang-learning/hellogo","Test":"TestHello","Output":"=== CONT  TestHello\n"}
{"Time":"2022-05-21T18:18:20.642477+02:00","Action":"output","Package":"github.com/golang-learning/hellogo","Test":"TestHello","Output":"--- PASS: TestHello (0.00s)\n"}
{"Time":"2022-05-21T18:18:20.6425+02:00","Action":"pass","Package":"github.com/golang-learning/hellogo","Test":"TestHello","Elapsed":0}
{"Time":"2022-05-21T18:18:20.642504+02:00","Action":"output","Package":"github.com/golang-learning/hellogo","Output":"PASS\n"}
{"Time":"2022-05-21T18:18:20.642505+02:00","Action":"output","Package":"github.com/golang-learning/hellogo","Output":"coverage: 100.0% of statements\n"}
{"Time":"2022-05-21T18:18:20.642507+02:00","Action":"output","Package":"github.com/golang-learning/hellogo","Output":"ok  \tgithub.com/golang-learning/hellogo\t(cached)\tcoverage: 100.0% of statements\n"}
{"Time":"2022-05-21T18:18:20.642518+02:00","Action":"pass","Package":"github.com/golang-learning/hellogo","Elapsed":0}
{"Time":"2022-05-21T18:18:20.643019+02:00","Action":"output","Package":"github.com/golang-learning/readme","Output":"?   \tgithub.com/golang-learning/readme\t[no test files]\n"}
{"Time":"2022-05-21T18:18:20.643033+02:00","Action":"skip","Package":"github.com/golang-learning/readme","Elapsed":0}
{"Time":"2022-05-21T18:18:20.643058+02:00","Action":"output","Package":"github.com/golang-learning/readme/docsparsing","Output":"?   \tgithub.com/golang-learning/readme/docsparsing\t[no test files]\n"}
{"Time":"2022-05-21T18:18:20.643061+02:00","Action":"skip","Package":"github.com/golang-learning/readme/docsparsing","Elapsed":0}
{"Time":"2022-05-21T18:18:21.135834+02:00","Action":"run","Package":"github.com/golang-learning/readme/testparsing","Test":"Test_parseResult"}
{"Time":"2022-05-21T18:18:21.135937+02:00","Action":"output","Package":"github.com/golang-learning/readme/testparsing","Test":"Test_parseResult","Output":"=== RUN   Test_parseResult\n"}
{"Time":"2022-05-21T18:18:21.135974+02:00","Action":"output","Package":"github.com/golang-learning/readme/testparsing","Test":"Test_parseResult","Output":"=== PAUSE Test_parseResult\n"}
{"Time":"2022-05-21T18:18:21.135978+02:00","Action":"pause","Package":"github.com/golang-learning/readme/testparsing","Test":"Test_parseResult"}
{"Time":"2022-05-21T18:18:21.135986+02:00","Action":"cont","Package":"github.com/golang-learning/readme/testparsing","Test":"Test_parseResult"}
{"Time":"2022-05-21T18:18:21.135988+02:00","Action":"output","Package":"github.com/golang-learning/readme/testparsing","Test":"Test_parseResult","Output":"=== CONT  Test_parseResult\n"}
{"Time":"2022-05-21T18:18:21.136048+02:00","Action":"output","Package":"github.com/golang-learning/readme/testparsing","Test":"Test_parseResult","Output":"--- PASS: Test_parseResult (0.00s)\n"}
{"Time":"2022-05-21T18:18:21.136056+02:00","Action":"pass","Package":"github.com/golang-learning/readme/testparsing","Test":"Test_parseResult","Elapsed":0}
{"Time":"2022-05-21T18:18:21.136064+02:00","Action":"output","Package":"github.com/golang-learning/readme/testparsing","Output":"PASS\n"}
{"Time":"2022-05-21T18:18:21.136184+02:00","Action":"output","Package":"github.com/golang-learning/readme/testparsing","Output":"coverage: 73.7% of statements\n"}
{"Time":"2022-05-21T18:18:21.136619+02:00","Action":"output","Package":"github.com/golang-learning/readme/testparsing","Output":"ok  \tgithub.com/golang-learning/readme/testparsing\t0.279s\tcoverage: 73.7% of statements\n"}
{"Time":"2022-05-21T18:18:21.137851+02:00","Action":"pass","Package":"github.com/golang-learning/readme/testparsing","Elapsed":0.28}
`
	docsOutput := `Package hellogo contains first steps in language.

FUNCTIONS

func Hello() string
    Hello is first function.

func Hello() string
    Hello is first function.
	Hello is first function.

`

	mockRunner := func(cmd string, args ...string) string {
		switch args[0] {
		case "test":
			return testOutput
		case "doc":
			return docsOutput
		}

		panic("unexpected input")
	}

	expected := []readme.Package{{
		Package:      "hellogo",
		TestCoverage: "100.0%",
		Description:  "Package hellogo contains first steps in language.",
		Functions: []readme.Function{
			{
				Interface: "func Hello() string",
				DocLines:  []string{"Hello is first function."},
			},
			{
				Interface: "func Hello() string",
				DocLines:  []string{"Hello is first function.", "Hello is first function."},
			},
		},
	}}

	require.Equal(t, expected, readme.GetPackages(mockRunner))
}
