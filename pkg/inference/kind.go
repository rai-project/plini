package inference

type Kind struct {
	arity int
}

var (
	KindStar = Kind{arity: 0}
)
