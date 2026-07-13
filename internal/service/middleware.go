package service

import (
	"context"
	"time"
)

func (s *Service) UpdateBoxStatus(ctx context.Context, device string) error {

	currentDate := time.Now().Format("2006-01-02T15:04:05.000Z07:00")

	if err := s.Device.UpdateConnection(ctx, device, currentDate); err != nil {
		return err
	}

	return nil
}
