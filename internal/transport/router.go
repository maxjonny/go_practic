package transport

import (
	"github.com/go-chi/chi/v5"
)

func InitRouter(handlers *Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Route("/checkbox/Z5", func(r chi.Router) {

		//r.Use(handlers.Test)

		r.Route("/{device}/actionapi/User", func(r chi.Router) {

			r.Use(handlers.UpdateConnect)

			r.Get("/GetUserCount", handlers.GetUserCount)
			r.Post("/UploadAlcohol", handlers.AddCardEvent)
			r.Post("/UploadUser", handlers.UploadUser)

		})

		r.Get("/{device}/actionapi/User/GetUserData", handlers.GetUserData)
		r.Head("/{device}", handlers.CheckConnect)
	})

	return r
}
