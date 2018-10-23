// +build local

package infrared

import (
	"testing"

	"github.com/AndreasAbdi/alexa-local-server/config"
)

func TestServiceTurnOnTv(t *testing.T) {
	conf := config.GetConfig()
	service := NewService(conf)
	service.SwitchTvPower()
}

func TestServiceTurnOnAc(t *testing.T) {
	conf := config.GetConfig()
	service := NewService(conf)
	service.SetACChill()
}
