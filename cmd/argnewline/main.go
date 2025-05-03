package main

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
	"strings"

	"github.com/otakakot/gofixit"
)

func main() {
	gofixit.Run(func(node ast.Node, fset *token.FileSet) bool {
		switch nn := node.(type) {
		case *ast.FuncDecl:
			return true
		case *ast.CallExpr:
			return true
		case *ast.InterfaceType:
			for _, field := range nn.Methods.List {
				ftype, ok := field.Type.(*ast.FuncType)
				if !ok || ftype.Params == nil || len(ftype.Params.List) == 0 {
					continue
				}

				lparen := ftype.Params.Opening
				rparen := ftype.Params.Closing

				posL := fset.Position(lparen)
				posR := fset.Position(rparen)

				if posL.Line != posR.Line {
					continue
				}

				buf := bytes.Buffer{}

				printer.Fprint(&buf, fset, ftype)

				sig := strings.Replace(buf.String(), "func", field.Names[0].Name, 1)

				if len(sig) <= 120 {
					continue
				}

				for i, param := range ftype.Params.List {
					param.Names[0].Name = "\n\t\t" + param.Names[0].Name
					if i == len(ftype.Params.List)-1 {
						if _, isEllipsis := param.Type.(*ast.Ellipsis); !isEllipsis {
							switch t := param.Type.(type) {
							case *ast.Ident:
								t.Name += ",\n"
							case *ast.SelectorExpr:
								if ident, ok := t.X.(*ast.Ident); ok {
									ident.Name += ",\n"
								}
							case *ast.StarExpr:
								if ident, ok := t.X.(*ast.Ident); ok {
									ident.Name += ",\n"
								}
							}
						}
					}
				}
			}

			return true
		}

		return true
	})
}
