package english

import (
	"go/ast"
	"unicode"

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
	return "zap-english-only"
}

func (r *ZapRule) Check(pass *analysis.Pass, call *ast.CallExpr, message astx.ExtractedMessage) {
	if !loggers.IsZapCall(pass, call) {
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
