package intent

import (
	"context"
	"net/http"

	"github.com/AndreasAbdi/alexa-local-server/server/infrared"

	"github.com/AndreasAbdi/alexa-local-server/server/alexa"
	"github.com/AndreasAbdi/alexa-local-server/server/cast"
	"github.com/AndreasAbdi/alexa-local-server/server/config"
	"github.com/mikeflynn/go-alexa/skillserver"
)

//custom intents
const intentMorning = "MorningIntent"
const intentHome = "HomeIntent"
const intentShutdown = "ShutdownIntent"

//infrared control intents
const intentToggleTv = "TurnOnTVIntent"
const intentToggleSoundbar = "ToggleSoundBarIntent"
const intentMuteSoundbar = "MuteSoundBarIntent"
const intentIncreaseVolume = "IncreaseVolumeSoundBarIntent"
const intentDecreaseVolume = "DecreaseVolumeSoundBarIntent"
const intentACTurnOn = "TurnOnACIntent"
const intentACTurnOff = "TurnOffACIntent"
const intentSwitchToFanAC = "TurnFanACIntent"
const intentSwitchToChillAC = "TurnChillACIntent"
const intentSwitchToHeatAC = "TurnHeatACIntent"

//general youtube and media intents
const intentPlayMedia = "PlayMediaIntent"
const intentPlayYoutubeSearch = "PlayYoutubeSearchIntent"
const intentPlayYoutube = "PlayYoutubeIntent"
const intentQuit = "QuitMediaIntent"
const intentPause = "PauseIntent"
const intentPlay = "PlayIntent"

//for resetting chromecast device linkage.
const intentResetDevice = "ResetIntent"

const intentRewind = "RewindIntent"
const intentSkip = "SkipIntent"
const intentSeek = "SeekIntent"
const intentClearPlaylist = "ClearPlaylistIntent"
const intentAddToPlaylist = "AddToPlaylistIntent"
const intentPlayNext = "PlayNextIntent"

//these are the built in intents to deal with
const intentFallback = "AMAZON.FallbackIntent"
const intentCancel = "AMAZON.CancelIntent"
const intentHelp = "AMAZON.HelpIntent"
const intentStop = "AMAZON.StopIntent"
const intentNext = "AMAZON.NextIntent"

//HandleIntent deals with handling intent actions.
func HandleIntent(conf config.Wrapper, castService *cast.Service, infraService *infrared.Service) alexa.HandlerFunc {
	intentToHandler := map[string]alexa.HandlerFunc{
		intentPlayMedia:         HandlePlayMedia(castService),
		intentPlayYoutube:       HandleHome(conf.GoogleKey, castService, infraService),
		intentPlay:              HandlePlay(castService),
		intentQuit:              HandleQuit(castService),
		intentPause:             HandlePause(castService),
		intentRewind:            HandleRewind(castService),
		intentSkip:              HandleSkip(castService),
		intentNext:              HandleSkip(castService),
		intentSeek:              HandleSeek(castService),
		intentShutdown:          HandleShutdown(castService, infraService),
		intentMorning:           HandleHome(conf.GoogleKey, castService, infraService),
		intentHome:              HandleHome(conf.GoogleKey, castService, infraService),
		intentPlayYoutubeSearch: HandleSearch(conf.GoogleKey, castService),
		intentClearPlaylist:     HandleClear(castService),
		intentAddToPlaylist:     HandleAddToPlaylist(conf.GoogleKey, castService),
		intentPlayNext:          HandlePlayNext(conf.GoogleKey, castService),
		intentFallback:          HandleFallback(),
		intentHelp:              HandleHelp(),
		intentResetDevice:       HandleReset(castService),
		intentStop:              HandleQuit(castService),
		intentCancel:            HandleQuit(castService),
		intentToggleTv:          HandleTVSwitch(infraService),
		intentToggleSoundbar:    HandleSoundBarSwitch(infraService),
		intentMuteSoundbar:      HandleSoundBarMute(infraService),
		intentIncreaseVolume:    HandleSoundBarIncreaseVolume(infraService),
		intentDecreaseVolume:    HandleSoundBarDecreaseVolume(infraService),
		intentACTurnOn:          HandleACTurnOn(infraService),
		intentACTurnOff:         HandleACTurnOff(infraService),
		intentSwitchToChillAC:   HandleACSwitchToChill(infraService),
		intentSwitchToFanAC:     HandleACSwitchToFan(infraService),
		intentSwitchToHeatAC:    HandleACSwitchToHeat(infraService),
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
