package rules

import (
	"go/ast"

	astx "selectelLinter/ast"
	"selectelLinter/rules/english"
	"selectelLinter/rules/lowercase"
	"selectelLinter/rules/sensitive"
	"selectelLinter/rules/special"

	"golang.org/x/tools/go/analysis"
)

type Rule interface {
	Name() string
	Check(pass *analysis.Pass, call *ast.CallExpr, message astx.ExtractedMessage)
}

func Default() []Rule {
	return []Rule{
		lowercase.NewSlogRule(),
		lowercase.NewZapRule(),
		english.NewSlogRule(),
		english.NewZapRule(),
		special.NewSlogRule(),
		special.NewZapRule(),
		sensitive.NewSlogRule(),
		sensitive.NewZapRule(),
	}
}

func ApplyAll(pass *analysis.Pass, ruleSet []Rule, call *ast.CallExpr, message astx.ExtractedMessage) {
	for _, rule := range ruleSet {
		rule.Check(pass, call, message)
	}
}
