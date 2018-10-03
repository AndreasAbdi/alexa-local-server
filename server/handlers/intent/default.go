package intent

import (
	"context"
	"net/http"

	"github.com/AndreasAbdi/alexa-local-server/server/alexa"
	"github.com/mikeflynn/go-alexa/skillserver"
)

//HandleDefault returns a function handler for unknown requests
func HandleDefault(intentType string) alexa.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *skillserver.EchoRequest) {
		alexaResp := skillserver.NewEchoResponse()
		alexaResp.OutputSpeech("I'm sorry, cast doesn't know what to do with intents of type " + intentType)
		alexa.WriteResponse(w, alexaResp)
	}
}
