package javascript

import (
	"github.com/robertkrimen/otto/ast"
	"github.com/robertkrimen/otto/parser"
)

func Parse(src string) (*ast.Program, error) {
	filename := ""
	program, err := parser.ParseFile(nil, filename, src, 0)

	return program, err
}
