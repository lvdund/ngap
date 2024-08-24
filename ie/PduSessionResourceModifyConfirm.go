package ie

import "ngap/aper"

type PduSessionResourceModifyConfirm struct {
	MessageType                          MessageType                            //`bitstring:"sizeLB:0,sizeUB:150"`
	AmfUeNgapId                          AmfUeNgapId                            //`bitstring:"sizeLB:0,sizeUB:150"`
	RanUeNgapId                          RanUeNgapId                            //`bitstring:"sizeLB:0,sizeUB:150"`
	PduSessionResourceModifyConfirmList  []PduSessionResourceModifyConfirmItem  //`bitstring:"sizeLB:0,sizeUB:150"`
	PduSessionResourceFailedToModifyList []PduSessionResourceFailedToModifyItem //`bitstring:"sizeLB:0,sizeUB:150"`
	CriticalityDiagnostics               CriticalityDiagnostics                 //`bitstring:"sizeLB:0,sizeUB:150"`
}

type PduSessionResourceModifyConfirmItem struct {
	PduSessionId                            PduSessionId     //`bitstring:"sizeLB:0,sizeUB:150"`
	PduSessionResourceModifyConfirmTransfer aper.OctetString //`octetstring:"sizeLB:0,sizeUB:150"`
}
