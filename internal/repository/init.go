package repository

import (
	db "main/internal/database"
)

type RepositoryInterface struct {
	User   IUserRepository
	Device IDeviceRepository
	Event  IEventRepository
}

func InitRepositoryInterface(s *db.Storage) RepositoryInterface {
	return RepositoryInterface{
		User:   NewUserRepository(s.Pg.Pool, s.Redis.Client),
		Device: NewDeviceRepository(s.Pg.Pool),
		Event:  NewEventRepository(s.Pg.Pool),
	}
}
