// Code generated by go-swagger; DO NOT EDIT.

package stored_alert

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetSeverityChartHandlerFunc turns a function with the right signature into a get severity chart handler
type GetSeverityChartHandlerFunc func(GetSeverityChartParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSeverityChartHandlerFunc) Handle(params GetSeverityChartParams) middleware.Responder {
	return fn(params)
}

// GetSeverityChartHandler interface for that can handle valid get severity chart params
type GetSeverityChartHandler interface {
	Handle(GetSeverityChartParams) middleware.Responder
}

// NewGetSeverityChart creates a new http.Handler for the get severity chart operation
func NewGetSeverityChart(ctx *middleware.Context, handler GetSeverityChartHandler) *GetSeverityChart {
	return &GetSeverityChart{Context: ctx, Handler: handler}
}

/*
	GetSeverityChart swagger:route GET /stored_alerts/severity_days stored_alert getSeverityChart

Get chart of stored alerts
*/
type GetSeverityChart struct {
	Context *middleware.Context
	Handler GetSeverityChartHandler
}

func (o *GetSeverityChart) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetSeverityChartParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
