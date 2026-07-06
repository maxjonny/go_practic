package models

type HumanImg struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type UserCard struct {
	GID           string   `json:"gID"`
	Img           HumanImg `json:"img"`
	GZBH          string   `json:"gZBH"`
	Name          string   `json:"name"`
	DeptName      string   `json:"deptName"`
	HumanID       int      `json:"human_id"`
	FromDevice    bool     `json:"fromDevice"`
	FingerFeature string   `json:"fingerFeature"`
	FaceFeature   string   `json:"faceFeature,omitempty"`
}
