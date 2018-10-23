package infrared

//Query is a JSON Body of an infrared blaster JSON request
type Query struct {
	Data   string `json:"data"`             //ir code data (hex code/ array of int values for raw IR sequence)
	Type   string `json:"type"`             //device type (RAW, ROOMBA, SAMSUNG, etc)
	KHZ    uint64 `json:"khz,omitempty"`    //needed only for raw requests
	Length uint64 `json:"length,omitempty"` //needed except for raw/roomba signals
	Pulse  uint64 `json:"pulse,omitempty"`  //repeat a signal rapidly.
	PDelay uint64 `json:"pdelay,omitempty"` //delay between pulses
	Repeat uint64 `json:"repeat,omitempty"` //number of times to repeat the signal
	RDelay uint64 `json:"rdelay,omitempty"` //delay between signal repeats
	Out    uint64 `json:"out,omitempty"`    //which target ir sender to use.
}
