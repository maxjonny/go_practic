package redis

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Client *redis.Client
}

func (r *Redis) connect(ctx context.Context) {
	ConnString := os.Getenv("REDIS_URL")

	opt, err := redis.ParseURL(ConnString)
	if err != nil {
		log.Fatal("REDIS_URL is empty or not set")
	}

	r.Client = redis.NewClient(opt)

	pingCtx, pingCancel := context.WithTimeout(ctx, 3*time.Second)
	defer pingCancel()

	if err := r.Client.Ping(pingCtx).Err(); err != nil {
		log.Fatalf("Connection error %s", err)
	}

	log.Printf("connected to %s", ConnString)
}
func (p *Redis) CloseConnecion() {
	p.Client.Close()
	log.Println("redis. Соединение закрыто")
}

func NewConnectRedis(ctx context.Context) Redis {
	db := Redis{}
	db.connect(ctx)
	return db
}
