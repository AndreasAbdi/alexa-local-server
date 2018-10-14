package infrared

import (
	"log"
	"path"

	"github.com/imroc/req"
)

//SwitchTvPower using ir blaster
func (s *Service) SwitchTvPower() {
	request := req.New()
	s.url.Path = path.Join(s.url.Path, "msg")
	params := req.QueryParam{
		keyPassword: s.password,
	}
	_, err := request.Get(s.url.String(), params)
	if err != nil {
		log.Println("Error with sending switch tv power request", err)
	}
}
