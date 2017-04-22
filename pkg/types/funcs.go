package types

import "github.com/chewxy/hm"

var (
	Identity = hm.Generalize(hm.SimpleEnv{},
		hm.NewFnType(hm.TypeVariable('a'), hm.TypeVariable('a')),
	)
	FMap = hm.Generalize(hm.SimpleEnv{},
		hm.NewFnType(
			hm.NewFnType(hm.TypeVariable('a'), hm.TypeVariable('b')),
			NewConstructor(hm.TypeVariable('f'), hm.TypeVariable('a')),
			NewConstructor(hm.TypeVariable('f'), hm.TypeVariable('b')),
		),
	)
	Map = hm.Generalize(hm.SimpleEnv{},
		hm.NewFnType(
			hm.NewFnType(hm.TypeVariable('a'), hm.TypeVariable('b')),
			NewList(hm.TypeVariable('a')),
			NewList(hm.TypeVariable('b')),
		),
	)
	Reduce = hm.Generalize(hm.SimpleEnv{},
		hm.NewFnType(
			hm.NewFnType(hm.TypeVariable('a'), hm.TypeVariable('b'), hm.TypeVariable('a')),
			hm.TypeVariable('a'),
			NewList(hm.TypeVariable('b')),
			hm.TypeVariable('a'),
		),
	)
	Transduce = hm.Generalize(hm.SimpleEnv{},
		hm.NewFnType(
			hm.NewFnType(hm.TypeVariable('c'), hm.TypeVariable('c')),
			hm.NewFnType(hm.TypeVariable('a'), hm.TypeVariable('b'), hm.TypeVariable('a')),
			hm.TypeVariable('a'),
			NewList(hm.TypeVariable('b')),
			hm.TypeVariable('a'),
		),
	)
	Reduced = hm.Generalize(hm.SimpleEnv{},
		hm.NewFnType(
			hm.TypeVariable('a'),
			hm.NewFnType(hm.TypeVariable('b'), hm.TypeVariable('b')),
			NewList(hm.TypeVariable('c')),
			hm.TypeVariable('a'),
		),
	)
	Into = hm.Generalize(hm.SimpleEnv{},
		hm.NewFnType(hm.TypeVariable('a'), hm.TypeVariable('*')),
	)
	Concat = hm.Generalize(hm.SimpleEnv{},
		hm.NewFnType(
			NewList(NewList(hm.TypeVariable('a'))),
			NewList(hm.TypeVariable('a')),
		),
	)
	ConcatMap = hm.Generalize(hm.SimpleEnv{},
		hm.NewFnType(
			hm.NewFnType(hm.TypeVariable('a'), NewList(hm.TypeVariable('b'))),
			NewList(hm.TypeVariable('a')),
			NewList(hm.TypeVariable('b')),
		),
	)
	Chain = hm.Generalize(hm.SimpleEnv{},
		hm.NewFnType(
			hm.NewFnType(hm.TypeVariable('a'), NewConstructor(hm.TypeVariable('m'), hm.TypeVariable('b'))),
			NewConstructor(hm.TypeVariable('m'), hm.TypeVariable('a')),
			NewConstructor(hm.TypeVariable('m'), hm.TypeVariable('b')),
		),
	)
	Zip = hm.Generalize(hm.SimpleEnv{},
		hm.NewFnType(
			NewList(hm.TypeVariable('a')),
			NewList(hm.TypeVariable('b')),
			NewList(hm.NewRecordType("zipped", hm.TypeVariable('a'), hm.TypeVariable('b'))),
		),
	)
	All = hm.Generalize(hm.SimpleEnv{},
		hm.NewFnType(
			hm.NewFnType(hm.TypeVariable('a'), Boolean),
			NewList(hm.TypeVariable('a')),
			Boolean,
		),
	)
	Any = hm.Generalize(hm.SimpleEnv{},
		hm.NewFnType(
			hm.NewFnType(hm.TypeVariable('a'), Boolean),
			NewList(hm.TypeVariable('a')),
			Boolean,
		),
	)
	Pipe = hm.Generalize(hm.SimpleEnv{},
		hm.NewFnType(
			hm.NewFnType(hm.TypeVariable('a'), hm.TypeVariable('b'), hm.TypeVariable('c'), hm.TypeVariable('o')),
			hm.NewFnType(hm.TypeVariable('o'), hm.TypeVariable('p')),
			hm.NewFnType(hm.TypeVariable('p'), hm.TypeVariable('q')),
			hm.NewFnType(hm.TypeVariable('q'), hm.TypeVariable('r')),
			hm.NewFnType(hm.TypeVariable('a'), hm.TypeVariable('b'), hm.TypeVariable('c'), hm.TypeVariable('r')),
		),
	)
	Compose = hm.Generalize(hm.SimpleEnv{},
		hm.NewFnType(
			hm.NewFnType(hm.TypeVariable('q'), hm.TypeVariable('r')),
			hm.NewFnType(hm.TypeVariable('p'), hm.TypeVariable('q')),
			hm.NewFnType(hm.TypeVariable('o'), hm.TypeVariable('p')),
			hm.NewFnType(hm.TypeVariable('a'), hm.TypeVariable('b'), hm.TypeVariable('c'), hm.TypeVariable('o')),
			hm.NewFnType(hm.TypeVariable('a'), hm.TypeVariable('b'), hm.TypeVariable('c'), hm.TypeVariable('r')),
		),
	)
)
