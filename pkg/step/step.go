package step

import "github.com/chewxy/hm"

type Step interface {
	ID() string
	Type() hm.Type
	SetDepends(string)
}
