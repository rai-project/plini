package types

import "github.com/chewxy/hm"

const (
	Bit       hm.TypeConst = "Bit"
	Boolean   hm.TypeConst = "Boolean"
	Integer   hm.TypeConst = "Integer"
	Integer8  hm.TypeConst = "Integer8"
	Integer16 hm.TypeConst = "Integer16"
	Integer32 hm.TypeConst = "Integer32"
	Integer64 hm.TypeConst = "Integer64"
	Float     hm.TypeConst = "Float"
	Float32   hm.TypeConst = "Float32"
	Float64   hm.TypeConst = "Float64"
	String    hm.TypeConst = "String"
)

func NewList(t hm.Type) hm.Type {
	return NewConstructor(hm.TypeConst("List"), t)
}
