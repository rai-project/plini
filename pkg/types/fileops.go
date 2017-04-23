package types

import "github.com/chewxy/hm"

var (
	Path = hm.Generalize(hm.SimpleEnv{},
		hm.NewFnType(
			File{Path: hm.TypeVariable('f')},
			hm.TypeVariable('f'),
		),
	)
	Reader = hm.TypeConst("Reader")
	Read   = hm.Generalize(hm.SimpleEnv{},
		hm.NewFnType(
			File{Path: hm.TypeVariable('f')},
			NewReader(File{Path: hm.TypeVariable('f')}),
		),
	)
	Lines = hm.Generalize(hm.SimpleEnv{},
		hm.NewFnType(
			NewReader(File{Path: hm.TypeVariable('f')}),
			NewList(String),
		),
	)
	Words = hm.Generalize(hm.SimpleEnv{},
		hm.NewFnType(
			String,
			NewList(String),
		),
	)
	Exists = hm.Generalize(hm.SimpleEnv{},
		hm.NewFnType(
			hm.TypeVariable('f'),
			Boolean,
		),
	)
)

func NewReader(t hm.Type) hm.Type {
	return NewConstructor(Reader, t)
}

var fileOps = map[string]*hm.Scheme{
	"Path":   Path,
	"Read":   Read,
	"Lines":  Lines,
	"Words":  Words,
	"Exists": Exists,
}
