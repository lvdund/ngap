package ie

import "ngap/aper"

type PduSessionResourceModifyIndication struct {
	MessageType                            MessageType                              `bitstring:"sizeLB:0,sizeUB:150"`
	AmfUeNgapId                            AmfUeNgapId                              `bitstring:"sizeLB:0,sizeUB:150"`
	RanUeNgapId                            RanUeNgapId                              `bitstring:"sizeLB:0,sizeUB:150"`
	PduSessionResourceModifyIndicationList []PduSessionResourceModifyIndicationItem `bitstring:"sizeLB:0,sizeUB:150"`
	UserLocationInformation                UserLocationInformation                  `bitstring:"sizeLB:0,sizeUB:150"`
}

type PduSessionResourceModifyIndicationItem struct {
	PduSessionId                               PduSessionId     `bitstring:"sizeLB:0,sizeUB:150"`
	PduSessionResourceModifyIndicationTransfer aper.OctetString `octetstring:"sizeLB:0,sizeUB:150"`
}
