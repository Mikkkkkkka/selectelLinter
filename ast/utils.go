package astx

import (
	"go/ast"
	"go/token"
)

type ExtractedMessage struct {
	Text string
	Vars []string
}

func ExtractStaticMessage(call *ast.CallExpr) (ExtractedMessage, bool) {
	if len(call.Args) == 0 {
		return ExtractedMessage{}, false
	}

	return extractStringLiteral(call.Args[0])
}

func extractStringLiteral(expr ast.Expr) (ExtractedMessage, bool) {
	switch v := expr.(type) {
	case *ast.BasicLit:
		if v.Kind != token.STRING {
			return ExtractedMessage{}, false
		}

		text, ok := unquote(v.Value)
		if !ok {
			return ExtractedMessage{}, false
		}

		return ExtractedMessage{Text: text}, true
	case *ast.Ident:
		return ExtractedMessage{Vars: []string{v.Name}}, true
	case *ast.BinaryExpr:
		if v.Op != token.ADD {
			return ExtractedMessage{}, false
		}

		left, ok := extractStringLiteral(v.X)
		if !ok {
			return ExtractedMessage{}, false
		}

		right, ok := extractStringLiteral(v.Y)
		if !ok {
			return ExtractedMessage{}, false
		}

		return ExtractedMessage{
			Text: left.Text + right.Text,
			Vars: append(left.Vars, right.Vars...),
		}, true
	default:
		return ExtractedMessage{}, false
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
