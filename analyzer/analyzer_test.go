package analyzer_test

import (
	"path/filepath"
	"testing"

	"selectelLinter/analyzer"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer_LowercaseStartRule(t *testing.T) {
	abs, err := filepath.Abs(filepath.Join("..", "testdata"))
	if err != nil {
		t.Fatalf("failed to resolve testdata path: %v", err)
	}

	analysistest.Run(
		t,
		abs,
		analyzer.Analyzer,
		"lowercase",
	)
}
