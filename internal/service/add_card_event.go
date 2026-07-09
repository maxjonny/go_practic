package service

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"main/internal/models"
	"main/internal/repository"
	pkg "main/pkg/models"
	"maps"
	"slices"
	"strconv"
	"time"
)

func (s *Service) AddCardEvent(event models.UserEvent) (bool, error) {

	ctx := context.Background()

	isFounded, err := s.Event.CheckDouble(ctx, event.CheckDate, event.EquipmentModel)
	if err != nil || !isFounded {
		return false, err
	}

	currentDate := time.Now().Format("2006-01-02T15:04:05.000Z07:00")
	fileName := ""

	if event.FaceFeature != "" {

		var sizeKbytes int

		data, err := base64.StdEncoding.DecodeString(event.FaceFeature)
		if err != nil {
			log.Printf("Ошибка декодирования фото: %s", event.GID)
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
			event.FaceFeature = ""
		}

		event.Img.Name = fileName
		event.Img.Path += fileName
	}

	deviceRel, err := s.Device.GetDeviceRelations(ctx, event.EquipmentModel)
	if err != nil {
		return false, err
	}

	if len(deviceRel) == 0 {
		return false, fmt.Errorf("Модель %s не привязяна к проекту\n", event.EquipmentModel)
	}

	mapProject := make(map[int]int)
	for _, rel := range deviceRel {
		mapProject[rel.NodeId] = rel.ProjectId
	}

	userRel, err := s.User.GetUserRelations(ctx, event.GID)
	if err != nil {
		return false, err
	}

	var eventProjects []int
	for _, node := range userRel.NodeIds {
		projectId, has := mapProject[node]
		if has {
			eventProjects = append(eventProjects, projectId)
		}
	}

	if len(eventProjects) == 0 {
		eventProjects = slices.Collect(maps.Values(mapProject))
	}

	if userRel.UserCardId == 0 {
		//Создание карточки
		//userRel.UserCardId от новой карты
	}

	event.HumanCardId = userRel.UserCardId

	for _, project_id := range eventProjects {

		event.ProjectId = project_id
		_, log, err := s.Common.AddModel("checkbox.human_events", event)
		if err != nil {
			return false, err
		} else {
			log.Doc.CreateDate = currentDate
		}

		if _, err := s.Common.AddHistoryPg(log); err != nil {
			return false, err
		}

		workerId, err := s.User.GetWorkerId(ctx, event.HumanCardId, event.ProjectId)
		if err != nil {
			return false, err
		}

		worker := repository.WorkerStatus{
			EnterDate:       currentDate,
			CheckResult:     event.EventType,
			AlcoholStrength: event.AlcoholStrength,
			ProjectId:       event.ProjectId,
			HumanCardId:     event.HumanCardId,
		}

		if event.AlcoholStrength >= "30" {
			worker.Status = "drunk"
		} else {
			worker.Status = "work"
		}

		if workerId == 0 {
			_, log, err := s.Common.AddModel("checkbox.workers", worker)
			if err != nil {
				return false, err
			} else {
				log.Doc.CreateDate = currentDate
			}

			if _, err := s.Common.AddHistoryPg(log); err != nil {
				return false, err
			}
		} else {
			_, log, err := s.Common.UpdatePg("checkbox.workers", workerId, worker)
			if err != nil {
				return false, err
			} else {
				log.Doc.CreateDate = currentDate
			}

			if _, err := s.Common.AddHistoryPg(log); err != nil {
				return false, err
			}
		}
	}

	return true, nil

}
