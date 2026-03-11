package sensitive

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
	return "zap-no-sensitive-data"
}

func (r *ZapRule) Check(pass *analysis.Pass, call *ast.CallExpr, message astx.ExtractedMessage) {
	if !loggers.IsZapCall(pass, call) {
		return
	}

	if containsSensitive(message.Text) {
		diagnostics.Report(pass, call.Lparen, "log message must not contain sensitive data")
	}
}
