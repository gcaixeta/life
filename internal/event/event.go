package event

import "time"

type Category string
type Project string
type Origin string

const (
	Manual    Origin = "manual"
	Automatic Origin = "automatic"
)

type Event struct {
	ID          string
	Title       string
	Description string
	StartedAt   time.Time
	EndedAt     time.Time
	Category    Category
	Project     Project
	Origin      Origin
}
