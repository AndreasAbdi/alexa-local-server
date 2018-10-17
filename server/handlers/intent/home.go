package intent

import (
	"context"
	"log"
	"net/http"

	"github.com/AndreasAbdi/alexa-local-server/server/alexa"
	"github.com/AndreasAbdi/alexa-local-server/server/cast"
	"github.com/AndreasAbdi/alexa-local-server/server/infrared"
	"github.com/AndreasAbdi/alexa-local-server/server/youtube"
	"github.com/mikeflynn/go-alexa/skillserver"
)

//HandleHome commands
func HandleHome(googleKey string, service *cast.Service, infraService *infrared.Service) alexa.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *skillserver.EchoRequest) {
		log.Println("Got a home intent")
		go func() {
			infraService.SwitchTvPower()
			infraService.SwitchSoundboxPower()
		}()
		device, err := service.GetDevice()
		if err != nil {
			alexa.WriteSpeech(w, "Sorry, we're having internal issues right now")
			return
		}
		id, _, err := youtube.SearchVideo(ctx, googleKey, "chilledcow")
		go func() {
			device.PlayYoutubeVideo(id)
		}()

		alexa.WriteSpeech(w, "starting up")
	}
}
