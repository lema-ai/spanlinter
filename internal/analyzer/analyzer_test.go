package analyzer_test

import (
	"testing"

	"github.com/lema-ai/spanlinter/internal/analyzer"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	// Order is load-bearing: b's expected diagnostic names a's call site.
	analysistest.Run(t, testdata, analyzer.Analyzer, "github.com/lema.ai/lemmata/services/a")
	analysistest.Run(t, testdata, analyzer.Analyzer, "github.com/lema.ai/lemmata/services/b")
}
