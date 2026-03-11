package diagnostics

import (
	"go/token"

	"golang.org/x/tools/go/analysis"
)

func Report(pass *analysis.Pass, pos token.Pos, msg string) {
	pass.Reportf(pos, msg)
}
