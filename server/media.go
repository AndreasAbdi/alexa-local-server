package app

import "net/http"

func (s *Server) handleMedia() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		device, err := getDevice()
		if err != nil {
			return
		}
		device.PlayMedia(
			"http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/images/BigBuckBunny.jpg",
			"image/jpeg")
	}
}
