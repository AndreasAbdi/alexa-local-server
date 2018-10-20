package intent

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/AndreasAbdi/alexa-local-server/server/alexa"
	"github.com/AndreasAbdi/alexa-local-server/server/infrared"
	"github.com/mikeflynn/go-alexa/skillserver"
)

const slotVolumeIncrease string = "volumeDelta"
const slotVolumeDecrease string = "volumeDelta"

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
	return func(ctx context.Context, w http.ResponseWriter, r *skillserver.EchoRequest) {
		log.Printf("Got a %s request", intentIncreaseVolume)
		go func() {
			volumeDelta := getValueOrZero(r, slotVolumeIncrease)
			service.VolumeIncreaseSoundbox(volumeDelta)

		}()
		alexaResp := skillserver.NewEchoResponse()
		alexa.WriteResponse(w, alexaResp)
	}
}

//HandleSoundBarDecreaseVolume infrared commands
func HandleSoundBarDecreaseVolume(service *infrared.Service) alexa.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *skillserver.EchoRequest) {
		log.Printf("Got a %s request", intentDecreaseVolume)
		go func() {
			volumeDelta := getValueOrZero(r, slotVolumeDecrease)
			service.VolumeDecreaseSoundbox(volumeDelta)
		}()
		alexaResp := skillserver.NewEchoResponse()
		alexa.WriteResponse(w, alexaResp)
	}
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

func getValueOrZero(r *skillserver.EchoRequest, slotType string) uint64 {
	var volumeDelta uint64
	increaseBy, err := r.GetSlotValue(slotType)
	if err != nil {
		return 0
	}
	log.Printf(increaseBy)
	volumeDelta, err = strconv.ParseUint(increaseBy, 10, 64)
	if err != nil {
		return 0
	}
	return volumeDelta
}
