package display

import (
	"fmt"

	"github.com/gcaixeta/life/internal/event"
)

func PrintEvent(e event.Event) {
	fmt.Printf("────────────────────────────\n")
	fmt.Printf("ID:            %d\n", e.ID)
	fmt.Printf("Title:         %s\n", e.Title)
	fmt.Printf("Description:   %s\n", e.Description)
	fmt.Printf("Início:        %s\n", e.StartedAt.Format("15:04"))
}
