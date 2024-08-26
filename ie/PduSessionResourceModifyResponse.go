package ie

import "ngap/aper"

type PduSessionResourceModifyResponse struct {
	MessageType                          MessageType                            `bitstring:"sizeLB:0,sizeUB:150"`
	AmfUeNgapId                          AmfUeNgapId                            `bitstring:"sizeLB:0,sizeUB:150"`
	RanUeNgapId                          RanUeNgapId                            `bitstring:"sizeLB:0,sizeUB:150"`
	PduSessionResourceModifyResponseList []PduSessionResourceModifyResponseItem `bitstring:"sizeLB:0,sizeUB:150"`
	PduSessionResourceFailedToModifyList []PduSessionResourceFailedToModifyItem `bitstring:"sizeLB:0,sizeUB:150"`
	UserLocationInformation              UserLocationInformation                `bitstring:"sizeLB:0,sizeUB:150"`
	CriticalityDiagnostics               CriticalityDiagnostics                 `bitstring:"sizeLB:0,sizeUB:150"`
}

type PduSessionResourceModifyResponseItem struct {
	PduSessionId                             PduSessionId     `bitstring:"sizeLB:0,sizeUB:150"`
	PduSessionResourceModifyResponseTransfer aper.OctetString `octetstring:"sizeLB:0,sizeUB:150"`
}

type PduSessionResourceFailedToModifyItem struct {
	PduSessionId                                 PduSessionId     `bitstring:"sizeLB:0,sizeUB:150"`
	PduSessionResourceModifyUnsuccessfulTransfer aper.OctetString `octetstring:"sizeLB:0,sizeUB:150"`
}
