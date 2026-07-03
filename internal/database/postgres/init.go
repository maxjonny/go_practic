package postgres

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	Pool *pgxpool.Pool
}

func (p *Postgres) connect(ctx context.Context) {
	ConnString := os.Getenv("DATABASE_URL")

	if ConnString == "" {
		log.Fatalf("DATABASE_URL is empty or not set")
	}

	dbpool, err := pgxpool.New(ctx, ConnString)
	if err != nil {
		log.Fatalf("Connection error %s", err)
	}

	pingCtx, pingCancel := context.WithTimeout(ctx, 3*time.Second)
	defer pingCancel()

	if err := dbpool.Ping(pingCtx); err != nil {
		dbpool.Close() // не забываем закрыть пул при ошибке
	}

	p.Pool = dbpool

	log.Printf("connected to %s", ConnString)
}

func (p *Postgres) CloseConnecion() {
	p.Pool.Close()
	log.Println("psql. Соединение закрыто")
}

func NewConnectPostgres(ctx context.Context) Postgres {
	db := Postgres{}
	db.connect(ctx)
	return db
}
