package readme

// Package represents exercises directory.
type Package struct {
	Package      string
	TestCoverage string     // TestCoverage from go test -cover
	Description  string     // Description from go doc comment above package
	Functions    []Function // Functions and their docs above
}

type Function struct {
	Interface string
	DocLines  []string
}
