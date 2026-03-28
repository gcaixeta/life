package event

import (
	"context"
	"database/sql"
	"time"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Insert(ctx context.Context, event *Event) error {
	query := `
		INSERT INTO events (title, description, started_at, ended_at, category, project, origin, rating_score, rating_note, rated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
	`

	result, err := r.db.ExecContext(ctx, query,
		event.Title,
		event.Description,
		event.StartedAt,
		event.EndedAt,
		event.Category,
		event.Project,
		event.Origin,
		event.Score,
		event.Note,
		event.RatedAt,
	)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	event.ID = id
	return nil
}

func (r *Repository) FindByDay(ctx context.Context, day time.Time) ([]Event, error) {
	startOfDay := time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, day.Location())
	endOfDay := time.Date(day.Year(), day.Month(), day.Day(), 23, 59, 59, 0, day.Location())

	query := `
		SELECT id, title, description, started_at, ended_at, category, project, origin, rating_score, rating_note, rated_at
		FROM events
		WHERE started_at BETWEEN ? AND ?;
	`
	rows, err := r.db.QueryContext(ctx, query, startOfDay, endOfDay)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		if err := rows.Scan(
			&event.ID,
			&event.Title,
			&event.Description,
			&event.StartedAt,
			&event.EndedAt,
			&event.Category,
			&event.Project,
			&event.Origin,
			&event.Score,
			&event.Note,
			&event.RatedAt,
		); err != nil {
			return events, err
		}
		events = append(events, event)
	}
	if err = rows.Err(); err != nil {
		return events, err
	}
	return events, nil
}

func (r *Repository) InsertAll(ctx context.Context, events []Event) error {
	for i := range events {
		err := r.Insert(ctx, &events[i])
		if err != nil {
			return err
		}
	}
	return nil
}
