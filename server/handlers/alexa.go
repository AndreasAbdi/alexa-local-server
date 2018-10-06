package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/AndreasAbdi/alexa-local-server/server/handlers/intent"

	"github.com/AndreasAbdi/alexa-local-server/server/alexa"
	"github.com/AndreasAbdi/alexa-local-server/server/cast"
	"github.com/AndreasAbdi/alexa-local-server/server/config"
	"github.com/mikeflynn/go-alexa/skillserver"
)

//HandleAlexa for alexa http requests
func HandleAlexa(conf config.Wrapper, castService *cast.Service) http.HandlerFunc {
	alexaApp := alexa.App{
		AppID:                   conf.AlexaAppID,
		LaunchHandler:           handleLaunch(),
		IntentHandler:           intent.HandleIntent(conf, castService),
		SessionEndedHandler:     handleSessionEnded(),
		AudioPlayerStateHandler: handleFunc("AudioPlayerStateChangeRequest"),
	}
	return alexa.HandleAlexaRequest(alexaApp)
}

func handleFunc(requestType string) alexa.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, req *skillserver.EchoRequest) {
		log.Printf("Received a %s request", requestType)
		alexaResp := skillserver.NewEchoResponse()
		alexa.WriteResponse(w, alexaResp)
	}
}
