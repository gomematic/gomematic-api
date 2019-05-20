// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DelteTeamFromUserHandlerFunc turns a function with the right signature into a delte team from user handler
type DelteTeamFromUserHandlerFunc func(DelteTeamFromUserParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DelteTeamFromUserHandlerFunc) Handle(params DelteTeamFromUserParams) middleware.Responder {
	return fn(params)
}

// DelteTeamFromUserHandler interface for that can handle valid delte team from user params
type DelteTeamFromUserHandler interface {
	Handle(DelteTeamFromUserParams) middleware.Responder
}

// NewDelteTeamFromUser creates a new http.Handler for the delte team from user operation
func NewDelteTeamFromUser(ctx *middleware.Context, handler DelteTeamFromUserHandler) *DelteTeamFromUser {
	return &DelteTeamFromUser{Context: ctx, Handler: handler}
}

/*DelteTeamFromUser swagger:route DELETE /teams/{teamID}/users team delteTeamFromUser

Remove a user from team

*/
type DelteTeamFromUser struct {
	Context *middleware.Context
	Handler DelteTeamFromUserHandler
}

func (o *DelteTeamFromUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDelteTeamFromUserParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}