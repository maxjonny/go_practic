package postgres

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	Pool *pgxpool.Pool
}

func (p *Postgres) ConnectPool() {
	ConnString := os.Getenv("DATABASE_URL")

	if ConnString == "" {
		log.Fatalf("DATABASE_URL is empty or not set")
	}

	dbpool, err := pgxpool.New(context.Background(), ConnString)
	if err != nil {
		log.Fatalf("Connection error %s", err)
	}

	p.Pool = dbpool

	log.Printf("connected to %s", ConnString)

}

func (p *Postgres) CloseConnecion() {
	p.Pool.Close()
	log.Println("psql. Соединение закрыто")
}

func NewConnectPostgres() Postgres {
	db := Postgres{}
	db.ConnectPool()
	return db
}
