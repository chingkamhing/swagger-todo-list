// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteOneHandlerFunc turns a function with the right signature into a delete one handler
type DeleteOneHandlerFunc func(DeleteOneParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteOneHandlerFunc) Handle(params DeleteOneParams) middleware.Responder {
	return fn(params)
}

// DeleteOneHandler interface for that can handle valid delete one params
type DeleteOneHandler interface {
	Handle(DeleteOneParams) middleware.Responder
}

// NewDeleteOne creates a new http.Handler for the delete one operation
func NewDeleteOne(ctx *middleware.Context, handler DeleteOneHandler) *DeleteOne {
	return &DeleteOne{Context: ctx, Handler: handler}
}

/*DeleteOne swagger:route DELETE /api/todo/{id} todos deleteOne

DeleteOne delete one API

*/
type DeleteOne struct {
	Context *middleware.Context
	Handler DeleteOneHandler
}

func (o *DeleteOne) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteOneParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
