package loglint

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "loglint",
	Doc:  "Verifies logs style",
	Run:  run,
}

func run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {

			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			sel, ok := call.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}

			ident, ok := sel.X.(*ast.Ident)
			if !ok {
				return true
			}

			if ident.Name == "fmt" && sel.Sel.Name == "Println" {

				pass.Reportf(sel.Pos(), "avoid fmt.Println in production code")
			}

			return true
		})
	}

	return nil, nil
}
