package models

type UserEvent struct {
	GID               string
	ImgByte           []byte
	Img               HumanImg
	Name              string
	DeptName          string
	IPAdress          string
	CheckDate         string
	EventType         string
	Project_id        int
	CheckResult       string
	HumanCardId       int
	Authentication    string
	EquipmentModel    string
	AlcoholStrength   string
	CheckSerialNumber string
	FaceFeature       string
}
