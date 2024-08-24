package ie

import "ngap/aper"

type InitialContextSetupResponse struct {
MessageType	MessageType	//`bitstring:"sizeLB:0,sizeUB:150"`
AmfUeNgapId	AmfUeNgapId	//`bitstring:"sizeLB:0,sizeUB:150"`
RanUeNgapId	RanUeNgapId	//`bitstring:"sizeLB:0,sizeUB:150"`
PduSessionResourceSetupResponseList	[]PduSessionResourceSetupResponseItem	//`bitstring:"sizeLB:0,sizeUB:150"`
PduSessionResourceFailedToSetupList	[]PduSessionResourceFailedToSetupItem	//`bitstring:"sizeLB:0,sizeUB:150"`
CriticalityDiagnostics	CriticalityDiagnostics	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type PduSessionResourceSetupResponseItem struct {
PduSessionId	PduSessionId	//`bitstring:"sizeLB:0,sizeUB:150"`
PduSessionResourceSetupResponseTransfer	aper.OctetString	//`octetstring:"sizeLB:0,sizeUB:150"`
}
