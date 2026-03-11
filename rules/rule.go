package rules

import (
	"go/ast"

	"selectelLinter/rules/english"
	"selectelLinter/rules/lowercase"

	"golang.org/x/tools/go/analysis"
)

type Rule interface {
	Name() string
	Check(pass *analysis.Pass, call *ast.CallExpr, message string)
}

func Default() []Rule {
	return []Rule{
		lowercase.NewSlogRule(),
		lowercase.NewZapRule(),
		english.NewSlogRule(),
		english.NewZapRule(),
	}
}

func ApplyAll(pass *analysis.Pass, ruleSet []Rule, call *ast.CallExpr, message string) {
	for _, rule := range ruleSet {
		rule.Check(pass, call, message)
	}
}
