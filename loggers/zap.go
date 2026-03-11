package loggers

import (
	"go/ast"
	"go/types"

	astx "selectelLinter/ast"

	"golang.org/x/tools/go/analysis"
)

func isZapPackageCall(pass *analysis.Pass, sel *ast.SelectorExpr) bool {
	ident, ok := astx.SelectorIdent(sel)
	if !ok {
		return false
	}

	pkgName, ok := pass.TypesInfo.Uses[ident].(*types.PkgName)
	if !ok {
		return false
	}

	if pkgName.Imported().Path() != "go.uber.org/zap" {
		return false
	}

	return isZapLevelMethod(sel.Sel.Name)
}

func isZapMethodCall(pass *analysis.Pass, sel *ast.SelectorExpr) bool {
	selection := pass.TypesInfo.Selections[sel]
	if selection == nil {
		return false
	}

	recv := selection.Recv()
	if recv == nil {
		return false
	}

	named, ok := deref(recv).(*types.Named)
	if !ok {
		return false
	}

	if named.Obj().Pkg() == nil || named.Obj().Pkg().Path() != "go.uber.org/zap" {
		return false
	}

	if !isZapLoggerType(named.Obj().Name()) {
		return false
	}

	return isZapLevelMethod(sel.Sel.Name)
}

func isZapLevelMethod(name string) bool {
	switch name {
	case "Debug", "Info", "Warn", "Error", "DPanic", "Panic", "Fatal":
		return true
	default:
		return false
	}
}

func isZapLoggerType(name string) bool {
	switch name {
	case "Logger", "SugaredLogger":
		return true
	default:
		return false
	}
}

func deref(t types.Type) types.Type {
	if p, ok := t.(*types.Pointer); ok {
		return p.Elem()
	}
	return t
}
