package cmd

import (
	"context"
	"fmt"

	"github.com/gcaixeta/life/internal/db"
	"github.com/gcaixeta/life/internal/event"
	"github.com/spf13/cobra"

	flag "github.com/spf13/pflag"
)

func init() {
	rootCmd.AddCommand()

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

		var events []Event

	},
}
