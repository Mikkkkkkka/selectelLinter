package analyzer

import (
	"go/ast"

	astx "selectelLinter/ast"
	"selectelLinter/loggers"
	"selectelLinter/rules"

	"golang.org/x/tools/go/analysis"
)

func run(pass *analysis.Pass) (any, error) {
	ruleSet := rules.Default()

	astx.WalkCalls(pass.Files, func(call *ast.CallExpr) {
		if !loggers.IsSupported(pass, call) {
			return
		}

		msg, ok := loggers.ExtractStaticMessage(call)
		if !ok {
			return
		}

		rules.ApplyAll(pass, ruleSet, call, msg)
	})

	return nil, nil
}
