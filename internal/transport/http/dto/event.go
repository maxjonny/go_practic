package dto

import (
	"encoding/base64"
	"log"
	"main/internal/models"
)

type EventDtoIn struct {
	GID               string `json:"gID"`
	Name              string `json:"name"`
	DeptName          string `json:"deptName"`
	IPAdress          string `json:"iPAdress"`
	CheckDate         string `json:"checkDate"`
	CheckResult       string `json:"checkResult"`
	Authentication    string `json:"authentication"`
	EquipmentModel    string `json:"equipmentModel"`
	AlcoholStrength   string `json:"alcoholStrength"`
	CheckSerialNumber string `json:"checkSerialNumber"`
	FaceFeature       string `json:"faceFeature,omitempty"`
}

func (dto EventDtoIn) IsValid() bool {
	return dto.GID != "" &&
		dto.Name != "" &&
		dto.CheckDate != "" &&
		dto.CheckResult != "" &&
		dto.AlcoholStrength != "" &&
		dto.EquipmentModel != "" &&
		dto.Authentication != ""
}

func (dto EventDtoIn) ToServiceModel() models.UserEvent {

	var data []byte
	var err error

	if dto.FaceFeature != "" {
		data, err = base64.StdEncoding.DecodeString(dto.FaceFeature)
		if err != nil {
			log.Printf("Ошибка декодирования фото: %s", dto.GID)
		}
	}

	return models.UserEvent{
		GID:               dto.GID,
		ImgByte:           data,
		Name:              dto.Name,
		DeptName:          dto.DeptName,
		IPAdress:          dto.IPAdress,
		CheckDate:         dto.CheckDate,
		EventType:         "enter",
		CheckSerialNumber: dto.CheckSerialNumber,
		CheckResult:       dto.CheckResult,
	}
}
