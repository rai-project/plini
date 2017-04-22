package types

import "github.com/chewxy/hm"

const (
	Bit     hm.TypeConst = "Bit"
	Boolean hm.TypeConst = "Boolean"
	Int     hm.TypeConst = "Int"
	Int8    hm.TypeConst = "Int8"
	Int16   hm.TypeConst = "Int16"
	Int32   hm.TypeConst = "Int32"
	Int64   hm.TypeConst = "Int64"
	Float   hm.TypeConst = "Float"
	Float32 hm.TypeConst = "Float32"
	Float64 hm.TypeConst = "Float64"
	String  hm.TypeConst = "String"
)

func NewList(t hm.Type) hm.Type {
	return NewConstructor(hm.TypeConst("List"), t)
}
