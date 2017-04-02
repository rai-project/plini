package clike

import (
	ast "github.com/abduld/castdiff/cc/ast"
	cc "github.com/abduld/castdiff/cc/parse"
)

func Parse(s string) (*ast.Prog, error) {
	return cc.ParseProg(s)
}
