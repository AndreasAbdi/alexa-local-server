package youtube

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/AndreasAbdi/alexa-local-server/server/alexa"
	"github.com/mikeflynn/go-alexa/skillserver"
	"google.golang.org/api/googleapi/transport"
	youtube "google.golang.org/api/youtube/v3"
)

var (
	query      = flag.String("query", "Google", "Search term")
	maxResults = flag.Int64("max-results", 25, "Max YouTube results")
)

const key = ""

func search(ctx context.Context, query string) ([]string, error) {
	client := &http.Client{
		Transport: &transport.APIKey{Key: key},
	}

	service, err := youtube.New(client)
	if err != nil {
		return nil, err
	}

	call := service.Search.List("id,snippet").Q(query).Type("video").MaxResults(25)
	response, err := call.Do()
	if err != nil {
		return nil, err
	}

	elements := make([]string, 0)
	for _, item := range response.Items {
		elements = append(elements, item.Id.VideoId)
	}
	return elements, nil
}

//HandleSearch returns a function handler for alexa requests
func HandleSearch() alexa.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *skillserver.EchoRequest) {

		query, err := r.GetSlotValue("searchQuery")
		if err != nil {
			http.Error(w, "no searchquery slot in unmarshalled alexa request", 500)
			return
		}
		results, err := search(ctx, query)
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
