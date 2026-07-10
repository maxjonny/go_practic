package http

import (
	"net/http"
)

type Router struct {
	mux *http.ServeMux
}

func InitRouter(handlers *Handler) *Router {
	r := Router{http.NewServeMux()}

	r.mux.Handle("/checkbox/Z5/{device}/actionapi/User/GetUserCount", r.with(handlers.GetUserCount, UpdateConnect))
	r.mux.Handle("/checkbox/Z5/{device}/actionapi/User/GetUserData", r.with(handlers.GetUserData))
	r.mux.Handle("/checkbox/Z5/{device}/actionapi/User/UploadAlcohol", r.with(handlers.AddCardEvent, UpdateConnect))
	r.mux.Handle("/checkbox/Z5/{device}/actionapi/User/UploadUser", r.with(handlers.UploadUser))

	r.mux.Handle("/checkbox/Z5/{device}/", r.with(handlers.CheckConnect))
	return &r
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}

func (r *Router) with(h http.HandlerFunc, middlewares ...func(http.Handler) http.Handler) http.Handler {
	var finalHandler http.Handler = http.HandlerFunc(h)
	for i := len(middlewares) - 1; i >= 0; i-- {
		finalHandler = middlewares[i](finalHandler)
	}
	return finalHandler
}

//mux.HandleFunc("GET /users", myHandler)
