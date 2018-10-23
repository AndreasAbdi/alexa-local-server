package infrared

import (
	"fmt"
	"strings"
)

const acType string = "RAW"
const acKHZ uint64 = 38
const acPulse uint64 = 3

var codeACTurnOff = []uint64{9144, 4682, 478, 642, 558, 670, 504, 618, 556, 1720, 520, 696, 478, 648, 526, 1724, 526, 710, 480, 644, 556, 670, 476, 656, 518, 1736, 558, 616, 558, 626, 550, 616, 554, 1716, 552, 622, 526, 1724, 554, 616, 558, 634, 530, 652, 550, 616, 530, 644, 530, 714, 478, 650, 552, 614, 562, 668, 478, 714, 484, 648, 546, 612, 562, 616, 534, 614, 554}

var codeACChill = []uint64{9134, 4628, 534, 654, 520, 640, 534, 656, 518, 1748, 520, 640, 534, 640, 534, 1730, 520, 658, 558, 626, 520, 644, 528, 656, 518, 1736, 532, 1732, 518, 656, 518, 654, 518, 672, 518, 1732, 522, 652, 520, 654, 520, 660, 532, 644, 530, 656, 518, 642, 532, 1748, 520, 640, 560, 626, 520, 656, 516, 664, 530, 654, 520, 654, 518, 654, 518, 628, 520}

var codeACFan = []uint64{9174, 4614, 522, 640, 534, 654, 518, 654, 520, 1736, 534, 654, 520, 654, 520, 1716, 560, 646, 520, 640, 534, 640, 532, 644, 528, 1748, 520, 640, 534, 1730, 520, 652, 522, 674, 518, 1730, 520, 640, 532, 652, 522, 670, 520, 656, 520, 640, 534, 638, 560, 1724, 518, 656, 520, 1730, 520, 1732, 520, 670, 518, 656, 518, 654, 520, 654, 520, 614, 534}

var codeACHeat = []uint64{9148, 4640, 520, 650, 524, 642, 532, 654, 518, 1736, 532, 658, 516, 654, 520, 1730, 520, 672, 518, 642, 532, 640, 560, 628, 518, 1736, 534, 654, 520, 640, 560, 628, 654, 1720, 530, 1730, 520, 640, 534, 654, 520, 658, 534, 654, 518, 654, 520, 646, 528, 1748, 520, 646, 526, 640, 558, 616, 532, 660, 532, 640, 532, 640, 532, 640, 562, 602, 518}

//SetACOff using ir blaster
func (s *Service) SetACOff() {
	query := constructACMessage(codeACTurnOff)
	sendMessageJSON(*s.url, s.password, query)
}

//SetACChill using ir blaster
func (s *Service) SetACChill() {
	query := constructACMessage(codeACChill)
	sendMessageJSON(*s.url, s.password, query)
}

//SetACHeat using ir blaster
func (s *Service) SetACHeat() {
	query := constructACMessage(codeACHeat)
	sendMessageJSON(*s.url, s.password, query)
}

//SetACFan using ir blaster
func (s *Service) SetACFan() {
	query := constructACMessage(codeACFan)
	sendMessageJSON(*s.url, s.password, query)
}

func constructACMessage(code []uint64) Query {
	return Query{
		Type:  acType,
		Data:  strings.Join(strings.Split(fmt.Sprint(code), " "), ""),
		KHZ:   acKHZ,
		Pulse: acPulse,
	}
}
