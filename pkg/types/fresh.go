package types

import "github.com/chewxy/hm"

type Fresh struct {
	count int
}

const letters = `abcdefghijklmnopqrstuvwxyz`

func NewFresh() *Fresh {
	return &Fresh{
		count: 0,
	}
}

func (f *Fresh) Fresh() hm.TypeVariable {
	retVal := letters[f.count]
	f.count++
	return hm.TypeVariable(retVal)
}
