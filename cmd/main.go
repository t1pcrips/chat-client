package main

import (
	"context"
	"github.com/t1pcrips/chat-client/internal/app"
	"log"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx, "local.env")
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = a.Run(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
