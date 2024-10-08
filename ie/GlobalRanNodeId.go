package ie

import (
	"fmt"
	"ngap/aper"
)

const (
	RAN_NODE_ID_GNB uint64 = iota + 1
	RAN_NODE_ID_NGENB
	//add more, should be generated
)

type GlobalRanNodeId struct {
	Choice uint64

	GlobalGnbId *GlobalGnbId
	NgEnb       *NgEnb
	//generate more choices
}

func (e *GlobalRanNodeId) Decode(r *aper.AperReader) (err error) {
	if e.Choice, err = r.ReadChoice(6, false); err != nil {
		return
	}
	fmt.Printf("Choice for RanId is: %d\n", e.Choice)
	switch e.Choice {
	case RAN_NODE_ID_GNB:
		var tmp GlobalGnbId
		if err = tmp.Decode(r); err != nil {
			return
		}
		e.GlobalGnbId = &tmp
	case RAN_NODE_ID_NGENB:
		//TODO:
		//generate more cases
	}

	return
}

func (ie *GlobalRanNodeId) Encode(r *aper.AperWriter) (err error) {

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
