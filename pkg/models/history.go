package models

import "time"

type HistoryDoc struct {
	SourceId        int      `json:"source_id"`
	Service         string   `json:"service"`
	EventType       string   `json:"event_type"`
	CreateDate      string   `json:"create_date"`
	CreateAccountId string   `json:"create_account_id"`
	CurrentModel    DbEntrie `json:"current_model"`
}

type History struct {
	Source string     `json:"source"`
	Doc    HistoryDoc `json:"doc"`
}

func CreateHistory(source string) *History {
	layout := "2006-01-02T15:04:05.000Z"
	currentDate := time.Now().UTC().Format(layout)

	return &History{
		Source: source,
		Doc: HistoryDoc{
			CreateDate: currentDate,
		},
	}
}
