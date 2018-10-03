package handler

import (
	"net/http"
	"time"

	"github.com/AndreasAbdi/alexa-local-server/server/cast"
)

//HandleQuit to terminate based on the quit endpoint
func HandleQuit(service *cast.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		device, err := service.GetDevice()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		device.QuitApplication(time.Second * 10)
	}
}
