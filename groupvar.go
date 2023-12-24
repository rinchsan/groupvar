package groupvar

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
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
				message := fmt.Sprintf("grouped %s declarations count should be up to 3", decl.Tok.String())
				diag, err := generateDiagnostic(pass, message, decl, spec)
				if err != nil {
					continue
				}
				pass.Report(diag)
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
					message := fmt.Sprintf("grouped %s declarations should be separated by types", decl.Tok.String())
					diag, err := generateDiagnostic(pass, message, decl, spec)
					if err != nil {
						continue
					}
					pass.Report(diag)
					break
				}
			}
		}
	})
	return nil, nil
}

func generateDiagnostic(pass *analysis.Pass, message string, decl *ast.GenDecl, spec *ast.ValueSpec) (analysis.Diagnostic, error) {
	diag := analysis.Diagnostic{
		Message: message,
	}

	fix := analysis.SuggestedFix{
		Message: message,
	}
	if decl.Lparen.IsValid() {
		diag.Pos = spec.Pos()
		diag.End = spec.End()

		specs := splitValueSpec(spec)
		texts := make([][]byte, len(specs))
		for i := range specs {
			text, err := formatNode(pass.Fset, specs[i])
			if err != nil {
				return analysis.Diagnostic{}, err
			}
			texts[i] = text
		}
		fix.TextEdits = append(fix.TextEdits, analysis.TextEdit{
			Pos:     spec.Pos(),
			End:     spec.End(),
			NewText: bytes.Join(texts, []byte("\n")),
		})
	} else {
		diag.Pos = decl.Pos()
		diag.End = decl.End()

		lparenPos := token.Pos(len("var"))
		if decl.Tok == token.CONST {
			lparenPos = token.Pos(len("const"))
		}
		newDecl := &ast.GenDecl{
			TokPos: decl.TokPos,
			Tok:    decl.Tok,
			Lparen: decl.TokPos + lparenPos,
			Specs:  splitValueSpec(spec),
		}
		newText, err := formatNode(pass.Fset, newDecl)
		if err != nil {
			return analysis.Diagnostic{}, err
		}
		fix.TextEdits = append(fix.TextEdits, analysis.TextEdit{
			Pos:     decl.Pos(),
			End:     decl.End(),
			NewText: newText,
		})
	}
	diag.SuggestedFixes = append(diag.SuggestedFixes, fix)

	return diag, nil
}

func splitValueSpec(spec *ast.ValueSpec) []ast.Spec {
	specs := make([]ast.Spec, len(spec.Names))
	for i := range spec.Names {
		vSpec := &ast.ValueSpec{
			Names: []*ast.Ident{spec.Names[i]},
			Type:  spec.Type,
		}
		if spec.Values != nil {
			vSpec.Values = []ast.Expr{spec.Values[i]}
		}
		specs[i] = vSpec
	}
	return specs
}

func formatNode(fset *token.FileSet, node any) ([]byte, error) {
	var buf bytes.Buffer
	if err := format.Node(&buf, fset, node); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
