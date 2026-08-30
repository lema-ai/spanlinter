package analyzer

import (
	"go/ast"
	"go/constant"
	"go/token"
	"go/types"
	"regexp"
	"sync"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const Doc = `spanlinter checks the span names passed to log.StartSpan.

A span name must be dotted snake_case with at least two segments
(e.g. "assessment.generate_report"), and no two call sites may use the same name.
Duplicate names make traces ambiguous in Coralogix.

Non-constant names (built at runtime) are skipped.

The duplicate check holds only for one golangci-lint run over the whole tree with
a cold cache; a cached package is never re-analyzed, so its names are not seen. A
duplicate is reported at one of the two call sites, arbitrarily, so it cannot be
suppressed with //nolint:spanlinter.
`

var Analyzer = &analysis.Analyzer{
	Name:     "spanlinter",
	Doc:      Doc,
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

var nameFormat = regexp.MustCompile(`^[a-z][a-z0-9_]*(\.[a-z0-9_]+)+$`)

// Shared across packages: golangci-lint runs every package in one process.
var (
	mu   sync.Mutex
	seen = map[string]token.Position{}
)

func run(pass *analysis.Pass) (any, error) {
	insp := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	insp.Preorder([]ast.Node{(*ast.CallExpr)(nil)}, func(n ast.Node) {
		call := n.(*ast.CallExpr)
		if len(call.Args) < 2 || !isStartSpan(pass.TypesInfo, call) {
			return
		}

		arg := call.Args[1]
		value := pass.TypesInfo.Types[arg].Value
		if value == nil || value.Kind() != constant.String {
			return
		}
		name := constant.StringVal(value)

		if !nameFormat.MatchString(name) {
			pass.Reportf(arg.Pos(), "span name %q must be dotted snake_case with at least two segments (e.g. \"assessment.generate_report\")", name)
			return
		}

		pos := pass.Fset.Position(arg.Pos())
		mu.Lock()
		previous, duplicate := seen[name]
		if !duplicate {
			seen[name] = pos
		}
		mu.Unlock()

		// Same position means the package was analyzed twice (e.g. its test variant).
		if duplicate && previous != pos {
			pass.Reportf(arg.Pos(), "duplicate log.StartSpan name %q, also used at %s", name, previous)
		}
	})

	return nil, nil
}

func isStartSpan(info *types.Info, call *ast.CallExpr) bool {
	selector, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	fn, ok := info.Uses[selector.Sel].(*types.Func)
	if !ok || fn.Name() != "StartSpan" || fn.Pkg() == nil {
		return false
	}

	switch fn.Pkg().Path() {
	case "github.com/lema.ai/lemmata/services/internal/log",
		"github.com/lema-ai/lemmata/services/internal/log":
		return true
	default:
		return false
	}
}
