package main

import (
	"github.com/AndreasAbdi/alexa-local-server/internal/server"
)

func main() {
	server := app.NewServer()
	server.Init()
}
