package service

import (
	"context"
	"encoding/base64"
	"log"
	"main/internal/models"
	pkg "main/pkg/models"
	"strconv"
	"time"
)

func (s *Service) UpsertUser(ctx context.Context, user models.UserCard) (bool, error) {

	currentDate := time.Now().Format("2006-01-02T15:04:05.000Z07:00")

	userCardId, userCard, err := s.User.GetByGid(ctx, user.GID)
	if err != nil {
		return false, err
	}

	if user.FaceFeature != "" {

		var sizeKbytes int
		var fileName string

		data, err := base64.StdEncoding.DecodeString(user.FaceFeature)
		if err != nil {
			log.Printf("Ошибка декодирования фото: %s", user.GID)
		}

		fileName = strconv.FormatInt(time.Now().UnixMilli(), 10) + "_checkbox.jpeg"
		sizeKbytes, err = s.SaveJpeg(fileName, data)
		if err != nil {
			log.Printf("Ошибка сохранения фото. %s\n", err)
			return false, err
		} else {

			fileLog := pkg.FileInfo{
				ServiceSourceId: 1,
				Type:            "checkbox",
				Name:            fileName,
				MimeType:        "image/jpeg",
				SizeKbyte:       sizeKbytes,
				Path:            "files/" + fileName,
				CreateDate:      currentDate,
			}

			_, log, err := s.Common.AddModel("main.files", fileLog)
			if err != nil {
				return false, err
			} else {
				log.Doc.CreateDate = currentDate
			}

			if _, err := s.Common.AddHistoryPg(log); err != nil {
				return false, err
			}
			user.FaceFeature = ""
		}

		user.Img.Name = fileName
		user.Img.Path += fileName
	}

	var log pkg.History
	if userCardId == 0 {
		_, log, err = s.Common.AddModel("checkbox.human_card", user)
		if err != nil {
			return false, err
		} else {
			log.Doc.CreateDate = currentDate
		}
	} else {

		if user.Img.Name != "" {
			userCard.Img = user.Img
		}

		_, log, err = s.Common.UpdatePg("checkbox.human_card", userCardId, userCard)
		if err != nil {
			return false, err
		} else {
			log.Doc.CreateDate = currentDate
		}
	}

	_, err = s.Common.AddHistoryPg(log)
	if err != nil {
		return false, err
	}

	return true, nil
}
