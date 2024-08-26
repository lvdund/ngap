package ie

import "ngap/aper"

type MulticastSessionActivationRequest struct {
	MessageType                               MessageType      `bitstring:"sizeLB:0,sizeUB:150"`
	MbsSessionId                              MbsSessionId     `bitstring:"sizeLB:0,sizeUB:150"`
	MulticastSessionActivationRequestTransfer aper.OctetString `octetstring:"sizeLB:0,sizeUB:150"`
}
