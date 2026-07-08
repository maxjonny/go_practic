package models

import "encoding/json"

type DbEntrie struct {
	Id  int             `json:"id"`
	Doc json.RawMessage `json:"doc"`
}
