package main

import (
	"github.com/AlexM141200/munros-api/src/api"
)

func main() {

	server := api.NewAPIServer(":8080")

	server.Run()

}
