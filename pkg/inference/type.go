package inference

type Type interface {
	ID() int
	Match(Type) bool
	Unify(Type) (*Substitution, error)
	Apply(*Substitution) (Type, error)
}
