package app

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/AndreasAbdi/alexa-local-server/server/alexa"
	"github.com/AndreasAbdi/alexa-local-server/server/youtube"
	"github.com/mikeflynn/go-alexa/skillserver"
)

//HandleIntent deals with handling intent actions.
func HandleIntent() alexa.HandlerFunc {

	return func(ctx context.Context, w http.ResponseWriter, req *skillserver.EchoRequest) {
		intent := req.GetIntentName()
		device, err := getDevice()
		if err != nil {
			log.Print("Failed to get device")
		}

		log.Println("Intent type is " + intent)
		switch intent {
		case "PlayMediaIntent":
			log.Println("Got a Play Media Intent")
			go func() {
				device.PlayMedia(
					"http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/images/BigBuckBunny.jpg",
					"image/jpeg")
			}()
		case "PlayYoutubeSearchIntent":
			log.Println("Got a Play Youtube Search Intent")
			youtube.HandleSearch()(ctx, w, req)
		case "PlayYoutubeIntent":
			log.Println("Got a play youtube intent")
			go func() { device.PlayYoutubeVideo("F1B9Fk_SgI0") }()
		case "QuitMediaIntent":
			log.Println("Got a quit media intent")
			device.QuitApplication(time.Second * 10)
		case "PauseIntent":
			log.Println("Got a pause media intent")
			device.MediaController.Pause(time.Second * 10)
		case "PlayIntent":
			log.Println("Got a play intent")
			device.MediaController.Play(time.Second * 10)
		case "MorningIntent":
			log.Println("Got a morning request")
		case "HomeIntent":
			log.Println("Got a welcome home intent")
		default:
			log.Println("IDK what to do with this intent type: " + intent)
		}
		alexaResp := skillserver.NewEchoResponse()
		alexa.WriteResponse(w, alexaResp)
	}
}
