package gql

import (
	"main/internal/repository"
	"main/internal/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.

type Resolver struct {
	service *service.Service
}

func InitResolvers(db repository.RepositoryInterface) *Resolver {
	return &Resolver{service: service.InitService(db)}
}
