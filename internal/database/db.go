package database

import (
	pg "main/internal/database/postgres"
	redis "main/internal/database/redis"
)

type Storage struct {
	Pg    pg.Postgres
	Redis redis.Redis
}

func (s *Storage) CloseConnection() {
	s.Pg.CloseConnecion()
	s.Redis.CloseConnecion()
}

func InitStorage() *Storage {
	db := Storage{}
	db.Pg = pg.NewConnectPostgres()
	db.Redis = redis.NewConnectRedis()
	return &db
}
