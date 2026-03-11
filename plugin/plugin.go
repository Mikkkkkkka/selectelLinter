package plugin

import (
	"selectelLinter/internal/analyzer/loglint"

	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("loglint", New)
}

type Plugin struct{}

func New(_ any) (register.LinterPlugin, error) {
	return &Plugin{}, nil
}

func (p *Plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{loglint.Analyzer}, nil
}

func (p *Plugin) GetLoadMode() string {
	return register.LoadModeSyntax
}
