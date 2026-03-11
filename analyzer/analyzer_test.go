package analyzer_test

import (
	"path/filepath"
	"testing"

	"selectelLinter/analyzer"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer_LowercaseStartRule(t *testing.T) {
	testdata := testdataPath(t)

	analysistest.Run(
		t,
		testdata,
		analyzer.Analyzer,
		"lowercase",
	)
}

func TestAnalyzer_EnglishOnlyRule(t *testing.T) {
	testdata := testdataPath(t)

	analysistest.Run(
		t,
		testdata,
		analyzer.Analyzer,
		"english",
	)
}

func TestAnalyzer_NoSpecialCharsRule(t *testing.T) {
	testdata := testdataPath(t)

	analysistest.Run(
		t,
		testdata,
		analyzer.Analyzer,
		"special",
	)
}

func TestAnalyzer_NoSensitiveDataRule(t *testing.T) {
	testdata := testdataPath(t)

	analysistest.Run(
		t,
		testdata,
		analyzer.Analyzer,
		"sensitive",
	)
}

func testdataPath(t *testing.T) string {
	abs, err := filepath.Abs(filepath.Join("..", "testdata"))
	if err != nil {
		t.Fatalf("failed to resolve testdata path: %v", err)
	}

	return abs
}
