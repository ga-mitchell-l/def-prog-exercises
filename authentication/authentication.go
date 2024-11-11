package authentication

import (
	"context"
	"slices"
)

// For simplicity you may use strings today.
type Privilege string

type privilegeKey struct{}

func Grant(c context.Context, ps ...Privilege) context.Context {
	if c.Value(privilegeKey{}) != nil {
		panic("Grant called multiple times")
	}
	return context.WithValue(c, privilegeKey{}, ps)
}

type checkedKey struct{}

func Check(c context.Context, p ...Privilege) (_ context.Context, ok bool) {
	granted, ok := c.Value(privilegeKey{}).([]Privilege)
	if !ok {
		return c, false
	}
	for _, ps := range p {
		if !slices.Contains(granted, ps) {
			return c, false
		}
	}
	return context.WithValue(c, checkedKey{}, struct{}{}), true
}

func Must(c context.Context) (ok bool) {
	return c.Value(checkedKey{}) != nil
}
