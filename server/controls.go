package app

import (
	"net/http"
	"time"
)

func (s *Server) handleQuit() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		device, err := getDevice()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		device.QuitApplication(time.Second * 10)
	}
}
