package handler

import (
	"context"
	"net/http"

	"github.com/AndreasAbdi/alexa-local-server/internal/alexa"
	"github.com/mikeflynn/go-alexa/skillserver"
)

//Creates a function handler for alexa sesion launched requests.
func handleLaunch() alexa.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, req *skillserver.EchoRequest) {
		alexaResp := skillserver.NewEchoResponse().EndSession(false).OutputSpeech("Chromecast launched. How may I help you?")
		alexa.WriteResponse(w, alexaResp)
	}
}
