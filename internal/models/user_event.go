package models

type UserEvent struct {
	GID               string   `json:"gID"`
	Img               HumanImg `json:"img"`
	Name              string   `json:"name"`
	DeptName          string   `json:"deptName"`
	IPAdress          string   `json:"iPAdress"`
	CheckDate         string   `json:"checkDate"`
	EventType         string   `json:"eventType"`
	ProjectId         int      `json:"project_id"`
	CheckResult       string   `json:"checkResult"`
	FaceFeature       string   `json:"faceFeature"`
	HumanCardId       int      `json:"human_card_id"`
	Authentication    string   `json:"authentication"`
	EquipmentModel    string   `json:"equipmentModel"`
	AlcoholStrength   string   `json:"alcoholStrength"`
	CheckSerialNumber string   `json:"checkSerialNumber"`
}
