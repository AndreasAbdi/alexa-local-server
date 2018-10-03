// +build local

package youtube_test

import (
	"context"
	"testing"

	"github.com/AndreasAbdi/alexa-local-server/server/youtube"
)

func TestYoutubeSearch(t *testing.T) {
	ctx := context.Background()
	results, err := youtube.Search(ctx, "hi")
	if err != nil {
		t.Error(err)
	}
	t.Log("Got search results.")
	t.Log("[RESULTS] START:")
	for _, result := range results {
		t.Log(result)
	}
	t.Log("[RESULTS] END.")
}
