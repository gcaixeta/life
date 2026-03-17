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
