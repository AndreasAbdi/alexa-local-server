package intent

import (
	"context"
	"log"
	"net/http"

	"github.com/AndreasAbdi/alexa-local-server/server/alexa"
	"github.com/AndreasAbdi/alexa-local-server/server/cast"
	"github.com/AndreasAbdi/alexa-local-server/server/youtube"
	"github.com/mikeflynn/go-alexa/skillserver"
)

//HandleHome commands
func HandleHome(googleKey string, service *cast.Service) alexa.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *skillserver.EchoRequest) {
		log.Println("Got a home intent")
		device, err := service.GetDevice()
		if err != nil {
			alexaResp := skillserver.NewEchoResponse()
			alexaResp.OutputSpeech("Sorry, we're having internal issues right now")
			alexa.WriteResponse(w, alexaResp)
			return
		}
		id, _, err := youtube.SearchVideo(ctx, googleKey, "lofi hip hop radio")
		go func() {
			device.PlayYoutubeVideo(id)
		}()

		alexaResp := skillserver.NewEchoResponse()
		alexaResp.OutputSpeech("starting up")
		alexa.WriteResponse(w, alexaResp)
	}
}
