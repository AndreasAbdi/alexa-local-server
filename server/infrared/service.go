package infrared

import "net/url"

//Service to get the device for chromecast.
type Service struct {
	initialized uint32
	url         *url.URL
	password    string
}

//NewService constructor
func NewService() *Service {
	return &Service{}
}
