package types

import "github.com/chewxy/hm"

var (
	Add = hm.Generalize(hm.SimpleEnv{},
		hm.NewFnType(
			hm.TypeVariable('a'),
			hm.TypeVariable('a'),
			hm.TypeVariable('a'),
		),
	)
	AddTo = hm.Generalize(hm.SimpleEnv{},
		hm.NewFnType(
			hm.TypeVariable('a'),
			hm.TypeVariable('a'),
		),
	)
	Subtract = hm.Generalize(hm.SimpleEnv{},
		hm.NewFnType(
			hm.TypeVariable('a'),
			hm.TypeVariable('a'),
			hm.TypeVariable('a'),
		),
	)
	SubtractFrom = hm.Generalize(hm.SimpleEnv{},
		hm.NewFnType(
			hm.TypeVariable('a'),
			hm.TypeVariable('a'),
		),
	)
	Divide = hm.Generalize(hm.SimpleEnv{},
		hm.NewFnType(
			hm.TypeVariable('a'),
			hm.TypeVariable('a'),
			hm.TypeVariable('a'),
		),
	)
	DivideBy = hm.Generalize(hm.SimpleEnv{},
		hm.NewFnType(
			hm.TypeVariable('a'),
			hm.TypeVariable('a'),
		),
	)
	Times = hm.Generalize(hm.SimpleEnv{},
		hm.NewFnType(
			hm.TypeVariable('a'),
			hm.TypeVariable('a'),
			hm.TypeVariable('a'),
		),
	)
	TimesBy = hm.Generalize(hm.SimpleEnv{},
		hm.NewFnType(
			hm.TypeVariable('a'),
			hm.TypeVariable('a'),
		),
	)
)

var mathOps = map[string]*hm.Scheme{
	"Add":          Add,
	"AddTo":        AddTo,
	"Subtract":     Subtract,
	"SubtractFrom": SubtractFrom,
	"Divide":       Divide,
	"DivideBy":     DivideBy,
	"Times":        Times,
	"TimesBy":      TimesBy,
}
