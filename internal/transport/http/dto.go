package http

import m "main/internal/models"



type UserCardDtoIn struct {
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

func (u *UserCardDtoIn) ToModel {
	return &{}
} 
