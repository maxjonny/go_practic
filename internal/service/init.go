package service

import (
	mainRep "main/internal/repository"
)

type Service struct {
	rep mainRep.RepositoryInterface
}

func InitService(rep mainRep.RepositoryInterface) *Service {
	return &Service{rep: rep}
}
