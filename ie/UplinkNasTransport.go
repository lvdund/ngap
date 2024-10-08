package ie

import "ngap/aper"

type UplinkNasTransport struct {
	MessageType             MessageType             `bitstring:"sizeLB:0,sizeUB:150"`
	AmfUeNgapId             AmfUeNgapId             `bitstring:"sizeLB:0,sizeUB:150"`
	RanUeNgapId             RanUeNgapId             `bitstring:"sizeLB:0,sizeUB:150"`
	NasPdu                  NasPdu                  `bitstring:"sizeLB:0,sizeUB:150"`
	UserLocationInformation UserLocationInformation `bitstring:"sizeLB:0,sizeUB:150"`
	WAgfIdentityInformation aper.OctetString        `octetstring:"sizeLB:0,sizeUB:150"`
	TngfIdentityInformation aper.OctetString        `octetstring:"sizeLB:0,sizeUB:150"`
	TwifIdentityInformation aper.OctetString        `octetstring:"sizeLB:0,sizeUB:150"`
}
