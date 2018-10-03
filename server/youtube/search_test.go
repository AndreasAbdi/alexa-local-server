// +build local

package youtube_test

import (
	"context"
	"testing"

	"github.com/AndreasAbdi/alexa-local-server/server/config"
	"github.com/AndreasAbdi/alexa-local-server/server/youtube"
)

func TestYoutubeSearchVideo(t *testing.T) {
	ctx := context.Background()
	conf := config.GetConfig()
	results, err := youtube.SearchVideo(ctx, conf.GoogleKey, "hi")
	if err != nil {
		t.Error(err)
	}
	t.Log("Got search results.")
	t.Log("[RESULTS] START:")
	for k, v := range results {
		t.Log(k, " : ", v)
	}
	t.Log("[RESULTS] END.")
}
