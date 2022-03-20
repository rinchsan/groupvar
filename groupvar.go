package groupvar

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// Analyzer finds low-readability variable/constant declarations.
var Analyzer = &analysis.Analyzer{
	Name: "groupvar",
	Doc:  "groupvar finds low-readability variable/constant declarations",
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.GenDecl)(nil),
	}
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		decl, ok := n.(*ast.GenDecl)
		if !ok {
			return
		}
		for _, spec := range decl.Specs {
			spec, ok := spec.(*ast.ValueSpec)
			if !ok {
				continue
			}
			if len(spec.Names) > 3 {
				pass.Reportf(spec.Pos(), "grouped %s declarations count should be up to 3", decl.Tok.String())
				continue
			}

			kind := token.ILLEGAL
			for _, value := range spec.Values {
				value, ok := value.(*ast.BasicLit)
				if !ok {
					continue
				}
				if kind == token.ILLEGAL {
					kind = value.Kind
					continue
				}
				if kind != value.Kind {
					pass.Reportf(value.Pos(), "grouped %s declarations should be separated by types", decl.Tok.String())
					break
				}
			}
		}
	})
	return nil, nil
}
