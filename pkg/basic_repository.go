package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"main/pkg/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type IBasicRepository interface {
	AddHistoryPg(Record models.History) (int, error)
	AddModel(table string, Record any) (id int, historyLog models.History, err error)
	UpdatePg(table string, recordId int, Record any) (id int, historyLog models.History, err error)
}

type BasicRepository struct {
	pgPool *pgxpool.Pool
}

func InitBasicRepository(pgPool *pgxpool.Pool) *BasicRepository {
	return &BasicRepository{pgPool}
}

func (br *BasicRepository) AddHistoryPg(Record models.History) (id int, err error) {

	docByte, err := json.Marshal(Record.Doc)
	if err != nil {
		return
	}

	queryString := `INSERT INTO main.history (doc, source) VALUES ($1, $2) RETURNING id`

	err = br.pgPool.QueryRow(context.Background(), queryString, string(docByte), Record.Source).Scan(&id)

	return
}

func (br *BasicRepository) AddModel(table string, Record any) (id int, historyLog models.History, err error) {

	var docByte []byte
	docByte, err = json.Marshal(Record)
	if err != nil {
		return
	}

	queryString := fmt.Sprintf(`INSERT INTO %s (doc) VALUES ($1) RETURNING id`,
		table)

	err = br.pgPool.QueryRow(context.Background(), queryString, string(docByte)).Scan(&id)

	historyLog = models.History{
		Source: table,
		Doc: models.HistoryDoc{
			SourceId:   id,
			EventType:  "create",
			CreateDate: "",
			CurrentModel: models.DbEntrie{
				Id:  id,
				Doc: json.RawMessage(docByte),
			},
		},
	}

	return
}

func (br *BasicRepository) UpdatePg(table string, recordId int, record any) (id int, historyLog models.History, err error) {

	var docByte []byte
	docByte, err = json.Marshal(record)
	if err != nil {
		return
	}

	queryString := fmt.Sprintf(`UPDATE %s SET doc = $1 WHERE id = $2 RETURNING id`, table)

	err = br.pgPool.QueryRow(context.Background(), queryString, string(docByte), recordId).Scan(&id)

	historyLog = models.History{
		Source: table,
		Doc: models.HistoryDoc{
			SourceId:   id,
			EventType:  "update",
			CreateDate: "",
			CurrentModel: models.DbEntrie{
				Id:  id,
				Doc: json.RawMessage(docByte),
			},
		},
	}

	return
}
