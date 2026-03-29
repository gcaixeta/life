package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gcaixeta/life/internal/db"
	"github.com/gcaixeta/life/internal/event"
	"github.com/gcaixeta/life/internal/review"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(reviewCmd)
	reviewCmd.Flags().StringP("date", "d", "", "Specifies a date of the events to review")
}

var reviewCmd = &cobra.Command{
	Use:   "review",
	Short: "Reviews the events",
	Long:  "Reviews the events happened on a given day. Defaults to yesterday. Use --date to specify another day.",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		conn, err := db.OpenConn()
		if err != nil {
			panic(err)
		}

		repo := event.NewRepository(conn)

		var date time.Time
		if cmd.Flags().Changed("date") {
			dateString, err := cmd.Flags().GetString("date")
			if err != nil {
				panic(err)
			}
			date, err = parseDateFromString(dateString)
			if err != nil {
				fmt.Println("Error parsing the informed date.")
				os.Exit(1)
			}
		} else {
			date = time.Now().AddDate(0, 0, -1)
		}

		events, err := repo.FindByDay(ctx, date)
		if err != nil {
			panic(err)
		}

		if len(events) == 0 {
			fmt.Printf("No events for %s. Do better!\n", date.Format(time.DateOnly))
			return
		}

		review.ReviewEvents(events)

		err = repo.UpdateAll(ctx, events)
		if err != nil {
			fmt.Println("Error trying to save the events after review.")
			os.Exit(1)
		}
	},
}

func parseDateFromString(dateString string) (time.Time, error) {
	if dateString == "today" {
		return time.Now(), nil
	}
	date, err := time.Parse(time.DateOnly, dateString)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}
