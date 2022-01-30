package graph

import (
	"github.com/cansirin/gezdimgordum/graphql/internal/backend"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Backend backend.Backender
}
