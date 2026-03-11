package main

import (
	"selectelLinter/internal/analyzer/loglint"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(loglint.Analyzer)
}
