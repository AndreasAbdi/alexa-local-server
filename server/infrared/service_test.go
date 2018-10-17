package infrared

import (
	"testing"

	"github.com/AndreasAbdi/alexa-local-server/server/config"
)

func TestServiceTurnOnTv(t *testing.T) {
	conf := config.Wrapper{
		IRBlasterAddress:  "http://sometestaddress.com",
		IRBlasterPassword: "somefakepassword",
	}
	service := NewService(conf)
	service.SwitchTvPower()
}

func TestServiceTurnOnSoundBar(t *testing.T) {
	conf := config.Wrapper{
		IRBlasterAddress:  "http://sometestaddress.com",
		IRBlasterPassword: "somefakepassword",
	}
	service := NewService(conf)
	service.SwitchSoundboxPower()
}
