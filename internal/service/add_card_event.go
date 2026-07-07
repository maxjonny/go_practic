package service

import (
	"context"
	"fmt"
	"main/internal/models"
)

func (s *Service) AddCardEvent(ctx context.Context, event models.UserEvent) error {
	event.ImgByte = nil
	fmt.Printf("%+v", event)
	return nil
}
