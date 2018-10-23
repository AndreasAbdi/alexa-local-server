package intent

import (
	"context"
	"log"
	"net/http"

	"github.com/AndreasAbdi/alexa-local-server/internal/alexa"
	"github.com/AndreasAbdi/alexa-local-server/internal/cast"
	"github.com/mikeflynn/go-alexa/skillserver"
)

//HandlePlayMedia commands
func HandlePlayMedia(service *cast.Service) alexa.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *skillserver.EchoRequest) {
		log.Println("Got a play media intent")
		device, err := service.GetDevice()
		if err == nil {
			go func() {
				device.PlayMedia(
					"http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/images/BigBuckBunny.jpg",
					"image/jpeg")
			}()
		}
		alexaResp := skillserver.NewEchoResponse()
		alexa.WriteResponse(w, alexaResp)
	}
}

//HandlePlayYoutube commands
func HandlePlayYoutube(service *cast.Service) alexa.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *skillserver.EchoRequest) {
		log.Println("Got a play youtube intent")
		device, err := service.GetDevice()
		if err == nil {
			go func() {
				device.PlayYoutubeVideo("F1B9Fk_SgI0")
			}()
		}
		alexaResp := skillserver.NewEchoResponse()
		alexa.WriteResponse(w, alexaResp)
	}
}
