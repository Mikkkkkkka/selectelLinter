package special

import (
	"go/ast"

	astx "selectelLinter/ast"
	"selectelLinter/diagnostics"
	"selectelLinter/loggers"

	"golang.org/x/tools/go/analysis"
)

type ZapRule struct{}

func NewZapRule() *ZapRule {
	return &ZapRule{}
}

func (r *ZapRule) Name() string {
	return "zap-no-special-chars"
}

func (r *ZapRule) Check(pass *analysis.Pass, call *ast.CallExpr, message astx.ExtractedMessage) {
	if !loggers.IsZapCall(pass, call) {
		return
	}

	if hasSpecialChars(message.Text) {
		diagnostics.Report(pass, call.Lparen, "log message must not contain special characters or emoji")
	}
}
