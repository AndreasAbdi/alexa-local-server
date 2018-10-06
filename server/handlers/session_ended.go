package handler

import (
	"context"
	"net/http"

	"github.com/AndreasAbdi/alexa-local-server/server/alexa"
	"github.com/mikeflynn/go-alexa/skillserver"
)

//Creates a function handler for alexa session ended command.
func handleSessionEnded() alexa.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, req *skillserver.EchoRequest) {
		alexaResp := skillserver.NewEchoResponse().EndSession(false).OutputSpeech("Leaving")
		alexa.WriteResponse(w, alexaResp)
	}
}
