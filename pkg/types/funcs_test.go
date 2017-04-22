package types

import (
	"fmt"
	"testing"

	"github.com/chewxy/hm"
	"github.com/stretchr/testify/assert"
)

func TestFMap(t *testing.T) {
	assert.Equal(t, "∀[a, b, f]: a → b → f(a) → f(b)", fmt.Sprintf("%v", FMap))
	ifmap := hm.Instantiate(NewFresh(), FMap)
	assert.Equal(t, "a → b → c(a) → c(b)", fmt.Sprintf("%v", ifmap))
	sub := hm.BorrowMSubs()
	sub.Add(hm.TypeVariable('x'), hm.TypeConst("List"))
	s := ifmap.Apply(sub).(*hm.FunctionType)
	assert.Equal(t, "a → b → c(a) → c(b)", fmt.Sprintf("%v", s))
}

func TestUnify(t *testing.T) {
	id := hm.Instantiate(NewFresh(), Identity)
	assert.Equal(t, "a → a", fmt.Sprintf("%v", id))
	rhs := hm.NewFnType(Integer, Integer)
	sub, err := hm.Unify(id, rhs)
	assert.NoError(t, err)
	fid := id.Apply(sub).(*hm.FunctionType)
	assert.Equal(t, "Integer → Integer", fmt.Sprintf("%v", fid))
}
