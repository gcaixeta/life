package cmd

import (
	"context"
	"fmt"

	"github.com/gcaixeta/life/internal/db"
	"github.com/gcaixeta/life/internal/event"
	"github.com/spf13/cobra"

	flag "github.com/spf13/pflag"
)

var title, description string

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&title, "title", "t", "", "Title of the event")
	addCmd.Flags().StringVarP(&description, "description", "d", "", "Description for the event")
	addCmd.MarkFlagsRequiredTogether("title", "description")
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

		var newEvent *event.Event

		if !flagsPresent(cmd.Flags()) {
			newEvent, err = event.NewEventFromPrompt()
			if err != nil {
				panic(err)
			}
		} else {
			newEvent, err = event.NewEventFromFlags(cmd.Flags())
			if err != nil {
				panic(err)
			}
		}

		repo := event.NewRepository(conn)
		err = repo.Insert(ctx, newEvent)
		if err != nil {
			panic(err)
		}

		fmt.Printf("\n%v\n", newEvent)
	},
}

func flagsPresent(flags *flag.FlagSet) bool {
	return flags.Changed("title") && flags.Changed("description")
}
