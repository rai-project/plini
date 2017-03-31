package inference

type Substitution struct {
	subs map[int]Type
}

var (
	EmptySubstitution = &Substitution{subs: map[int]Type{}}
)
