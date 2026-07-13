package http

import (
	"github.com/go-chi/chi/v5"
)

func InitRouter(handlers *Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Route("/checkbox/Z5/", func(r chi.Router) {

		r.Use(UpdateConnect)

		r.Get("{device}/actionapi/User/GetUserCount", handlers.GetUserCount)
		r.Get("{device}/actionapi/User/GetUserData", handlers.GetUserData)
		r.Post("{device}/actionapi/User/UploadAlcohol", handlers.AddCardEvent)
		r.Post("{device}/actionapi/User/UploadUser", handlers.UploadUser)

		r.Head("/", handlers.CheckConnect)
	})

	return r
}
