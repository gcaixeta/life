package event

import (
	"context"
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Insert(ctx context.Context, event *Event) error {
	query := `
		INSERT INTO events (title, description, started_at, ended_at, category, project, origin)
		VALUES (?, ?, ?, ?, ?, ?, ?);
	`

	result, err := r.db.ExecContext(ctx, query,
		event.Title,
		event.Description,
		event.StartedAt,
		event.EndedAt,
		event.Category,
		event.Project,
		event.Origin,
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

func (r *Repository) FindAll(ctx context.Context) ([]Event, error) {
	query := `
		SELECT * FROM events;
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	var events []Event

	for rows.Next() {
		var event Event
		if err := rows.Scan(&event.ID, &event.Title, &event.Description,
			&event.Category, &event.StartedAt, &event.EndedAt,
			&event.Project, &event.Origin); err != nil {
			return events, err
		}

		events = append(events, event)
	}
	if err = rows.Err(); err != nil {
		return events, err
	}
	return events, nil
}
