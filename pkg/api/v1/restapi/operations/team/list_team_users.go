// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ListTeamUsersHandlerFunc turns a function with the right signature into a list team users handler
type ListTeamUsersHandlerFunc func(ListTeamUsersParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListTeamUsersHandlerFunc) Handle(params ListTeamUsersParams) middleware.Responder {
	return fn(params)
}

// ListTeamUsersHandler interface for that can handle valid list team users params
type ListTeamUsersHandler interface {
	Handle(ListTeamUsersParams) middleware.Responder
}

// NewListTeamUsers creates a new http.Handler for the list team users operation
func NewListTeamUsers(ctx *middleware.Context, handler ListTeamUsersHandler) *ListTeamUsers {
	return &ListTeamUsers{Context: ctx, Handler: handler}
}

/*ListTeamUsers swagger:route GET /teams/{teamID}/users team listTeamUsers

Fetch all users assigned to team

*/
type ListTeamUsers struct {
	Context *middleware.Context
	Handler ListTeamUsersHandler
}

func (o *ListTeamUsers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListTeamUsersParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}