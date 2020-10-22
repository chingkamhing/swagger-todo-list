// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// AddOneHandlerFunc turns a function with the right signature into a add one handler
type AddOneHandlerFunc func(AddOneParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddOneHandlerFunc) Handle(params AddOneParams) middleware.Responder {
	return fn(params)
}

// AddOneHandler interface for that can handle valid add one params
type AddOneHandler interface {
	Handle(AddOneParams) middleware.Responder
}

// NewAddOne creates a new http.Handler for the add one operation
func NewAddOne(ctx *middleware.Context, handler AddOneHandler) *AddOne {
	return &AddOne{Context: ctx, Handler: handler}
}

/*AddOne swagger:route POST /api/todo/ todos addOne

AddOne add one API

*/
type AddOne struct {
	Context *middleware.Context
	Handler AddOneHandler
}

func (o *AddOne) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddOneParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
