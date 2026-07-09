package dto

import (
	"main/internal/models"
	"time"
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

func (dto *EventDtoIn) IsValid() bool {

	eventTime, err := time.ParseInLocation("2006/01/02 15:04:05", dto.CheckDate, time.Local)
	if err != nil {
		return false
	}
	dto.CheckDate = eventTime.Format("2006-01-02T15:04:05.000Z07:00")

	return dto.GID != "" &&
		dto.Name != "" &&
		dto.CheckResult != "" &&
		dto.AlcoholStrength != "" &&
		dto.EquipmentModel != "" &&
		dto.Authentication != ""
}

func (dto EventDtoIn) ToServiceModel() models.UserEvent {
	return models.UserEvent{
		GID:               dto.GID,
		Name:              dto.Name,
		DeptName:          dto.DeptName,
		IPAdress:          dto.IPAdress,
		CheckDate:         dto.CheckDate,
		FaceFeature:       dto.FaceFeature,
		EventType:         "enter",
		CheckSerialNumber: dto.CheckSerialNumber,
		CheckResult:       dto.CheckResult,
		EquipmentModel:    dto.EquipmentModel,
		Authentication:    dto.Authentication,
		AlcoholStrength:   dto.AlcoholStrength,
		Img: models.HumanImg{
			Path: "files/",
		},
	}
}
