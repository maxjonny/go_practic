package dto

import (
	"main/internal/models"
)

type UserDto struct {
	GID           string `json:"gID"`
	GZBH          string `json:"gZBH"`
	Name          string `json:"name"`
	DeptName      string `json:"deptName"`
	FingerFeature string `json:"fingerFeature"`
	FaceFeature   string `json:"faceFeature"`
}

func (dto *UserDto) IsValid() bool {
	return dto.GID != "" &&
		dto.GZBH != "" &&
		dto.Name != ""
}

func (dto *UserDto) ToServiceModel() models.UserCard {
	return models.UserCard{
		GID:           dto.GID,
		GZBH:          dto.GZBH,
		Name:          dto.Name,
		DeptName:      dto.DeptName,
		FromDevice:    true,
		FingerFeature: dto.FingerFeature,
		FaceFeature:   dto.FaceFeature,
		Img: models.HumanImg{
			Path: "files/",
		},
	}
}

func (dto *UserDto) FromServiceModel(user *models.UserCard) {
	dto.GID = user.GID
	dto.GZBH = user.GZBH
	dto.Name = user.Name
	dto.DeptName = user.DeptName
	dto.FingerFeature = user.FingerFeature
	dto.FaceFeature = user.FaceFeature
}
