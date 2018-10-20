package intent

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/AndreasAbdi/alexa-local-server/server/alexa"
	"github.com/AndreasAbdi/alexa-local-server/server/cast"
	gocast "github.com/AndreasAbdi/gochromecast"
	"github.com/mikeflynn/go-alexa/skillserver"
)

const slotSeekTime = "seekTime"

//HandlePlay commands
func HandlePlay(service *cast.Service) alexa.HandlerFunc {
	return getGenericControlFunc(service, intentPlay, func(device *gocast.Device) {
		device.MediaController.Play(time.Second * 10)
	})
}

//HandlePause commands
func HandlePause(service *cast.Service) alexa.HandlerFunc {
	return getGenericControlFunc(service, intentPause, func(device *gocast.Device) {
		device.MediaController.Pause(time.Second * 10)
	})
}

//HandleClear comamnds
func HandleClear(service *cast.Service) alexa.HandlerFunc {
	return getGenericControlFunc(service, intentClearPlaylist, func(device *gocast.Device) {
		device.YoutubeController.ClearPlaylist()
	})
}

//HandleSkip commands
func HandleSkip(service *cast.Service) alexa.HandlerFunc {
	return getGenericControlFunc(service, intentSeek, func(device *gocast.Device) {
		device.MediaController.Skip(time.Second * 10)
	})
}

//HandleRewind commands
func HandleRewind(service *cast.Service) alexa.HandlerFunc {
	return getGenericControlFunc(service, intentRewind, func(device *gocast.Device) {
		device.MediaController.Rewind(time.Second * 10)
	})
}

//HandleQuit commands
func HandleQuit(service *cast.Service) alexa.HandlerFunc {
	return getGenericControlFunc(service, intentQuit, func(device *gocast.Device) {
		device.QuitApplication(time.Second * 10)
	})
}

//HandleSeek commands
func HandleSeek(service *cast.Service) alexa.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *skillserver.EchoRequest) {
		seekTimeString, err := r.GetSlotValue(slotSeekTime)
		if err != nil {
			http.Error(w, "no seek time slot in unmarshalled alexa request", 500)
			return
		}
		seekTime, err := time.ParseDuration(seekTimeString)
		if err != nil {
			http.Error(w, "Failed to convert slot value for seek time to float64", 500)
			return
		}
		log.Printf("Got a seek request")
		go func() {
			device, err := service.GetDevice()
			if err == nil {

				device.MediaController.Seek(seekTime.Seconds(), time.Second*10)
			}
		}()
		alexaResp := skillserver.NewEchoResponse()
		alexa.WriteResponse(w, alexaResp)
	}
}

func getGenericControlFunc(service *cast.Service, name string, controlCommand func(device *gocast.Device)) alexa.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *skillserver.EchoRequest) {
		log.Printf("Got a %s request", name)
		go func() {
			device, err := service.GetDevice()
			if err == nil {
				controlCommand(device)
			}
		}()
		alexaResp := skillserver.NewEchoResponse()
		alexa.WriteResponse(w, alexaResp)
	}
}
