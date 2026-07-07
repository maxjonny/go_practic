package service

import (
	"context"
	mainRep "main/internal/repository"
)

type Service struct {
	rep mainRep.RepositoryInterface
}

func InitService(rep mainRep.RepositoryInterface) *Service {
	return &Service{rep: rep}
}

func (s *Service) SaveErrEvent(ctx context.Context, errEvent []byte) error {

	if err := s.rep.Event.SaveErrEvent(ctx, errEvent); err != nil {
		return err
	}
	return nil
}
