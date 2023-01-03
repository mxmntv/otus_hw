package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgconn"
	"github.com/mxmntv/otus_hw/hw12_calendar/internal/model"
	"github.com/mxmntv/otus_hw/hw12_calendar/pkg/logger"
	"github.com/mxmntv/otus_hw/hw12_calendar/pkg/postgres"
)

type EventRepository struct {
	*postgres.Postgres
	*logger.Logger
}

func NewEventRepoPg(pg *postgres.Postgres, l *logger.Logger) *EventRepository {
	return &EventRepository{pg, l}
}

func (repo EventRepository) CreateEvent(ctx context.Context, event *model.Event) error {
	query := `
		INSERT INTO event
		(title, date, duration, description, owner, notification)
		VALUES
		($1, $2, $3, $4, $5, $6)
		RETURNING id
	`
	if err := repo.Pool.QueryRow(ctx, query, event).Scan(&event.ID); err != nil {
		var pgErr *pgconn.PgError
		if errors.Is(err, pgErr) {
			// todo implement pg error handling (create error & loggin it)
			repo.Logger.Error(err.Error())
			return err
		}
		return err
	}
	return nil
}

func (repo EventRepository) UpdateEvent(ctx context.Context, event *model.Event) error {
	query := `
		UPDATE event SET 
		title=COALESCE($2, title),
		date=COALESCE($3, date),
		duration=COALESCE($4, duration),
		description=COALESCE($5, description), 
		owner=COALESCE($6, owner), 
		notification=COALESCE($7, notification)
		WHERE id=$1
	`

	tag, err := repo.Pool.Exec(ctx, query, &event.ID, &event.Title, &event.DateTime,
		&event.Duration, &event.Description, &event.OwnerID, &event.NotificationTime)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.Is(err, pgErr) {
			// todo implement pg error handling (create error & loggin it)
			repo.Logger.Error(err.Error())
			return err
		}
		return err
	}
	fmt.Println(tag.RowsAffected()) // todo log this
	return nil
}

func (repo EventRepository) EventList(ctx context.Context) ([]model.Event, error) {
	query := `
		SELECT id, title, description, duration, owner, notification FROM event
	`
	rows, err := repo.Pool.Query(ctx, query)
	if err != nil {
		// todo implement pg error handling (create error & loggin it)
		repo.Logger.Error(err.Error())
		return nil, err
	}

	events := make([]model.Event, 0)

	for rows.Next() {
		var event model.Event

		err := rows.Scan(&event.ID, &event.Title, &event.Description,
			&event.Duration, &event.OwnerID, &event.NotificationTime)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func (repo EventRepository) DeleteEvent(ctx context.Context, id string) error {
	query := `
		DELETE event WHERE id=$1
	`
	tag, err := repo.Pool.Exec(ctx, query, id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.Is(err, pgErr) {
			// todo implement pg error handling (create error & loggin it)
			repo.Logger.Error(err.Error())
			return err
		}
		return err
	}
	fmt.Println(tag.RowsAffected()) // todo log this
	return nil
}
