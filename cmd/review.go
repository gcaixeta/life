package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/gcaixeta/life/internal/db"
	"github.com/gcaixeta/life/internal/event"
	"github.com/gcaixeta/life/internal/review"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(reviewCmd)
}

var reviewCmd = &cobra.Command{
	Use:   "review",
	Short: "Reviews the events",
	Long:  "Reviews the events happened between a certain period. By default, the period is yesterdays. Accepts flags to change it.",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		conn, err := db.OpenConn()
		if err != nil {
			panic(err)
		}

		var events []event.Event

		repo := event.NewRepository(conn)
		events, err = repo.FindByDay(ctx, time.Now().AddDate(0, 0, -1))
		if err != nil {
			panic(err)
		}

		if len(events) == 0 {
			fmt.Println("No events for yesterdays. Do better!")
		}

		review.ReviewEvents(events)
		err = repo.InsertAll(ctx, events)
		if err != nil {
			fmt.Println("Error trying to save the events after review")
		}

		fmt.Println("Are events updated?")
		for i := range events {
			score := *events[i].Score
			ratedAt := *events[i].RatedAt
			note := *events[i].Note
			fmt.Printf("Event: %v. Score: %d. Rated At: %v. Note: %v",
				events[i].Title,
				score,
				ratedAt,
				note,
			)
		}
	},
}
