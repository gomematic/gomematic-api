// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// VerifyAuthHandlerFunc turns a function with the right signature into a verify auth handler
type VerifyAuthHandlerFunc func(VerifyAuthParams) middleware.Responder

// Handle executing the request and returning a response
func (fn VerifyAuthHandlerFunc) Handle(params VerifyAuthParams) middleware.Responder {
	return fn(params)
}

// VerifyAuthHandler interface for that can handle valid verify auth params
type VerifyAuthHandler interface {
	Handle(VerifyAuthParams) middleware.Responder
}

// NewVerifyAuth creates a new http.Handler for the verify auth operation
func NewVerifyAuth(ctx *middleware.Context, handler VerifyAuthHandler) *VerifyAuth {
	return &VerifyAuth{Context: ctx, Handler: handler}
}

/*VerifyAuth swagger:route GET /auth/verify/{token} auth verifyAuth

Verify validity for an authentication token

*/
type VerifyAuth struct {
	Context *middleware.Context
	Handler VerifyAuthHandler
}

func (o *VerifyAuth) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewVerifyAuthParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}