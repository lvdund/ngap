package ie

import "ngap/aper"

type BroadcastSessionSetupResponse struct {
	MessageType                     MessageType            //`bitstring:"sizeLB:0,sizeUB:150"`
	MbsSessionId                    MbsSessionId           //`bitstring:"sizeLB:0,sizeUB:150"`
	MbsSessionSetupResponseTransfer aper.OctetString       //`octetstring:"sizeLB:0,sizeUB:150"`
	CriticalityDiagnostics          CriticalityDiagnostics //`bitstring:"sizeLB:0,sizeUB:150"`
}
