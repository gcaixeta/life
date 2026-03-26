package event

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	flag "github.com/spf13/pflag"
)

type Origin string

const (
	Manual   Origin = "manual"
	Hook     Origin = "hook"
	Pomodoro Origin = "pomodoro"
)

type Event struct {
	ID          int64
	Title       string
	Description string
	StartedAt   time.Time
	EndedAt     *time.Time
	Category    string
	Project     string
	Origin      Origin
	Score       *int
	Note        *string
	RatedAt     *time.Time
}

func NewEventFromFlags(flags *flag.FlagSet) (*Event, error) {
	title, err := flags.GetString("title")
	if err != nil {
		return nil, err
	}

	description, err := flags.GetString("description")
	if err != nil {
		return nil, err
	}

	return &Event{
		Title:       title,
		Description: description,
		StartedAt:   time.Now(),
		Origin:      Manual,
	}, nil
}

func NewEventFromPrompt() (*Event, error) {
	reader := bufio.NewReader(os.Stdin)

	var title, description string

	fmt.Print("Title: ")
	title, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	title = strings.TrimSpace(title)

	fmt.Print("Description: ")
	description, err = reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	description = strings.TrimSpace(description)

	return &Event{
		Title:       title,
		Description: description,
		StartedAt:   time.Now(),
		Origin:      Manual,
	}, nil

}

func (e *Event) String() string {
	ended := ""
	if e.EndedAt != nil {
		ended = " → " + e.EndedAt.Format("15:04")
	}

	return fmt.Sprintf("Event #%d created\n%s · %s%s · %s",
		e.ID,
		e.Title,
		e.StartedAt.Format("02/01/2006 15:04"),
		ended,
		e.Origin,
	)
}
