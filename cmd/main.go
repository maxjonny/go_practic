package main

import (
	"log"
	db "main/internal/database"
	"net/http"

	t "main/internal/transport"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	db.InitStorage()
	router := t.InitRouter()

	port := "8080"
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
