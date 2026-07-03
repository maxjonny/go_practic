package database

import (
	"context"
	pg "main/internal/database/postgres"
	redis "main/internal/database/redis"
)

type Storage struct {
	Pg    pg.Postgres
	Redis redis.Redis
}

func (s *Storage) CloseConnection() {
	defer s.Pg.CloseConnecion()
	defer s.Redis.CloseConnecion()
}

func InitStorage(ctx context.Context) *Storage {
	db := Storage{}
	db.Pg = pg.NewConnectPostgres(ctx)
	db.Redis = redis.NewConnectRedis(ctx)
	return &db
}
