package sensitive

import (
	"go/ast"
	"selectelLinter/diagnostics"
	"selectelLinter/loggers"

	"golang.org/x/tools/go/analysis"
)

type SlogRule struct{}

func NewSlogRule() *SlogRule {
	return &SlogRule{}
}

func (r *SlogRule) Name() string {
	return "slog-no-sensitive-data"
}

func (r *SlogRule) Check(pass *analysis.Pass, call *ast.CallExpr, message string) {
	if !loggers.IsSlogCall(pass, call) {
		return
	}

	if containsSensitive(message) {
		diagnostics.Report(pass, call.Lparen, "log message must not contain sensitive data")
	}
}
