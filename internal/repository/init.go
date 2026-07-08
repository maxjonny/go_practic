package repository

import (
	db "main/internal/database"
	common "main/pkg"
)

type RepositoryInterface struct {
	Common common.IBasicRepository

	User   IUserRepository
	Device IDeviceRepository
	Event  IEventRepository
}

func InitRepositoryInterface(s *db.Storage) RepositoryInterface {
	return RepositoryInterface{
		Common: common.InitBasicRepository(s.Pg.Pool),

		User:   NewUserRepository(s.Pg.Pool, s.Redis.Client),
		Device: NewDeviceRepository(s.Pg.Pool),
		Event:  NewEventRepository(s.Pg.Pool),
	}
}
