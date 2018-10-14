package infrared

//Service to get the device for chromecast.
type Service struct {
	initialized uint32
}

//NewService constructor
func NewService() *Service {
	return &Service{}
}
