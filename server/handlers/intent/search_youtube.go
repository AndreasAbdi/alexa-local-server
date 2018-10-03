package intent

import (
	"context"
	"fmt"
	"net/http"

	"github.com/AndreasAbdi/alexa-local-server/server/alexa"
	"github.com/AndreasAbdi/alexa-local-server/server/cast"
	"github.com/AndreasAbdi/alexa-local-server/server/youtube"
	"github.com/mikeflynn/go-alexa/skillserver"
)

//HandleSearch returns a function handler for alexa requests
func HandleSearch(googleKey string, service *cast.Service) alexa.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *skillserver.EchoRequest) {

		query, err := r.GetSlotValue("searchQuery")
		if err != nil {
			http.Error(w, "no searchquery slot in unmarshalled alexa request", 500)
			return
		}
		id, title, err := youtube.SearchVideo(ctx, googleKey, query)
		if err != nil {
			http.Error(w, "Failed to perform search", 500)
			return
		}
		if err != nil {
			fmt.Println("error:", err)
		}
		go func() {
			playOnCast(id, service)
		}()
		writeHandleSearchVideoResponse(title, w)
		return
	}
}

func playOnCast(videoID string, service *cast.Service) {
	device, err := service.GetDevice()
	if err != nil {
		return
	}
	device.PlayYoutubeVideo(videoID)
}

func writeHandleSearchVideoResponse(title string, w http.ResponseWriter) {
	resp := skillserver.NewEchoResponse()
	resp.OutputSpeech("Playing " + title)
	alexa.WriteResponse(w, resp)
}
