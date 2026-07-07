package models

type UserCard struct {
	GID           string
	Img           HumanImg
	GZBH          string
	Name          string
	DeptName      string
	HumanID       int
	FromDevice    bool
	FingerFeature string
	FaceFeature   string
}
