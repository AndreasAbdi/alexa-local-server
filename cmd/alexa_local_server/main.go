/*
A server for controlling a variety of smart home devices.

Current build is for personal use, manipulate at will for your use case.
*/
package main

import (
	"github.com/AndreasAbdi/alexa-local-server/internal/server"
)

func main() {
	server := app.NewServer()
	server.Init()
}
