package loggers

import (
	"go/ast"
	"go/types"

	astx "selectelLinter/ast"

	"golang.org/x/tools/go/analysis"
)

func isSlogPackageCall(pass *analysis.Pass, sel *ast.SelectorExpr) bool {
	ident, ok := astx.SelectorIdent(sel)
	if !ok {
		return false
	}

	pkgName, ok := pass.TypesInfo.Uses[ident].(*types.PkgName)
	if !ok {
		return false
	}

	if pkgName.Imported().Path() != "log/slog" {
		return false
	}

	return isSlogLevelMethod(sel.Sel.Name)
}

func isSlogMethodCall(pass *analysis.Pass, sel *ast.SelectorExpr) bool {
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

	if named.Obj().Pkg() == nil || named.Obj().Pkg().Path() != "log/slog" {
		return false
	}

	if named.Obj().Name() != "Logger" {
		return false
	}

	return isSlogLevelMethod(sel.Sel.Name)
}

func isSlogLevelMethod(name string) bool {
	switch name {
	case "Debug", "Info", "Warn", "Error":
		return true
	default:
		return false
	}
}
