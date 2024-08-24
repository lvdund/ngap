package ie

import "ngap/aper"

type WAgfUserLocationInformation struct {
	ChoiceWAgfUserLocationInformation ChoiceWAgfUserLocationInformation //`bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceWAgfUserLocationInformation struct {
	GlobalLineId     GlobalLineId     //`bitstring:"sizeLB:0,sizeUB:150"`
	HfcNodeId        HfcNodeId        //`bitstring:"sizeLB:0,sizeUB:150"`
	GlobalCableId    GlobalCableId    //`bitstring:"sizeLB:0,sizeUB:150"`
	HfcNodeIdNew     HfcNodeIdNew     //`bitstring:"sizeLB:0,sizeUB:150"`
	GlobalCableIdNew GlobalCableIdNew //`bitstring:"sizeLB:0,sizeUB:150"`
}

type GlobalLineId struct {
	GlobalLineId aper.OctetString //`octetstring:"sizeLB:0,sizeUB:150"`
	LineType     []byte           //`bitstring:"sizeLB:0,sizeUB:150"`
	Tai          Tai              //`bitstring:"sizeLB:0,sizeUB:150"`
}

type HfcNodeId struct {
	HfcNodeId aper.OctetString //`octetstring:"sizeLB:0,sizeUB:150"`
}

type GlobalCableId struct {
	GlobalCableId aper.OctetString //`octetstring:"sizeLB:0,sizeUB:150"`
}

type HfcNodeIdNew struct {
	HfcNodeId aper.OctetString //`octetstring:"sizeLB:0,sizeUB:150"`
	Tai       Tai              //`bitstring:"sizeLB:0,sizeUB:150"`
}

type GlobalCableIdNew struct {
	GlobalCableId aper.OctetString //`octetstring:"sizeLB:0,sizeUB:150"`
	Tai           Tai              //`bitstring:"sizeLB:0,sizeUB:150"`
}
