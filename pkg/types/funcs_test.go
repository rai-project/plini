package types

import (
	"fmt"
	"testing"

	"github.com/chewxy/hm"
)

func TestFMap(t *testing.T) {
	fmt.Printf("%v\n", FMap)
	sub := hm.BorrowMSubs()
	sub.Add(hm.TypeVariable('f'), hm.TypeConst("List"))
	s := hm.Instantiate(NewFresh(), FMap).Apply(sub).(*hm.FunctionType)
	fmt.Printf("after sub = %v\n", s)
}
