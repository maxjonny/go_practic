package repository

import (
	"context"
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
	queryString := `insert into checkbox.err_events (doc) values ($1)`

	rows, err := er.pgPool.Query(ctx, queryString, eventStr)
	if err != nil {
		log.Println("Ошибака сохранения в сheckbox.err_events")
		return err
	}
	defer rows.Close()

	return nil
}

func (er *EventRepository) CheckDouble(ctx context.Context, checkDate string, equipmentModel string) (bool, error) {

	queryString := `SELECT count(*) as cnt FROM checkbox.human_events
                                    WHERE doc->>'checkDate' = $1 and doc->>'equipmentModel' = $2`

	var recordsCount int
	err := er.pgPool.QueryRow(ctx, queryString, checkDate, equipmentModel).Scan(&recordsCount)
	if err != nil {
		log.Println("Ошибака чтения из сheckbox.human_events")
	}

	if recordsCount > 0 {
		return false, err
	}
	return true, err
}
