package main

import (
	"context"
	"log"
	db "main/internal/database"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	rep "main/internal/repository"
	api "main/internal/transport/http"

	"github.com/joho/godotenv"
)

func main() {

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	storage := db.InitStorage(ctx)
	defer storage.CloseConnection()

	storageInterface := rep.InitRepositoryInterface(storage)
	handlers := api.InitHandlers(storageInterface)
	router := api.InitRouter(handlers)

	port := "8080"

	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	} else {
		log.Printf("Сервер запущен на порту %s\n", port)
	}

	<-ctx.Done()
	log.Println("Получен сигнал завершения, останавливаем сервер...")
}
