package lowercase

import (
	"go/ast"
	"strings"
	"unicode"
	"unicode/utf8"

	"selectelLinter/diagnostics"
	"selectelLinter/loggers"

	"golang.org/x/tools/go/analysis"
)

type SlogRule struct{}

func NewSlogRule() *SlogRule {
	return &SlogRule{}
}

func (r *SlogRule) Name() string {
	return "slog-lowercase-start"
}

func (r *SlogRule) Check(pass *analysis.Pass, call *ast.CallExpr, message string) {
	if !loggers.IsSlogCall(pass, call) {
		return
	}

	trimmed := strings.TrimLeftFunc(message, unicode.IsSpace)
	if trimmed == "" {
		return
	}

	first, _ := utf8.DecodeRuneInString(trimmed)
	if unicode.IsLetter(first) && !unicode.IsLower(first) {
		diagnostics.Report(pass, call.Lparen, "log message should start with a lowercase letter")
	}
}
