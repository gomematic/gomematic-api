// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ListUserTeamsHandlerFunc turns a function with the right signature into a list user teams handler
type ListUserTeamsHandlerFunc func(ListUserTeamsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListUserTeamsHandlerFunc) Handle(params ListUserTeamsParams) middleware.Responder {
	return fn(params)
}

// ListUserTeamsHandler interface for that can handle valid list user teams params
type ListUserTeamsHandler interface {
	Handle(ListUserTeamsParams) middleware.Responder
}

// NewListUserTeams creates a new http.Handler for the list user teams operation
func NewListUserTeams(ctx *middleware.Context, handler ListUserTeamsHandler) *ListUserTeams {
	return &ListUserTeams{Context: ctx, Handler: handler}
}

/*ListUserTeams swagger:route GET /users/{userID}/teams user listUserTeams

Fetch all teams assigned to user

*/
type ListUserTeams struct {
	Context *middleware.Context
	Handler ListUserTeamsHandler
}

func (o *ListUserTeams) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListUserTeamsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}