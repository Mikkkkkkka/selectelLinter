package loggers

import (
	"go/ast"

	astx "selectelLinter/ast"

	"golang.org/x/tools/go/analysis"
)

func IsSupported(pass *analysis.Pass, call *ast.CallExpr) bool {
	sel, ok := astx.Selector(call)
	if !ok {
		return false
	}

	if isSlogPackageCall(pass, sel) || isZapPackageCall(pass, sel) {
		return true
	}

	if isSlogMethodCall(pass, sel) || isZapMethodCall(pass, sel) {
		return true
	}

	return false
}

func ExtractStaticMessage(call *ast.CallExpr) (string, bool) {
	return astx.ExtractStaticMessage(call)
}
