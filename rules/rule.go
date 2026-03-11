package rules

import (
	"go/ast"

	slogrules "selectelLinter/rules/slog"
	zaprules "selectelLinter/rules/zap"

	"golang.org/x/tools/go/analysis"
)

type Rule interface {
	Name() string
	Check(pass *analysis.Pass, call *ast.CallExpr, message string)
}

func Default() []Rule {
	return []Rule{
		slogrules.NewLowercaseStartRule(),
		zaprules.NewLowercaseStartRule(),
	}
}

func ApplyAll(pass *analysis.Pass, ruleSet []Rule, call *ast.CallExpr, message string) {
	for _, rule := range ruleSet {
		rule.Check(pass, call, message)
	}
}
