package main

import (
	"log"
	db "main/internal/database"
	"net/http"

	rep "main/internal/repository"
	api "main/internal/transport/http"

	"github.com/joho/godotenv"
)

func main() {

	//загрузка конфига
	//создание бд
	//создание интерфейсов
	//создание роутера

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	storage := db.InitStorage()
	storageInterface := rep.InitRepositoryInterface(storage)
	handlers := api.InitHandlers(storageInterface)
	router := api.InitRouter(handlers)

	port := "8080"
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
