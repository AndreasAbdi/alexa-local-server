package cast

import (
	"sync"
	"sync/atomic"
	"time"

	castv2 "github.com/AndreasAbdi/go-castv2"
)

//Service to get the device for chromecast.
type Service struct {
	device      *castv2.Device
	initialized uint32
	mutex       *sync.Mutex
}

//NewService constructor
func NewService() *Service {
	return &Service{
		mutex: &sync.Mutex{},
	}
}

//GetDevice from the service.
func (s *Service) GetDevice() (*castv2.Device, error) {

	if atomic.LoadUint32(&s.initialized) == 1 {
		return s.device, nil
	}

	return s.setDevice()
}

//Reset the service.
func (s *Service) Reset() {
	s.setDevice()
}

// NoChromecastError describing that there was no chromecast to find.
type NoChromecastError struct {
}

func (e *NoChromecastError) Error() string {
	return "Was not able to find a chromecast device"
}

func getDevice() (*castv2.Device, error) {
	devices := make(chan *castv2.Device, 100)
	castv2.FindDevices(time.Second*5, devices)
	for device := range devices {
		return device, nil
	}
	return nil, &NoChromecastError{}
}

func (s *Service) setDevice() (*castv2.Device, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	device, err := getDevice()
	if err != nil {
		return nil, err
	}
	s.device = device
	atomic.StoreUint32(&s.initialized, 1)
	return device, nil
}
