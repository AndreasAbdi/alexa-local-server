package handler

import (
	"net/http"

	"github.com/AndreasAbdi/alexa-local-server/internal/cast"
)

//HandleYoutube for http calls to the youtube endpoint
func HandleYoutube(service *cast.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		device, err := service.GetDevice()
		if err != nil {
			return
		}
		device.PlayYoutubeVideo("F1B9Fk_SgI0")

	}
}
