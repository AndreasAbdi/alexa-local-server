package infrared

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AndreasAbdi/alexa-local-server/server/config"
)

func TestServiceTurnOnTv(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//do nothing. IR blaster endpoint doesn't return anything.
	})
	server := httptest.NewServer(handler)
	defer server.Close()
	conf := config.Wrapper{
		IRBlasterAddress:  server.URL,
		IRBlasterPassword: "somefakepassword",
	}
	service := NewService(conf)
	service.SwitchTvPower()
}

func TestServiceTurnOnSoundBar(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//do nothing. IR blaster endpoint doesn't return anything.
	})
	server := httptest.NewServer(handler)
	defer server.Close()

	conf := config.Wrapper{
		IRBlasterAddress:  server.URL,
		IRBlasterPassword: "somefakepassword",
	}
	service := NewService(conf)
	service.SwitchSoundboxPower()
}
