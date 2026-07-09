package models

type HistoryDoc struct {
	SourceId        int      `json:"source_id"`
	EventType       string   `json:"event_type"`
	CreateDate      string   `json:"create_date"`
	CurrentModel    DbEntrie `json:"current_model"`
	CreateAccountId int      `json:"create_account_id"`
}

type History struct {
	Source string     `json:"source"`
	Doc    HistoryDoc `json:"doc"`
}
