package intent

import (
	"context"
	"net/http"

	"github.com/AndreasAbdi/alexa-local-server/internal/alexa"
	"github.com/AndreasAbdi/alexa-local-server/internal/cast"
	"github.com/mikeflynn/go-alexa/skillserver"
)

//HandleReset returns a function handler for unknown requests
func HandleReset(service *cast.Service) alexa.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *skillserver.EchoRequest) {
		service.Reset()
		alexaResp := skillserver.NewEchoResponse().
			OutputSpeech("Device has been reset")
		alexa.WriteResponse(w, alexaResp)
	}
}
