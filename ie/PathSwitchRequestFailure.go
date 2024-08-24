package ie

import "ngap/aper"

type PathSwitchRequestFailure struct {
	MessageType                    MessageType                      //`bitstring:"sizeLB:0,sizeUB:150"`
	AmfUeNgapId                    AmfUeNgapId                      //`bitstring:"sizeLB:0,sizeUB:150"`
	RanUeNgapId                    RanUeNgapId                      //`bitstring:"sizeLB:0,sizeUB:150"`
	PduSessionResourceReleasedList []PduSessionResourceReleasedItem //`bitstring:"sizeLB:0,sizeUB:150"`
	CriticalityDiagnostics         CriticalityDiagnostics           //`bitstring:"sizeLB:0,sizeUB:150"`
}

type PduSessionResourceReleasedItem struct {
	PduSessionId                          PduSessionId     //`bitstring:"sizeLB:0,sizeUB:150"`
	PathSwitchRequestUnsuccessfulTransfer aper.OctetString //`octetstring:"sizeLB:0,sizeUB:150"`
}
