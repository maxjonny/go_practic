package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type IEventRepository interface {
	SaveErrEvent(ctx context.Context, errEvent []byte) error
	CheckDouble(ctx context.Context, checkDate string, equipmentModel string) (bool, error)
}

type EventRepository struct {
	pgPool *pgxpool.Pool
}

func NewEventRepository(pgPool *pgxpool.Pool) *EventRepository {
	return &EventRepository{pgPool}
}

func (er *EventRepository) SaveErrEvent(ctx context.Context, event []byte) error {

	eventStr := string(event)
	queryString := fmt.Sprintf(`insert into checkbox.err_events (doc) values ('%s')
		`, eventStr)

	rows, err := er.pgPool.Query(ctx, queryString)
	if err != nil {
		log.Println("Ошибака сохранения в сheckbox.err_events")
		return err
	}
	defer rows.Close()

	return nil
}

func (er *EventRepository) CheckDouble(ctx context.Context, checkDate string, equipmentModel string) (bool, error) {

	queryString := fmt.Sprintf(`SELECT count(*) as cnt FROM checkbox.human_events
                                    WHERE doc->>'checkDate' = '%s' and doc->>'equipmentModel' = '%s'`,
		checkDate, equipmentModel)

	var recordsCount int
	err := er.pgPool.QueryRow(ctx, queryString).Scan(&recordsCount)
	if err != nil {
		log.Println("Ошибака чтения из сheckbox.human_events")
	}

	if recordsCount > 0 {
		return false, err
	}
	return true, err
}
