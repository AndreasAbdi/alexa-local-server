package main

import (
	"github.com/AndreasAbdi/alexa-local-server/server"
)

func main() {
	server := app.NewServer(":8000", "test")
	server.Init()
}
