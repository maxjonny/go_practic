package service

import (
	"context"
	"fmt"
	m "main/internal/models"
)

func (s *Service) GetUserCount(ctx context.Context, device string) (int, error) {

	var users []m.UserCard
	var err error

	nodeIds, err := s.Device.GetActiveNode(ctx, device)
	if err != nil {
		return 0, fmt.Errorf("GetActiveNode: %w", err)
	}

	if len(nodeIds) > 0 {
		users, err = s.User.GetUsersByNodes(ctx, nodeIds)
		if err != nil {
			return 0, fmt.Errorf("GetUsersByNodes: %w", err)
		}
	}

	if len(users) > 0 {
		if err := s.User.DropCache(ctx, device); err != nil {
			return 0, fmt.Errorf("DropCache: %w", err)
		}
		if err := s.User.CreateCache(ctx, device, users); err != nil {
			return 0, fmt.Errorf("CreateCache: %w", err)
		}
	}

	return len(users), nil
}
