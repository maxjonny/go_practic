package service

import (
	"context"
	"log"
	m "main/internal/models"
)

func (s *Service) GetUserData(ctx context.Context, device string, index string) (*m.UserCard, error) {

	user, err := s.User.GetUser(ctx, device, index)
	if err != nil {
		log.Println(err)
	}

	return user, nil

}
