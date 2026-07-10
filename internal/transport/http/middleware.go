package http

import "net/http"

func UpdateConnect(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println("Лог: пришел запрос на", r.URL.Path)

		// Передаем управление следующему обработчику
		next.ServeHTTP(w, r)
	})
}
