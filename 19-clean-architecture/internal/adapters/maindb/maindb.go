package maindb

import (
	"context"
	"time"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"

	"github.com/otusteam/go/cleancalendar/internal/domain/entities"
)

// implements domain.interfaces.EventStorage
type PgEventStorage struct {
	db *sqlx.DB
}

func NewPgEventStorage(dsn string) (*PgEventStorage, error) {
	db, err := sqlx.Open("pgx", dsn) // *sql.DB
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &PgEventStorage{db: db}, nil
}

func (pges *PgEventStorage) SaveEvent(ctx context.Context, event *entities.Event) error {
	query := `
		INSERT INTO events(id, owner, title, text, start_time, end_time)
		VALUES (:id, :owner, :title, :text, :start_time, :end_time)
	`
	_, err := pges.db.NamedExecContext(ctx, query, map[string]interface{}{
		"id":         event.ID.String(),
		"owner":      event.Owner,
		"title":      event.Title,
		"text":       event.Text,
		"start_time": event.StartTime,
		"end_time":   event.EndTime,
	})
	return err
}

func (pges *PgEventStorage) GetEventByID(ctx context.Context, id string) (*entities.Event, error) {
	var res entities.Event

	query := `
		SELECT
			*
		FROM events
		WHERE id = $1
	`

	err := pges.db.GetContext(ctx, &res, query, id)

	return &res, err
}

func (pges *PgEventStorage) GetEventsByOwnerStartDate(ctx context.Context, owner string, startTime time.Time) *[]entities.Event {
	var res []entities.Event

	query := `
		SELECT
			*
		FROM events
		WHERE owner = $1 AND start_time = $2
	`

	if err := pges.db.SelectContext(ctx, &res, query, owner, startTime); err != nil {
		return nil
	}

	return &res
}
