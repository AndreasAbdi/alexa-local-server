package youtube

import (
	"context"
	"net/http"

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

	call := service.Search.
		List("id,snippet").
		Q(query).
		Type("video").
		MaxResults(maxResults).
		Context(ctx)
	response, err := call.Do()
	if err != nil {
		return "", "", err
	}

	for _, item := range response.Items {
		return item.Id.VideoId, item.Snippet.Title, nil
	}
	return "", "", nil
}
