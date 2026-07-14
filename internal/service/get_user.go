package service

import (
	"context"
	"encoding/base64"
	"log"
	m "main/internal/models"
)

func (s *Service) GetUserData(ctx context.Context, device string, index string) (*m.UserCard, error) {

	user, err := s.User.GetUserCache(ctx, device, index)
	if err != nil {
		return nil, err
	}

	imgBytes, err := s.LoadJpeg(user.Img.Name)
	if err != nil {
		log.Printf("Фото не найдено! gId:%s", user.GID)
	}

	if len(imgBytes) != 0 {
		encodedImg := base64.StdEncoding.EncodeToString(imgBytes)
		user.FaceFeature = encodedImg
	}

	return user, nil

}
