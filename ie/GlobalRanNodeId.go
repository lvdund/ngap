package ie

import "ngap/aper"

type GlobalRanNodeId struct {
	ChoiceNgRanNode ChoiceNgRanNode `bitstring:"sizeLB:0,sizeUB:150"`
}

func (ie *GlobalRanNodeId) Decode(r aper.AperReader) error {

	return nil
}

func (ie *GlobalRanNodeId) Encode(r aper.AperWriter) (err error) {

	return nil
}

type ChoiceNgRanNode struct {
	Gnb   Gnb   `bitstring:"sizeLB:0,sizeUB:150"`
	NgEnb NgEnb `bitstring:"sizeLB:0,sizeUB:150"`
	N3Iwf N3Iwf `bitstring:"sizeLB:0,sizeUB:150"`
	Tngf  Tngf  `bitstring:"sizeLB:0,sizeUB:150"`
	Twif  Twif  `bitstring:"sizeLB:0,sizeUB:150"`
	WAgf  WAgf  `bitstring:"sizeLB:0,sizeUB:150"`
}

type Gnb struct {
	GlobalGnbId GlobalGnbId `bitstring:"sizeLB:0,sizeUB:150"`
}

type NgEnb struct {
	GlobalNgEnbId GlobalNgEnbId `bitstring:"sizeLB:0,sizeUB:150"`
}

type N3Iwf struct {
	GlobalN3IwfId GlobalN3IwfId `bitstring:"sizeLB:0,sizeUB:150"`
}

type Tngf struct {
	GlobalTngfId GlobalTngfId `bitstring:"sizeLB:0,sizeUB:150"`
}

type Twif struct {
	GlobalTwifId GlobalTwifId `bitstring:"sizeLB:0,sizeUB:150"`
}

type WAgf struct {
	GlobalWAgfId GlobalWAgfId `bitstring:"sizeLB:0,sizeUB:150"`
}
