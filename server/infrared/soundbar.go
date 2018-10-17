package infrared

const soundBarType string = "SAMSUNG"
const soundBarCodeLength uint64 = 32
const soundBarPowerCode string = "34347887"
const soundBarMuteCode string = "3434F807"
const soundBarIncreaseSoundCode string = "3434E817"
const soundBarDecreaseSoundCode string = "34346897"

//SwitchSoundboxPower using ir blaster
func (s *Service) SwitchSoundboxPower() {
	sendMessageDefault(*s.url, s.password, soundBarPowerCode, soundBarType, tvPowerCodeLength)
}

//VolumeDecreaseSoundbox using ir blaster
func (s *Service) VolumeDecreaseSoundbox() {
	sendMessageDefault(*s.url, s.password, soundBarDecreaseSoundCode, soundBarType, tvPowerCodeLength)
}

//VolumeIncreaseSoundbox using ir blaster
func (s *Service) VolumeIncreaseSoundbox() {
	sendMessageDefault(*s.url, s.password, soundBarIncreaseSoundCode, soundBarType, tvPowerCodeLength)
}

//MuteSoundbox using ir blaster
func (s *Service) MuteSoundbox() {
	sendMessageDefault(*s.url, s.password, soundBarMuteCode, soundBarType, tvPowerCodeLength)
}
