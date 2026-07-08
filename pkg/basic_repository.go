package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"main/pkg/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type IBasicRepository interface {
	AddHistoryPg(Record models.History) (int, error)
	AddModel(table string, Record any) (id int, doc string, err error)
}

type BasicRepository struct {
	pgPool *pgxpool.Pool
}

func InitBasicRepository(pgPool *pgxpool.Pool) *BasicRepository {
	return &BasicRepository{pgPool}
}

func (br *BasicRepository) AddHistoryPg(Record models.History) (int, error) {

	var ReturningId int

	docByte, err := json.Marshal(Record.Doc)
	if err != nil {
		return ReturningId, err
	}

	queryString := fmt.Sprintf(`INSERT INTO main.history (doc, source) VALUES ('%s', '%s') RETURNING id`,
		string(docByte), Record.Source)

	err = br.pgPool.QueryRow(context.Background(), queryString).Scan(&ReturningId)
	if err != nil {
		log.Printf("Ошибка сохранения истории, %s", err)
	}

	return ReturningId, err
}

func (br *BasicRepository) AddModel(table string, Record any) (id int, doc string, err error) {

	var docByte []byte
	docByte, err = json.Marshal(Record)
	if err != nil {
		return
	}

	queryString := fmt.Sprintf(`INSERT INTO %s (doc) VALUES ('%s') RETURNING id, doc`,
		table, string(docByte))

	err = br.pgPool.QueryRow(context.Background(), queryString).Scan(&id, &doc)
	if err != nil {
		log.Printf("Ошибка сохранения истории, %s", err)
	}

	return
}
