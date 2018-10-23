package intent

import (
	"context"
	"net/http"

	"github.com/AndreasAbdi/alexa-local-server/internal/alexa"
	"github.com/mikeflynn/go-alexa/skillserver"
)

//HandleFallback returns a function handler for unknown requests
func HandleFallback() alexa.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *skillserver.EchoRequest) {
		alexaResp := skillserver.NewEchoResponse().
			EndSession(false).
			OutputSpeech("Command examples include: play some song; pause; skip; stop; and start.")
		alexa.WriteResponse(w, alexaResp)
	}
}
