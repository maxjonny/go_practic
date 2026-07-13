package models

type FileInfo struct {
	CreateDate      string `json:"create_date"`
	AccountId       int    `json:"account_id"`
	ServiceSourceId int    `json:"service_source_id"`
	Type            string `json:"type"`
	Name            string `json:"name"`
	MimeType        string `json:"mime_type"`
	SizeKbyte       int    `json:"size_kbyte"`
	Path            string `json:"path"`
}
