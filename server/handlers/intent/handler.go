package intent

import (
	"context"
	"net/http"

	"github.com/AndreasAbdi/alexa-local-server/server/alexa"
	"github.com/AndreasAbdi/alexa-local-server/server/cast"
	"github.com/AndreasAbdi/alexa-local-server/server/config"
	"github.com/mikeflynn/go-alexa/skillserver"
)

const intentPlayMedia = "PlayMediaIntent"
const intentPlayYoutubeSearch = "PlayYoutubeSearchIntent"
const intentPlayYoutube = "PlayYoutubeIntent"
const intentQuit = "QuitMediaIntent"
const intentPause = "PauseIntent"
const intentPlay = "PlayIntent"
const intentMorning = "MorningIntent"
const intentHome = "HomeIntent"
const intentRewind = "RewindIntent"
const intentSkip = "SkipIntent"
const intentSeek = "SeekIntent"
const intentClearPlaylist = "ClearPlaylistIntent"
const intentAddToPlaylist = "AddToPlaylistIntent"
const intentPlayNext = "AddPlayIntent"

//HandleIntent deals with handling intent actions.
func HandleIntent(conf config.Wrapper, castService *cast.Service) alexa.HandlerFunc {
	intentToHandler := map[string]alexa.HandlerFunc{
		intentPlayMedia:         HandlePlayMedia(castService),
		intentPlayYoutube:       HandlePlayYoutube(castService),
		intentPlay:              HandlePlay(castService),
		intentQuit:              HandleQuit(castService),
		intentPause:             HandlePause(castService),
		intentRewind:            HandleRewind(castService),
		intentSkip:              HandleSkip(castService),
		intentSeek:              HandleSeek(castService),
		intentMorning:           HandleDefault(intentMorning),
		intentHome:              HandleHome(intentHome, castService),
		intentPlayYoutubeSearch: HandleSearch(conf.GoogleKey, castService),
		intentClearPlaylist:     HandleClear(castService),
		intentAddToPlaylist:     HandleAddToPlaylist(conf.GoogleKey, castService),
		intentPlayNext:          HandlePlayNext(conf.GoogleKey, castService),
	}
	return func(ctx context.Context, w http.ResponseWriter, req *skillserver.EchoRequest) {
		intent := req.GetIntentName()
		if _, ok := intentToHandler[intent]; ok {
			intentToHandler[intent](ctx, w, req)
			return
		}
		HandleDefault(intent)(ctx, w, req)
	}
}
