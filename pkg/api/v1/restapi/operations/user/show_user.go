// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ShowUserHandlerFunc turns a function with the right signature into a show user handler
type ShowUserHandlerFunc func(ShowUserParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ShowUserHandlerFunc) Handle(params ShowUserParams) middleware.Responder {
	return fn(params)
}

// ShowUserHandler interface for that can handle valid show user params
type ShowUserHandler interface {
	Handle(ShowUserParams) middleware.Responder
}

// NewShowUser creates a new http.Handler for the show user operation
func NewShowUser(ctx *middleware.Context, handler ShowUserHandler) *ShowUser {
	return &ShowUser{Context: ctx, Handler: handler}
}

/*ShowUser swagger:route GET /users/{userID} user showUser

Fetch a specific user

*/
type ShowUser struct {
	Context *middleware.Context
	Handler ShowUserHandler
}

func (o *ShowUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewShowUserParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}