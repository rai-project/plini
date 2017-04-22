package types

import (
	"fmt"

	"github.com/chewxy/hm"
)

type File struct {
	Path hm.Type
}

func (f File) Name() string {
	return "File"
}

func (f File) String() string {
	return fmt.Sprintf("File(%#v)", f.Path)
}

func (f File) Apply(sub hm.Subs) hm.Substitutable {
	f.Path = f.Path.Apply(sub).(hm.Type)
	return f
}

func (f File) FreeTypeVar() hm.TypeVarSet {
	return f.Path.FreeTypeVar()
}

func (f File) Normalize(k, v hm.TypeVarSet) (hm.Type, error) {
	var err error
	if f.Path, err = f.Path.Normalize(k, v); err != nil {
		return nil, err
	}

	return f, nil
}

func (f File) Types() hm.Types {
	ts := hm.BorrowTypes(1)
	ts[0] = f.Path
	return ts
}

func (f File) Eq(other hm.Type) bool {
	if ot, ok := other.(File); ok {
		return f.Path.Eq(ot.Path)
	}
	return false
}

func (f File) Format(state fmt.State, c rune) {
	fmt.Fprintf(state, "File(%#v)", f.Path)
}
