package step

import "github.com/chewxy/hm"

type Step interface {
	ID() string
	Name() string
	Type() hm.Type
	SetDepends(string)
}
