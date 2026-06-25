package repository

import (
	db "main/internal/database"
)

type RepositoryInterface struct {
	User   IUserRepository
	Device IDeviceRepository
}

func InitRepositoryInterface(s *db.Storage) RepositoryInterface {
	return RepositoryInterface{
		User:   NewUserRepository(s.Pg.Pool, s.Redis.Client),
		Device: NewDeviceRepository(s.Pg.Pool),
	}
}
