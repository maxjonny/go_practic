package service

import (
	"context"
	"log"
	m "main/internal/models"
)

func (s *Service) GetUserCount(ctx context.Context, device string) (int, error) {

	var users []m.UserCard
	var err error

	nodeIds, err := s.Device.GetActiveNode(ctx, device)
	if err != nil {
		log.Println(err)
	}

	if len(nodeIds) > 0 {
		users, err = s.User.GetUsersByNodes(ctx, nodeIds)
		if err != nil {
			log.Println(err)
		}
	}

	if len(users) > 0 {
		s.User.DropCache(ctx, device)
		s.User.CreateCache(ctx, device, users)
	}

	return len(users), nil
}
