package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type IEventRepository interface {
	SaveErrEvent(ctx context.Context, errEvent []byte) error
}

type EventRepository struct {
	pgPool *pgxpool.Pool
}

func NewEventRepository(pgPool *pgxpool.Pool) *EventRepository {
	return &EventRepository{pgPool}
}

func (er *EventRepository) SaveErrEvent(ctx context.Context, event []byte) error {

	eventStr := string(event)
	queryString := fmt.Sprintf(`
		insert into checkbox.err_events (doc) values ('%s')
	`, eventStr)

	fmt.Println(queryString)

	rows, err := er.pgPool.Query(ctx, queryString)
	if err != nil {
		log.Println("Ошибака сохранения в сheckbox.err_events")
		return err
	}
	defer rows.Close()

	return nil
}
