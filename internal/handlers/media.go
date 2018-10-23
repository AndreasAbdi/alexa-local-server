package handler

import (
	"net/http"

	"github.com/AndreasAbdi/alexa-local-server/internal/cast"
)

//HandleMedia for http calls to the media endpoint. Specifically for testing chromecast endpoint.
func HandleMedia(service *cast.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		device, err := service.GetDevice()
		if err != nil {
			return
		}
		device.PlayMedia(
			"http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/images/BigBuckBunny.jpg",
			"image/jpeg")
	}
}
