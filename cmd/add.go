package cmd

import (
	"context"
	"fmt"

	"github.com/gcaixeta/life/internal/db"
	"github.com/gcaixeta/life/internal/event"
	"github.com/spf13/cobra"
)

func init() {
	var title, description string
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&title)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new event",
	Long:  "Life basic entity, the command prompts user for the new event details",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		conn, err := db.OpenConn()
		if err != nil {
			panic(err)
		}

		newEvent, err := event.NewEventFromPrompt()
		if err != nil {
			panic(err)
		}

		repo := event.NewRepository(conn)
		err = repo.Insert(ctx, newEvent)
		if err != nil {
			panic(err)
		}

		fmt.Printf("\n%v\n", newEvent)

	},
}
