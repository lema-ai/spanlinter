package plugin

import (
	"github.com/golangci/plugin-module-register/register"
	"github.com/lema-ai/spanlinter/internal/analyzer"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("spanlinter", New)
}

// New creates a new instance of the spanlinter plugin
func New(settings any) (register.LinterPlugin, error) {
	return &SpanLinter{}, nil
}

// SpanLinter implements the LinterPlugin interface
type SpanLinter struct{}

// BuildAnalyzers returns the analyzers provided by this linter
func (l *SpanLinter) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{analyzer.Analyzer}, nil
}

// GetLoadMode returns the load mode required by this linter
func (l *SpanLinter) GetLoadMode() string {
	return register.LoadModeTypesInfo
}
