// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetHostRequirementsHandlerFunc turns a function with the right signature into a get host requirements handler
type GetHostRequirementsHandlerFunc func(GetHostRequirementsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetHostRequirementsHandlerFunc) Handle(params GetHostRequirementsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetHostRequirementsHandler interface for that can handle valid get host requirements params
type GetHostRequirementsHandler interface {
	Handle(GetHostRequirementsParams, interface{}) middleware.Responder
}

// NewGetHostRequirements creates a new http.Handler for the get host requirements operation
func NewGetHostRequirements(ctx *middleware.Context, handler GetHostRequirementsHandler) *GetHostRequirements {
	return &GetHostRequirements{Context: ctx, Handler: handler}
}

/*GetHostRequirements swagger:route GET /host_requirements installer getHostRequirements

Get minimum host requirements.

*/
type GetHostRequirements struct {
	Context *middleware.Context
	Handler GetHostRequirementsHandler
}

func (o *GetHostRequirements) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetHostRequirementsParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
