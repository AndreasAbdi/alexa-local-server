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

//these are the built in intents to deal with
const intentFallback = "FallbackIntent"
const intentCancel = "CancelIntent"
const intentHelp = "HelpIntent"
const intentStop = "StopIntent"
const intentNext = "NextIntent"

//HandleIntent deals with handling intent actions.
func HandleIntent(conf config.Wrapper, castService *cast.Service) alexa.HandlerFunc {
	intentToHandler := map[string]alexa.HandlerFunc{
		intentPlayMedia:         HandlePlayMedia(castService),
		intentPlayYoutube:       HandleHome(conf.GoogleKey, castService),
		intentPlay:              HandlePlay(castService),
		intentQuit:              HandleQuit(castService),
		intentPause:             HandlePause(castService),
		intentRewind:            HandleRewind(castService),
		intentSkip:              HandleSkip(castService),
		intentNext:              HandleSkip(castService),
		intentSeek:              HandleSeek(castService),
		intentMorning:           HandleHome(conf.GoogleKey, castService),
		intentHome:              HandleHome(conf.GoogleKey, castService),
		intentPlayYoutubeSearch: HandleSearch(conf.GoogleKey, castService),
		intentClearPlaylist:     HandleClear(castService),
		intentAddToPlaylist:     HandleAddToPlaylist(conf.GoogleKey, castService), //todo
		intentPlayNext:          HandlePlayNext(conf.GoogleKey, castService),      //todo
		intentFallback:          HandleFallback(),
		intentHelp:              HandleHelp(),
		intentStop:              HandleQuit(castService),
		intentCancel:            HandleQuit(castService),
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
