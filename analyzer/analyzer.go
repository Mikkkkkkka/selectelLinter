package analyzer

import "golang.org/x/tools/go/analysis"

var Analyzer = &analysis.Analyzer{
	Name: "loglint",
	Doc:  "Checks log messages for style and safety rules",
	Run:  run,
}
