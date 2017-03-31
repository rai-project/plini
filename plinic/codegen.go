package plinic

import (
	"fmt"

	. "github.com/dave/jennifer/jen"
)

func CodeGen() string {
	f := NewFile("main")
	f.Func().Id("main").Params().Block(
		Qual("fmt", "Println").Call(Lit("Hello, world")),
	)
	fmt.Printf("%#v", f)

	return ""
}
