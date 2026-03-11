package english

import (
	"go/ast"
	"unicode"

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
	return "slog-english-only"
}

func (r *SlogRule) Check(pass *analysis.Pass, call *ast.CallExpr, message astx.ExtractedMessage) {
	if !loggers.IsSlogCall(pass, call) {
		return
	}

	for _, ch := range message.Text {
		if !unicode.IsLetter(ch) {
			continue
		}

		if !unicode.In(ch, unicode.Latin) {
			diagnostics.Report(pass, call.Lparen, "log message must be in English")
			return
		}
	}
}
