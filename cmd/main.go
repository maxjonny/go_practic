package main

import (
	"context"
	"log"
	db "main/internal/database"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	httpApp := api.CreateServer(storage)
	httpApp.Run()

	<-ctx.Done()
	log.Println("Остановка сервера, закрытие соеднений")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	httpApp.Stop(shutdownCtx)
	storage.CloseConnection()

	log.Println("Сервер остановлен")
}
