# golang-learning

[Learn go with tests](https://quii.gitbook.io/learn-go-with-tests/) exercises repository template.

# About me

Short bio and motivation in learning golang.

# Learned lessons
{{range .}}
### {{.Package}} - {{.TestCoverage}}
{{.Description}}{{range .Functions}}
<details>
  <summary><code>{{.Interface}}</code></summary>
{{range .DocLines}}
    {{.}}{{end}}
</details>
{{end}}{{end}}