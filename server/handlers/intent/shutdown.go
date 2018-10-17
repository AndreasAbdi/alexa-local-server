package intent

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/AndreasAbdi/alexa-local-server/server/alexa"
	"github.com/AndreasAbdi/alexa-local-server/server/cast"
	"github.com/AndreasAbdi/alexa-local-server/server/infrared"
	"github.com/mikeflynn/go-alexa/skillserver"
)

//HandleShutdown commands
func HandleShutdown(castService *cast.Service, infraService *infrared.Service) alexa.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *skillserver.EchoRequest) {
		log.Println("Got a shutdown intent")
		go func() {
			infraService.SwitchTvPower()
			infraService.SwitchSoundboxPower()
		}()
		device, err := castService.GetDevice()
		if err != nil {
			alexa.WriteSpeech(w, "Sorry, we're having internal issues right now")
			return
		}
		go func() {
			device.QuitApplication(time.Second * 10)
		}()
		alexa.WriteEmpty(w)
	}
}
