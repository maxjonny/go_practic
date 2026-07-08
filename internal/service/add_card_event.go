package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"main/internal/models"
	pkg "main/pkg/models"
	"strconv"
	"time"
)

func (s *Service) AddCardEvent(event models.UserEvent) (bool, error) {

	isFounded, err := s.Event.CheckDouble(context.Background(), event.CheckDate, event.EquipmentModel)
	if err != nil || !isFounded {
		return false, err
	}

	currentDate := time.Now().Format("2006-01-02T15:04:05.000Z07:00")
	fileName := ""

	if len(event.ImgByte) != 0 {

		var sizeKbytes int

		fileName = strconv.FormatInt(time.Now().UnixMilli(), 10) + "_checkbox.jpeg"
		sizeKbytes, err = s.SaveJpeg(fileName, event.ImgByte)
		if err != nil {
			log.Printf("Ошибка сохранения фото. %s\n", err)
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

			recordId, record, err := s.Common.AddModel("main.files", fileLog)
			if err != nil {
				return false, err
			}

			historyLog := pkg.History{
				Source: "main.files",
				Doc: pkg.HistoryDoc{
					SourceId:   recordId,
					Service:    "checkbox",
					EventType:  "create",
					CreateDate: currentDate,
					CurrentModel: pkg.DbEntrie{
						Id:  recordId,
						Doc: json.RawMessage(record),
					},
				},
			}

			if _, err := s.Common.AddHistoryPg(historyLog); err != nil {
				return false, err
			}
			event.ImgByte = nil
		}

		event.Img.Name = fileName
		event.Img.Path += fileName
	}

	fmt.Printf("%+v\n", event)

	return true, nil

}
