package astx

import "go/ast"

func Selector(call *ast.CallExpr) (*ast.SelectorExpr, bool) {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	return sel, ok
}

func SelectorIdent(sel *ast.SelectorExpr) (*ast.Ident, bool) {
	ident, ok := sel.X.(*ast.Ident)
	return ident, ok
}
