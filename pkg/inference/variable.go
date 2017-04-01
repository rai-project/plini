package inference

import "github.com/pkg/errors"

type Variable struct {
	id   int
	name string
}

func NewVariable(name string) (Type, error) {
	return &Variable{
		id:   <-genSym,
		name: name,
	}, nil
}

func (v *Variable) ID() int {
	return v.id
}

func (v *Variable) Unify(t Type) (*Substitution, error) {
	switch t := t.(type) {
	case *Constructor:
		return &Substitution{
			subs: map[int]Type{
				v.ID(): t,
			},
		}, nil
	case *Variable:
		return nil, errors.Errorf("not implemented")
	}
	return nil,
		errors.Errorf("unexpected rhs %v when unifying constructor", t)
}

func (v *Variable) Match(t Type) bool {
	_, err := v.Unify(t)
	return err == nil
}

func (v *Variable) Apply(sub *Substitution) (Type, error) {
	return nil, nil
}
