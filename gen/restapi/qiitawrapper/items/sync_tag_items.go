// Code generated by go-swagger; DO NOT EDIT.

package items

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// SyncTagItemsHandlerFunc turns a function with the right signature into a sync tag items handler
type SyncTagItemsHandlerFunc func(SyncTagItemsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SyncTagItemsHandlerFunc) Handle(params SyncTagItemsParams) middleware.Responder {
	return fn(params)
}

// SyncTagItemsHandler interface for that can handle valid sync tag items params
type SyncTagItemsHandler interface {
	Handle(SyncTagItemsParams) middleware.Responder
}

// NewSyncTagItems creates a new http.Handler for the sync tag items operation
func NewSyncTagItems(ctx *middleware.Context, handler SyncTagItemsHandler) *SyncTagItems {
	return &SyncTagItems{Context: ctx, Handler: handler}
}

/*SyncTagItems swagger:route GET /items/sync/{tag} items syncTagItems

同期的にタグの記事を取得。

指定日付に投稿された記事を一覧で取得。一度に取得できる記事の数は50固定。

*/
type SyncTagItems struct {
	Context *middleware.Context
	Handler SyncTagItemsHandler
}

func (o *SyncTagItems) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewSyncTagItemsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
