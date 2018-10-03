package youtube

import (
	"context"
	"log"
	"net/http"

	"github.com/AndreasAbdi/alexa-local-server/server/alexa"
	"github.com/AndreasAbdi/alexa-local-server/server/cast"
	"github.com/mikeflynn/go-alexa/skillserver"
	"google.golang.org/api/googleapi/transport"
	youtube "google.golang.org/api/youtube/v3"
)

const maxResults = 25

//SearchVideo looks to youtube for videos with a defined query value and return id and title.
func SearchVideo(ctx context.Context, googleKey string, query string) (idToTitle map[string]string, err error) {
	idToTitle = make(map[string]string)

	client := &http.Client{
		Transport: &transport.APIKey{Key: googleKey},
	}

	service, err := youtube.New(client)
	if err != nil {
		return nil, err
	}

	call := service.Search.List("id,snippet").Q(query).Type("video").MaxResults(maxResults)
	response, err := call.Do()
	if err != nil {
		return nil, err
	}

	for _, item := range response.Items {
		idToTitle[item.Id.VideoId] = item.Snippet.Title
	}
	return idToTitle, nil
}

//HandleSearch returns a function handler for alexa requests
func HandleSearch(googleKey string, service *cast.Service) alexa.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *skillserver.EchoRequest) {

		query, err := r.GetSlotValue("searchQuery")
		if err != nil {
			http.Error(w, "no searchquery slot in unmarshalled alexa request", 500)
			return
		}
		results, err := SearchVideo(ctx, googleKey, query)
		if err != nil {
			http.Error(w, "Failed to perform search", 500)
			return
		}
		for id, title := range results {
			log.Printf("First entry for search was %s: %s", id, title)
			go func() {
				playOnCast(id, service)
			}()
			writeHandleSearchVideoResponse(title, w)
			return
		}
		log.Print("Was able to search but no entries found")
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
