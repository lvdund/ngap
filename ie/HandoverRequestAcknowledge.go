package ie

import "ngap/aper"

type HandoverRequestAcknowledge struct {
	MessageType                         MessageType                           `bitstring:"sizeLB:0,sizeUB:150"`
	AmfUeNgapId                         AmfUeNgapId                           `bitstring:"sizeLB:0,sizeUB:150"`
	RanUeNgapId                         RanUeNgapId                           `bitstring:"sizeLB:0,sizeUB:150"`
	PduSessionResourceAdmittedList      []PduSessionResourceAdmittedItem      `bitstring:"sizeLB:0,sizeUB:150"`
	PduSessionResourceFailedToSetupList []PduSessionResourceFailedToSetupItem `bitstring:"sizeLB:0,sizeUB:150"`
	TargetToSourceTransparentContainer  TargetToSourceTransparentContainer    `bitstring:"sizeLB:0,sizeUB:150"`
	CriticalityDiagnostics              CriticalityDiagnostics                `bitstring:"sizeLB:0,sizeUB:150"`
	NpnAccessInformation                NpnAccessInformation                  `bitstring:"sizeLB:0,sizeUB:150"`
	RedcapIndication                    RedcapIndication                      `bitstring:"sizeLB:0,sizeUB:150"`
}

type PduSessionResourceAdmittedItem struct {
	PduSessionId                       PduSessionId     `bitstring:"sizeLB:0,sizeUB:150"`
	HandoverRequestAcknowledgeTransfer aper.OctetString `octetstring:"sizeLB:0,sizeUB:150"`
}
