package transport

import (
	"context"
	"log"
	"main/internal/database"
	"main/internal/repository"
	"net/http"
)

type server struct {
	server *http.Server
}

func CreateServer(storage *database.Storage) *server {

	storageInterface := repository.InitRepositoryInterface(storage)
	handlers := InitHandlers(storageInterface)
	router := InitRouter(handlers)

	return &server{
		server: &http.Server{
			Addr:    ":4010",
			Handler: router,
		},
	}
}

func (a *server) Run() {
	go func() {
		log.Printf("Http. Сервер запущен на порту %s", a.server.Addr)
		if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Ошибка запуска: %s", err)
		}
	}()
}

func (a *server) Stop(ctx context.Context) {
	if err := a.server.Shutdown(ctx); err != nil {
		log.Printf("Принудительное завершение: %v", err)
	}
}
