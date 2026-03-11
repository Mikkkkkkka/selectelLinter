package astx

import "go/ast"

func WalkCalls(files []*ast.File, fn func(*ast.CallExpr)) {
	for _, file := range files {
		ast.Inspect(file, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			fn(call)
			return true
		})
	}
}
