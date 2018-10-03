package youtube

import (
	"context"
	"fmt"
	"net/http"

	"github.com/AndreasAbdi/alexa-local-server/server/alexa"
	"github.com/AndreasAbdi/alexa-local-server/server/cast"
	"github.com/mikeflynn/go-alexa/skillserver"
	"google.golang.org/api/googleapi/transport"
	youtube "google.golang.org/api/youtube/v3"
)

const maxResults = 25

//SearchVideo looks to youtube for videos with a defined query value and return id and title.
func SearchVideo(ctx context.Context, googleKey string, query string) (id string, title string, err error) {

	client := &http.Client{
		Transport: &transport.APIKey{Key: googleKey},
	}

	service, err := youtube.New(client)
	if err != nil {
		return "", "", err
	}

	call := service.Search.List("id,snippet").Q(query).Type("video").MaxResults(maxResults)
	response, err := call.Do()
	if err != nil {
		return "", "", err
	}

	for _, item := range response.Items {
		return item.Id.VideoId, item.Snippet.Title, nil
	}
	return "", "", nil
}

//HandleSearch returns a function handler for alexa requests
func HandleSearch(googleKey string, service *cast.Service) alexa.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *skillserver.EchoRequest) {

		query, err := r.GetSlotValue("searchQuery")
		if err != nil {
			http.Error(w, "no searchquery slot in unmarshalled alexa request", 500)
			return
		}
		id, title, err := SearchVideo(ctx, googleKey, query)
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
