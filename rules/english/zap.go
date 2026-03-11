package english

import (
	"go/ast"
	"unicode"

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

func (r *ZapRule) Check(pass *analysis.Pass, call *ast.CallExpr, message string) {
	if !loggers.IsZapCall(pass, call) {
		return
	}

	for _, ch := range message {
		if !unicode.IsLetter(ch) {
			continue
		}

		if !unicode.In(ch, unicode.Latin) {
			diagnostics.Report(pass, call.Lparen, "log message must be in English")
			return
		}
	}
}
