package intent

import (
	"context"
	"log"
	"net/http"

	"github.com/AndreasAbdi/alexa-local-server/server/alexa"
	"github.com/AndreasAbdi/alexa-local-server/server/infrared"
	"github.com/mikeflynn/go-alexa/skillserver"
)

//HandleTVSwitch infrared commands
func HandleTVSwitch(service *infrared.Service) alexa.HandlerFunc {
	return getGenericInfraredFunc(intentToggleTv, func() {
		service.SwitchTvPower()
	})
}

//HandleSoundBarSwitch infrared commands
func HandleSoundBarSwitch(service *infrared.Service) alexa.HandlerFunc {
	return getGenericInfraredFunc(intentToggleSoundbar, func() {
		service.SwitchSoundboxPower()
	})
}

//HandleSoundBarMute infrared commands
func HandleSoundBarMute(service *infrared.Service) alexa.HandlerFunc {
	return getGenericInfraredFunc(intentMuteSoundbar, func() {
		service.SwitchTvPower()
	})
}

//HandleSoundBarIncreaseVolume infrared commands
func HandleSoundBarIncreaseVolume(service *infrared.Service) alexa.HandlerFunc {
	return getGenericInfraredFunc(intentIncreaseVolume, func() {
		service.VolumeIncreaseSoundbox()
	})
}

//HandleSoundBarDecreaseVolume infrared commands
func HandleSoundBarDecreaseVolume(service *infrared.Service) alexa.HandlerFunc {
	return getGenericInfraredFunc(intentDecreaseVolume, func() {
		service.VolumeDecreaseSoundbox()
	})
}

//HandleACTurnOn infrared commands
func HandleACTurnOn(service *infrared.Service) alexa.HandlerFunc {
	return getGenericInfraredFunc(intentACTurnOn, func() {
		service.SetACChill() //ac doesn't have set to on mode available, always sets to chill.
	})
}

//HandleACTurnOff infrared commands
func HandleACTurnOff(service *infrared.Service) alexa.HandlerFunc {
	return getGenericInfraredFunc(intentACTurnOff, func() {
		service.SetACOff()
	})
}

//HandleACSwitchToChill infrared commands
func HandleACSwitchToChill(service *infrared.Service) alexa.HandlerFunc {
	return getGenericInfraredFunc(intentSwitchToChillAC, func() {
		service.SetACChill()
	})
}

//HandleACSwitchToFan infrared commands
func HandleACSwitchToFan(service *infrared.Service) alexa.HandlerFunc {
	return getGenericInfraredFunc(intentSwitchToFanAC, func() {
		service.SetACFan()
	})
}

//HandleACSwitchToHeat infrared commands
func HandleACSwitchToHeat(service *infrared.Service) alexa.HandlerFunc {
	return getGenericInfraredFunc(intentSwitchToHeatAC, func() {
		service.SetACHeat()
	})
}

func getGenericInfraredFunc(name string, controlCommand func()) alexa.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *skillserver.EchoRequest) {
		log.Printf("Got a %s request", name)
		go func() {
			controlCommand()

		}()
		alexaResp := skillserver.NewEchoResponse()
		alexa.WriteResponse(w, alexaResp)
	}
}
