package infrared

import (
	"net/url"

	"github.com/AndreasAbdi/alexa-local-server/internal/config"
)

//Service to get the device for chromecast.
type Service struct {
	url      *url.URL
	password string
}

//NewService constructor
func NewService(config config.Wrapper) *Service {
	urlObject, err := url.Parse(config.IRBlasterAddress)
	if err != nil {
		urlObject, _ = url.Parse("http://192.168.0.1") //loopback
	}
	return &Service{
		url:      urlObject,
		password: config.IRBlasterPassword,
	}
}
