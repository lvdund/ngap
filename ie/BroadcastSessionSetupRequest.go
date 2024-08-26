package ie

import "ngap/aper"

type BroadcastSessionSetupRequest struct {
	MessageType                    MessageType      `bitstring:"sizeLB:0,sizeUB:150"`
	MbsSessionId                   MbsSessionId     `bitstring:"sizeLB:0,sizeUB:150"`
	SNssai                         SNssai           `bitstring:"sizeLB:0,sizeUB:150"`
	MbsServiceArea                 MbsServiceArea   `bitstring:"sizeLB:0,sizeUB:150"`
	MbsSessionSetupRequestTransfer aper.OctetString `octetstring:"sizeLB:0,sizeUB:150"`
}
