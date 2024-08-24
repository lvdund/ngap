package ie

import "ngap/aper"

type MulticastSessionDeactivationRequest struct {
	MessageType                                 MessageType      //`bitstring:"sizeLB:0,sizeUB:150"`
	MbsSessionId                                MbsSessionId     //`bitstring:"sizeLB:0,sizeUB:150"`
	MulticastSessionDeactivationRequestTransfer aper.OctetString //`octetstring:"sizeLB:0,sizeUB:150"`
}
