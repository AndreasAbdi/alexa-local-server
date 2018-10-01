package app

import "net/http"

func (s *Server) handleYoutube() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		device, err := getDevice()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		device.PlayYoutubeVideo("F1B9Fk_SgI0")
	}
}
