package main

import (
	"fmt"

	"github.com/AndreasAbdi/alexa-local-server/server"
)

func main() {
	server := app.NewServer()
	server.Init()
	fmt.Print("Local Server deployed!")
}
