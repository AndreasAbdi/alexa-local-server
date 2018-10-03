package youtube

import (
	"context"
	"log"
	"net/http"

	"github.com/AndreasAbdi/alexa-local-server/server/alexa"
	"github.com/mikeflynn/go-alexa/skillserver"
	"google.golang.org/api/googleapi/transport"
	youtube "google.golang.org/api/youtube/v3"
)

const maxResults = 25
const key = ""

//SearchVideo looks to youtube for videos with a defined query value and return id and title.
func SearchVideo(ctx context.Context, query string) (idToTitle map[string]string, err error) {
	idToTitle = make(map[string]string)

	client := &http.Client{
		Transport: &transport.APIKey{Key: key},
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
func HandleSearch() alexa.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *skillserver.EchoRequest) {

		query, err := r.GetSlotValue("searchQuery")
		if err != nil {
			http.Error(w, "no searchquery slot in unmarshalled alexa request", 500)
			return
		}
		results, err := SearchVideo(ctx, query)
		if err != nil {
			http.Error(w, "Failed to perform search", 500)
			return
		}
		for _, result := range results {
			log.Print("First entry for search was" + result)
			return
		}
		log.Print("Was able to search")
		return
	}
}

func writeResponse(w http.ResponseWriter) {
	alexaResp := skillserver.NewEchoResponse()
	alexaResp.OutputSpeech("Finished searching")
	alexa.WriteResponse(w, alexaResp)
}
