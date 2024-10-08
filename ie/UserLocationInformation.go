package ie

import "ngap/aper"

type UserLocationInformation struct {
	ChoiceUserLocationInformation ChoiceUserLocationInformation `bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceUserLocationInformation struct {
	EUtraUserLocationInformation EUtraUserLocationInformation `bitstring:"sizeLB:0,sizeUB:150"`
	NrUserLocationInformation    NrUserLocationInformation    `bitstring:"sizeLB:0,sizeUB:150"`
	N3IwfUserLocationInformation N3IwfUserLocationInformation `bitstring:"sizeLB:0,sizeUB:150"`
	TngfUserLocationInformation  TngfUserLocationInformation  `bitstring:"sizeLB:0,sizeUB:150"`
	TwifUserLocationInformation  TwifUserLocationInformation  `bitstring:"sizeLB:0,sizeUB:150"`
	WAgfUserLocationInformation  WAgfUserLocationInformation  `bitstring:"sizeLB:0,sizeUB:150"`
}

type EUtraUserLocationInformation struct {
	EUtraCgi          EUtraCgi  `bitstring:"sizeLB:0,sizeUB:150"`
	Tai               Tai       `bitstring:"sizeLB:0,sizeUB:150"`
	AgeOfLocation     TimeStamp `bitstring:"sizeLB:0,sizeUB:150"`
	PscellInformation NgRanCgi  `bitstring:"sizeLB:0,sizeUB:150"`
}

type NrUserLocationInformation struct {
	NrCgi               NrCgi               `bitstring:"sizeLB:0,sizeUB:150"`
	Tai                 Tai                 `bitstring:"sizeLB:0,sizeUB:150"`
	AgeOfLocation       TimeStamp           `bitstring:"sizeLB:0,sizeUB:150"`
	PscellInformation   NgRanCgi            `bitstring:"sizeLB:0,sizeUB:150"`
	Nid                 Nid                 `bitstring:"sizeLB:0,sizeUB:150"`
	NrNtnTaiInformation NrNtnTaiInformation `bitstring:"sizeLB:0,sizeUB:150"`
}

type N3IwfUserLocationInformation struct {
	IpAddress  TransportLayerAddress `bitstring:"sizeLB:0,sizeUB:150"`
	PortNumber aper.OctetString      `octetstring:"sizeLB:2,sizeUB:2"`
	Tai        Tai                   `bitstring:"sizeLB:0,sizeUB:150"`
}

type TngfUserLocationInformation struct {
	TnapId     aper.OctetString      `octetstring:"sizeLB:0,sizeUB:150"`
	IpAddress  TransportLayerAddress `bitstring:"sizeLB:0,sizeUB:150"`
	PortNumber aper.OctetString      `octetstring:"sizeLB:2,sizeUB:2"`
	Tai        Tai                   `bitstring:"sizeLB:0,sizeUB:150"`
}

type TwifUserLocationInformation struct {
	TwapId     aper.OctetString      `octetstring:"sizeLB:0,sizeUB:150"`
	IpAddress  TransportLayerAddress `bitstring:"sizeLB:0,sizeUB:150"`
	PortNumber aper.OctetString      `octetstring:"sizeLB:2,sizeUB:2"`
	Tai        Tai                   `bitstring:"sizeLB:0,sizeUB:150"`
}
