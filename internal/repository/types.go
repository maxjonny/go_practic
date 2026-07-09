package repository

type UserRelations struct {
	UserCardId int
	NodeIds    []int
}

type DeviceRelations struct {
	NodeId    int
	ProjectId int
}

type WorkerStatus struct {
	Status          string `json:"status"`
	ExitDate        string `json:"exitDate"`
	EnterDate       string `json:"enterDate"`
	ProjectId       int    `json:"project_id"`
	CheckResult     string `json:"checkResult"`
	HumanCardId     int    `json:"human_card_id"`
	AlcoholStrength string `json:"alcoholStrength"`
}
