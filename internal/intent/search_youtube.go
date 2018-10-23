package intent

import (
	"context"
	"net/http"
	"strings"

	"github.com/AndreasAbdi/alexa-local-server/internal/alexa"
	"github.com/AndreasAbdi/alexa-local-server/internal/cast"
	"github.com/AndreasAbdi/alexa-local-server/internal/youtube"
	gocast "github.com/AndreasAbdi/gochromecast"
	"github.com/mikeflynn/go-alexa/skillserver"
)

const slotSearchQuery = "searchQuery"

//HandlePlayNext returns a function handler for alexa requests that adds a video to the youtube playlist.
func HandlePlayNext(googleKey string, service *cast.Service) alexa.HandlerFunc {
	return genericSearch(googleKey, service, func(id string, device *gocast.Device) {
		device.YoutubeController.PlayNext(id)
	})
}

//HandleAddToPlaylist returns a function handler for alexa requests that adds a video to the youtube playlist.
func HandleAddToPlaylist(googleKey string, service *cast.Service) alexa.HandlerFunc {
	return genericSearch(googleKey, service, func(id string, device *gocast.Device) {
		device.YoutubeController.AddToQueue(id)
	})
}

//HandleSearch returns a function handler for alexa requests
func HandleSearch(googleKey string, service *cast.Service) alexa.HandlerFunc {
	return genericSearch(googleKey, service, func(id string, device *gocast.Device) {
		device.PlayYoutubeVideo(id)
	})
}

func genericSearch(googleKey string, service *cast.Service, command func(string, *gocast.Device)) alexa.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *skillserver.EchoRequest) {
		query, err := r.GetSlotValue(slotSearchQuery)
		if err != nil {
			http.Error(w, "no searchquery slot in unmarshalled alexa request", 500)
			return
		}
		id, title, err := youtube.SearchVideo(ctx, googleKey, query)
		if err != nil {
			http.Error(w, "Failed to perform search", 500)
			return
		}
		go func() {
			device, err := service.GetDevice()
			if err != nil {
				return
			}
			command(id, device)
		}()
		alexa.WriteSpeech(w, "Playing "+shortenTitle(title))
		return
	}
}

func shortenTitle(title string) string {
	shortened := strings.Split(title, "(")
	shortened = strings.Split(shortened[0], "|")
	return shortened[0]
}
