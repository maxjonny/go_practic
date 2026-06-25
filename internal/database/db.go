package database

import (
	pg "main/internal/database/postgres"
	redis "main/internal/database/redis"
	rep "main/internal/repository"
)

type db struct {
	Pg    pg.Postgres
	Redis redis.Redis
	Rep   rep.RepositoryInterface
}

var Storage db

func InitStorage() {
	Storage.Pg = pg.NewConnectPostgres()
	Storage.Redis = redis.NewConnectRedis()
	Storage.Rep = rep.InitRepositoryInterface(Storage.Pg.Pool, Storage.Redis.Client)
}
