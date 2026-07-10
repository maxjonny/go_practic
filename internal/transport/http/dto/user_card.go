package dto

import (
	"main/internal/models"
)

type UserDtoIn struct {
	GID           string `json:"gID"`
	GZBH          string `json:"gZBH"`
	Name          string `json:"name"`
	DeptName      string `json:"deptName"`
	FingerFeature string `json:"fingerFeature"`
	FaceFeature   string `json:"faceFeature"`
}

func (dto *UserDtoIn) IsValid() bool {
	return dto.GID != "" &&
		dto.GZBH != "" &&
		dto.Name != ""
}

func (dto *UserDtoIn) ToServiceModel() models.UserCard {
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
