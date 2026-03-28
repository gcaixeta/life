package review

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gcaixeta/life/internal/display"
	"github.com/gcaixeta/life/internal/event"
)

func ReviewEvents(events []event.Event) {
	reader := bufio.NewReader(os.Stdin)

	for i := range events {
		display.PrintEvent(events[i])

		score, err := askRating(reader)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		events[i].Score = &score

		if score < 3 || score == 5 {
			var message string
			if score == 5 {
				message = "great"
			} else {
				message = "bad"
			}

			note, err := askNote(reader, message)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				events[i].Note = &note
			}
		}

		now := time.Now()
		events[i].RatedAt = &now

		fmt.Printf("Saved rating %d for event '%s'\n", score, events[i].Title)
	}
}

func askRating(r *bufio.Reader) (int, error) {
	for {
		fmt.Print("Give a score between 1 and 5 to this event: ")
		input, err := r.ReadString('\n')
		if err != nil {
			return 0, err
		}

		score, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			fmt.Println("Please enter a valid number. Try again!")
			continue
		}

		if score < 1 || score > 5 {
			fmt.Println("Score should be from 1 to 5. Try again!")
			continue
		}

		return score, nil
	}
}

func askNote(r *bufio.Reader, message string) (string, error) {
	fmt.Printf("Tell me what went %s about it: ", message)

	input, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(input), nil
}

