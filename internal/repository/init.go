package repository

import (
	"github.com/redis/go-redis/v9"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RepositoryInterface struct {
	User   IUserRepository
	Device IDeviceRepository
}

func InitRepositoryInterface(pgPool *pgxpool.Pool, rClient *redis.Client) RepositoryInterface {
	return RepositoryInterface{
		User:   NewUserRepository(pgPool, rClient),
		Device: NewDeviceRepository(pgPool),
	}
}
