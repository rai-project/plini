package types

import (
	"fmt"

	"github.com/chewxy/hm"
)

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
	Select = hm.Generalize(hm.SimpleEnv{},
		hm.NewFnType(
			NewList(hm.TypeVariable('a')),
			hm.NewFnType(hm.TypeVariable('a'), Boolean),
			NewList(hm.TypeVariable('a')),
		),
	)
	Append = hm.Generalize(hm.SimpleEnv{},
		hm.NewFnType(
			NewList(hm.TypeVariable('a')),
			hm.TypeVariable('a'),
			NewList(hm.TypeVariable('a')),
		),
	)
	Prepend = hm.Generalize(hm.SimpleEnv{},
		hm.NewFnType(
			NewList(hm.TypeVariable('a')),
			hm.TypeVariable('a'),
			NewList(hm.TypeVariable('a')),
		),
	)
	Join = hm.Generalize(hm.SimpleEnv{},
		hm.NewFnType(
			NewList(hm.TypeVariable('a')),
			NewList(hm.TypeVariable('a')),
			NewList(hm.TypeVariable('a')),
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
	First = hm.Generalize(hm.SimpleEnv{},
		hm.NewFnType(
			NewList(hm.TypeVariable('a')),
			hm.TypeVariable('a'),
		),
	)
	Last = hm.Generalize(hm.SimpleEnv{},
		hm.NewFnType(
			NewList(hm.TypeVariable('a')),
			hm.TypeVariable('a'),
		),
	)
	Rest = hm.Generalize(hm.SimpleEnv{},
		hm.NewFnType(
			NewList(hm.TypeVariable('a')),
			NewList(hm.TypeVariable('a')),
		),
	)
	Partition = hm.Generalize(hm.SimpleEnv{},
		hm.NewFnType(
			NewList(hm.TypeVariable('a')),
			Integer,
			NewList(NewList(hm.TypeVariable('a'))),
		),
	)
)

var listOps = map[string]*hm.Scheme{
	"Identity":  Identity,
	"FMap":      FMap,
	"Map":       Map,
	"Select":    Select,
	"Append":    Append,
	"Prepend":   Prepend,
	"Join":      Join,
	"Reduce":    Reduce,
	"Transduce": Transduce,
	"Reduced":   Reduced,
	"Into":      Into,
	"Concat":    Concat,
	"ConcatMap": ConcatMap,
	"Chain":     Chain,
	"Zip":       Zip,
	"All":       All,
	"Any":       Any,
	"Pipe":      Pipe,
	"Compose":   Compose,
	"First":     First,
	"Rest":      Rest,
	"Last":      Last,
	"Partition": Partition,
}

func init() {
	for k, v := range listOps {
		fmt.Printf("%v = %#v\n", k, v)
	}
}
