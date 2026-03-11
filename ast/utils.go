package astx

import (
	"go/ast"
	"go/token"
)

func ExtractStaticMessage(call *ast.CallExpr) (string, bool) {
	if len(call.Args) == 0 {
		return "", false
	}

	return extractStringLiteral(call.Args[0])
}

func extractStringLiteral(expr ast.Expr) (string, bool) {
	switch v := expr.(type) {
	case *ast.BasicLit:
		if v.Kind != token.STRING {
			return "", false
		}

		return unquote(v.Value)
	case *ast.BinaryExpr:
		if v.Op != token.ADD {
			return "", false
		}

		left, ok := extractStringLiteral(v.X)
		if !ok {
			return "", false
		}

		right, ok := extractStringLiteral(v.Y)
		if !ok {
			return "", false
		}

		return left + right, true
	default:
		return "", false
	}
}

func unquote(raw string) (string, bool) {
	if len(raw) < 2 {
		return "", false
	}

	if raw[0] == '`' && raw[len(raw)-1] == '`' {
		return raw[1 : len(raw)-1], true
	}

	if raw[0] == '"' && raw[len(raw)-1] == '"' {
		return raw[1 : len(raw)-1], true
	}

	return "", false
}
