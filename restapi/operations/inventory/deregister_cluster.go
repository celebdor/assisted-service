// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeregisterClusterHandlerFunc turns a function with the right signature into a deregister cluster handler
type DeregisterClusterHandlerFunc func(DeregisterClusterParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeregisterClusterHandlerFunc) Handle(params DeregisterClusterParams) middleware.Responder {
	return fn(params)
}

// DeregisterClusterHandler interface for that can handle valid deregister cluster params
type DeregisterClusterHandler interface {
	Handle(DeregisterClusterParams) middleware.Responder
}

// NewDeregisterCluster creates a new http.Handler for the deregister cluster operation
func NewDeregisterCluster(ctx *middleware.Context, handler DeregisterClusterHandler) *DeregisterCluster {
	return &DeregisterCluster{Context: ctx, Handler: handler}
}

/*DeregisterCluster swagger:route DELETE /clusters/{cluster_id} inventory deregisterCluster

Deregister OpenShift bare metal cluster

*/
type DeregisterCluster struct {
	Context *middleware.Context
	Handler DeregisterClusterHandler
}

func (o *DeregisterCluster) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeregisterClusterParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}