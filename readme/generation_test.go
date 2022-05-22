package readme_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/golang-learning/readme"
)

func TestGeneration(t *testing.T) {
	t.Parallel()

	input := []readme.Package{
		{
			Package:      "hello",
			TestCoverage: "100.0%",
			Description:  "Opa opa opa pa",
			Functions: []readme.Function{{
				Interface: "func Opa() string",
				DocLines:  []string{"Opa", "Opa"},
			}},
		},
		{
			Package:      "second",
			TestCoverage: "50.0%",
			Description:  "Ra-ta-ta-ta-ta-ta-ta",
			Functions: []readme.Function{{
				Interface: "func Bla() string",
				DocLines:  []string{"Bla"},
			}},
		},
	}

	expected := `# golang-learning

[Learn go with tests](https://quii.gitbook.io/learn-go-with-tests/) exercises repository template.

# About me

Short bio and motivation in learning golang.

# Learned lessons

### hello - 100.0%
Opa opa opa pa
<details>
  <summary><code>func Opa() string</code></summary>

    Opa
    Opa
</details>

### second - 50.0%
Ra-ta-ta-ta-ta-ta-ta
<details>
  <summary><code>func Bla() string</code></summary>

    Bla
</details>
`

	require.Equal(t, expected, readme.MakeContent(input))
}
