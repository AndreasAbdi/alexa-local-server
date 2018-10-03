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
	id, title, err := youtube.SearchVideo(ctx, conf.GoogleKey, "summertime magic")
	if err != nil {
		t.Error(err)
	}
	t.Logf("Got search results. title: %s , id: %s", title, id)
}
