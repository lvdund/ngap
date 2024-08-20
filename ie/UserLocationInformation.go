package ie

type UserLocationInformation struct {
	ChoiceUserLocationInformation *ChoiceUserLocationInformation
}

type ChoiceUserLocationInformation struct {
	EUtraUserLocationInformation *EUtraUserLocationInformation
	NrUserLocationInformation    *NrUserLocationInformation
	N3IwfUserLocationInformation *N3IwfUserLocationInformation
	TngfUserLocationInformation  *TngfUserLocationInformation
	TwifUserLocationInformation  *TwifUserLocationInformation
	WAgfUserLocationInformation  *WAgfUserLocationInformation
}

type EUtraUserLocationInformation struct {
	EUtraCgi          *EUtraCgi
	Tai               *Tai
	AgeOfLocation     *TimeStamp
	PscellInformation *NgRanCgi
}

type NrUserLocationInformation struct {
	NrCgi               *NrCgi
	Tai                 *Tai
	AgeOfLocation       *TimeStamp
	PscellInformation   *NgRanCgi
	Nid                 *Nid
	NrNtnTaiInformation *NrNtnTaiInformation
}

type N3IwfUserLocationInformation struct {
	IpAddress  *TransportLayerAddress
	PortNumber []byte //`bitstring:"sizeLB:2,sizeUB:2"`
	Tai        *Tai
}

type TngfUserLocationInformation struct {
	TnapId     *[]byte
	IpAddress  *TransportLayerAddress
	PortNumber []byte //`bitstring:"sizeLB:2,sizeUB:2"`
	Tai        *Tai
}

type TwifUserLocationInformation struct {
	TwapId     *[]byte
	IpAddress  *TransportLayerAddress
	PortNumber []byte //`bitstring:"sizeLB:2,sizeUB:2"`
	Tai        *Tai
}
