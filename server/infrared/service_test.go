package infrared

import (
	"testing"

	"github.com/AndreasAbdi/alexa-local-server/server/config"
)

func TestServiceTurnOnTv(t *testing.T) {
	conf := config.GetConfig()
	service := NewService(conf)
	service.SwitchTvPower()
}
