package ir

import (
	"github.com/chewxy/hm"
	"github.com/rai-project/plini/pkg/dataset"
)

type Op interface {
	ID() string
	Container() string
	Type() hm.Type
	Run(*dataset.Dataset) (*dataset.Dataset, error)
}
