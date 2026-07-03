package http

import (
	"net/http"
)

type Router struct {
	mux *http.ServeMux
}

func InitRouter(handlers *Handler) *Router {
	r := Router{http.NewServeMux()}

	r.mux.HandleFunc("/checkbox/Z5/{device}/actionapi/User/GetUserCount", handlers.GetUserCount)
	r.mux.HandleFunc("/checkbox/Z5/{device}/actionapi/User/GetUserData", handlers.GetUserData)

	r.mux.HandleFunc("/checkbox/Z5/{device}/", handlers.CheckConnect)
	return &r
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}
