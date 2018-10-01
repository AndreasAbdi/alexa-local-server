package app

import (
	"time"

	castv2 "github.com/AndreasAbdi/go-castv2"
)

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
