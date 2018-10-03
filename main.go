package main

import (
	"github.com/AndreasAbdi/alexa-local-server/server"
)

func main() {
	server := app.NewServer(":8000", "amzn1.ask.skill.45c29c7b-8da0-4c1f-92c2-7fb87258f9cb")
	server.Init()
}
