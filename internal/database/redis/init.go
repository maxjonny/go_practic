package redis

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Client *redis.Client
}

func (r *Redis) GetClient() {
	ConnString := os.Getenv("REDIS_URL")

	opt, err := redis.ParseURL(ConnString)
	if err != nil {
		log.Fatal("REDIS_URL is empty or not set")
	}

	r.Client = redis.NewClient(opt)

	if err := r.Client.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Connection error %s", err)
	}

	log.Printf("connected to %s", ConnString)
}
func (p *Redis) CloseConnecion() {
	p.Client.Close()
	log.Println("redis. Соединение закрыто")
}

func NewConnectRedis() Redis {
	db := Redis{}
	db.GetClient()
	return db
}
