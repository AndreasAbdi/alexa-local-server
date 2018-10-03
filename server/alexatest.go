package app

import (
	"context"
	"log"
	"net/http"

	"github.com/AndreasAbdi/alexa-local-server/server/alexa"
	"github.com/AndreasAbdi/alexa-local-server/server/cast"
	"github.com/AndreasAbdi/alexa-local-server/server/config"
	"github.com/mikeflynn/go-alexa/skillserver"
)

func (s *Server) handleAlexa(conf config.Wrapper, castService *cast.Service) http.HandlerFunc {
	alexaApp := alexa.App{
		AppID:                   conf.AlexaAppID,
		LaunchHandler:           handleFunc("LaunchRequest"),
		IntentHandler:           HandleIntent(conf, castService),
		SessionEndedHandler:     handleFunc("SessionEndedRequest"),
		AudioPlayerStateHandler: handleFunc("AudioPlayerStateChangeRequest"),
	}
	return alexa.HandleAlexaRequest(alexaApp, s.encodingService)
}

func handleFunc(requestType string) alexa.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, req *skillserver.EchoRequest) {
		log.Printf("Received a %s request", requestType)
		alexaResp := skillserver.NewEchoResponse()
		alexa.WriteResponse(w, alexaResp)
	}
}
