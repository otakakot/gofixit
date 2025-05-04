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
			if nn.Type.Params == nil || len(nn.Type.Params.List) == 0 {
				return true
			}

			lparen := nn.Type.Func
			rparen := nn.Type.Results.End()

			posL := fset.Position(lparen)
			posR := fset.Position(rparen)

			if posL.Line != posR.Line {
				return true
			}

			if posR.Column <= 120 {
				return true
			}

			for i, param := range nn.Type.Params.List {
				param.Names[0].Name = "\n\t" + param.Names[0].Name
				if i == len(nn.Type.Params.List)-1 {
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

			return true
		case *ast.CallExpr:
			if !nn.Lparen.IsValid() || !nn.Rparen.IsValid() || len(nn.Args) == 0 {
				return true
			}

			posL := fset.Position(nn.Lparen)
			posR := fset.Position(nn.Rparen)

			if posL.Line != posR.Line {
				return true
			}

			if posR.Column <= 120 {
				return true
			}

			for i, arg := range nn.Args {
				switch a := arg.(type) {
				case *ast.BasicLit:
					a.Value = "\n\t" + a.Value
					if i == len(nn.Args)-1 {
						a.Value += ",\n"
					}
				}
			}

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

				if err := printer.Fprint(&buf, fset, ftype); err != nil {
					return false
				}

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
