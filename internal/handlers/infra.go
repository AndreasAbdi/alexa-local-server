package handler

import (
	"log"
	"net/http"

	"github.com/AndreasAbdi/alexa-local-server/internal/infrared"
)

//HandleInfra sends a request to power the tv.
func HandleInfra(service *infrared.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("Handling an infra request")
		service.SwitchTvPower()
	}
}
