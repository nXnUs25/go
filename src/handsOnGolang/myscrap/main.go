package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

// parse a Go program into an AST representation.
func parse(program string) (*token.FileSet, *ast.File, error) {
	fs := token.NewFileSet()
	tree, err := parser.ParseFile(fs, "example.go", program, 0)
	if err != nil {
		return nil, nil, err
	}
	return fs, tree, nil
}

// inspectVariables visits each AST node and prints any encountered Go variable.
func inspectVariables(fs *token.FileSet, tree *ast.File) {
	ast.Inspect(tree, func(n ast.Node) bool {
		ident, ok := n.(*ast.Ident)
		if !ok || ident.Obj == nil || ident.Obj.Kind != ast.Var {
			return true
		}

		fmt.Printf("%s:\tvariable %q\n", fs.Position(n.Pos()), ident)
		return true
	})
}

func main() {
	fs, tree, err := parse(`
        package foo 

        var global = "foo"

		const (
			TESTING = "testing ASM package"
		)

        func main(){ 
			x := 42 
			y := 4.4
			var xyz string
		}
    `)

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}

	inspectVariables(fs, tree)
}
