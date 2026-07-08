package service

import (
	"context"
	"log"
	mainRep "main/internal/repository"
	"os"
)

type Service struct {
	mainRep.RepositoryInterface
}

func InitService(rep mainRep.RepositoryInterface) *Service {
	return &Service{rep}
}

func (s *Service) SaveErrEvent(ctx context.Context, errEvent []byte) error {

	if err := s.Event.SaveErrEvent(ctx, errEvent); err != nil {
		return err
	}
	return nil
}

func (s *Service) SaveJpeg(imgName string, imgBytes []byte) (int, error) {

	fileSizeInBytes := len(imgBytes)
	fileSizeInKB := int(float64(fileSizeInBytes) / 1024)

	filePath := "./upload_files/" + imgName
	err := os.WriteFile(filePath, imgBytes, 0644)
	if err != nil {
		log.Printf("Ошибка при сохранении файла на диск: %v", err)
	}
	return fileSizeInKB, err
}
