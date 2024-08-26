package ie

import "ngap/aper"

type BroadcastSessionModificationRequest struct {
	MessageType                           MessageType      `bitstring:"sizeLB:0,sizeUB:150"`
	MbsSessionId                          MbsSessionId     `bitstring:"sizeLB:0,sizeUB:150"`
	MbsServiceArea                        MbsServiceArea   `bitstring:"sizeLB:0,sizeUB:150"`
	MbsSessionModificationRequestTransfer aper.OctetString `octetstring:"sizeLB:0,sizeUB:150"`
}
