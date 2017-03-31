package inference

import (
	"github.com/pkg/errors"
)

type Constructor struct {
	id    int
	name  string
	arity int
}

func NewConstructor(name string, arity int) (Type, error) {
	return &Constructor{
		id:    <-genSym,
		name:  name,
		arity: arity,
	}, nil
}

func (c *Constructor) ID() int {
	return c.id
}

func (c *Constructor) Unify(t Type) (*Substitution, error) {
	switch t := t.(type) {
	case *Constructor:
		if c.name == t.name {
			return EmptySubstitution, nil
		}
		return nil,
			errors.Errorf("unable to unify constructor %v and constructor %v", c.name, t.name)
	case *Variable:
		return t.Unify(c)
	}
	return nil,
		errors.Errorf("unexpected rhs %v when unifying constructor", t)
}

func (c *Constructor) Match(t Type) bool {
	_, err := c.Unify(t)
	return err == nil
}

func (c *Constructor) Apply(sub *Substitution) (Type, error) {
	return nil, nil
}
