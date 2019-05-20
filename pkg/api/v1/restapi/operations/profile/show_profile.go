// Code generated by go-swagger; DO NOT EDIT.

package profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ShowProfileHandlerFunc turns a function with the right signature into a show profile handler
type ShowProfileHandlerFunc func(ShowProfileParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ShowProfileHandlerFunc) Handle(params ShowProfileParams) middleware.Responder {
	return fn(params)
}

// ShowProfileHandler interface for that can handle valid show profile params
type ShowProfileHandler interface {
	Handle(ShowProfileParams) middleware.Responder
}

// NewShowProfile creates a new http.Handler for the show profile operation
func NewShowProfile(ctx *middleware.Context, handler ShowProfileHandler) *ShowProfile {
	return &ShowProfile{Context: ctx, Handler: handler}
}

/*ShowProfile swagger:route GET /profile/self profile showProfile

Retrieve an unlimited auth token

*/
type ShowProfile struct {
	Context *middleware.Context
	Handler ShowProfileHandler
}

func (o *ShowProfile) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewShowProfileParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}