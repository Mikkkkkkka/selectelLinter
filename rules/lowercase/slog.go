package lowercase

import (
	"go/ast"
	"strings"
	"unicode"
	"unicode/utf8"

	astx "selectelLinter/ast"
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

func (r *SlogRule) Check(pass *analysis.Pass, call *ast.CallExpr, message astx.ExtractedMessage) {
	if !loggers.IsSlogCall(pass, call) {
		return
	}

	trimmed := strings.TrimLeftFunc(message.Text, unicode.IsSpace)
	if trimmed == "" {
		return
	}

	first, _ := utf8.DecodeRuneInString(trimmed)
	if unicode.IsLetter(first) && !unicode.IsLower(first) {
		diagnostics.Report(pass, call.Lparen, "log message should start with a lowercase letter")
	}
}
