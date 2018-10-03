package intent

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/AndreasAbdi/alexa-local-server/server/alexa"
	"github.com/AndreasAbdi/alexa-local-server/server/cast"
	"github.com/mikeflynn/go-alexa/skillserver"
)

//HandlePlay commands
func HandlePlay(service *cast.Service) alexa.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *skillserver.EchoRequest) {
		log.Println("Got a play intent")
		device, err := service.GetDevice()
		if err == nil {
			go func() {
				device.MediaController.Play(time.Second * 10)
			}()
		}
		alexaResp := skillserver.NewEchoResponse()
		alexa.WriteResponse(w, alexaResp)
	}
}

//HandlePause commands
func HandlePause(service *cast.Service) alexa.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *skillserver.EchoRequest) {
		log.Println("Got a pause intent")
		device, err := service.GetDevice()
		if err == nil {
			go func() {
				device.MediaController.Pause(time.Second * 10)
			}()
		}
		alexaResp := skillserver.NewEchoResponse()
		alexa.WriteResponse(w, alexaResp)
	}
}

//HandleQuit commands
func HandleQuit(service *cast.Service) alexa.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *skillserver.EchoRequest) {
		log.Println("Got a quit intent")
		device, err := service.GetDevice()
		if err == nil {
			go func() {
				device.QuitApplication(time.Second * 10)
			}()
		}
		alexaResp := skillserver.NewEchoResponse()
		alexa.WriteResponse(w, alexaResp)
	}
}
