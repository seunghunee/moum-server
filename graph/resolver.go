package graph

//go:generate go run github.com/99designs/gqlgen

import (
	"github.com/seunghunee/moum-server/accessor"
)

// Resolver is the root graph resolver type.
type Resolver struct {
	Accessor accessor.Accessor
}
