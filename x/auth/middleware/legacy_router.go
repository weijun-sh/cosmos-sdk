package middleware

import (
	"fmt"

	sdk "github.com/weijun-sh/cosmos-sdk/types"
)

type LegacyRouter struct {
	routes map[string]sdk.Handler
}

var _ sdk.Router = NewLegacyRouter()

// NewRouter returns a reference to a new router.
func NewLegacyRouter() *LegacyRouter {
	return &LegacyRouter{
		routes: make(map[string]sdk.Handler),
	}
}

// AddRoute adds a route path to the router with a given handler. The route must
// be alphanumeric.
func (rtr *LegacyRouter) AddRoute(route sdk.Route) sdk.Router {
	if !sdk.IsAlphaNumeric(route.Path()) {
		panic("route expressions can only contain alphanumeric characters")
	}
	if rtr.routes[route.Path()] != nil {
		panic(fmt.Sprintf("route %s has already been initialized", route.Path()))
	}

	rtr.routes[route.Path()] = route.Handler()
	return rtr
}

// Route returns a handler for a given route path.
//
// TODO: Handle expressive matches.
func (rtr *LegacyRouter) Route(_ sdk.Context, path string) sdk.Handler {
	return rtr.routes[path]
}
