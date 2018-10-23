package infrared

const tvPowerCode string = "E0E040BF"
const tvType string = "SAMSUNG"
const tvPowerCodeLength uint64 = 32

//SwitchTvPower using ir blaster
func (s *Service) SwitchTvPower() {
	sendMessageDefault(*s.url, s.password, tvPowerCode, tvType, tvPowerCodeLength, nil)
}
