package main

import (
	"context"
	"fmt"

	"github.com/gcaixeta/life/internal/db"
	"github.com/gcaixeta/life/internal/event"
)

func main() {
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
}
