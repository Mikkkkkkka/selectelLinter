package special

import (
	"go/ast"

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
	return "slog-no-special-chars"
}

func (r *SlogRule) Check(pass *analysis.Pass, call *ast.CallExpr, message astx.ExtractedMessage) {
	if !loggers.IsSlogCall(pass, call) {
		return
	}

	if hasSpecialChars(message.Text) {
		diagnostics.Report(pass, call.Lparen, "log message must not contain special characters or emoji")
	}
}
