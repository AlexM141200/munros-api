package main

import (
	"context"

	"github.com/AlexM141200/munros-api/src/api"
)

func main() {

	ctx := context.Background()
	server := api.NewAPIServer(":8080")

	server.Run(ctx)

}
