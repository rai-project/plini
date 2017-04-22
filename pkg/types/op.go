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
)
