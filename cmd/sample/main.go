package main

import (
	"go/ast"
	"go/token"

	"github.com/otakakot/gofixit"
)

func main() {
	gofixit.Run(func(n ast.Node, _ *token.FileSet) bool {
		lit, ok := n.(*ast.BasicLit)
		if !ok || lit.Kind != token.STRING {
			return true
		}

		lit.Value = `"sample"`

		return true
	})
}
