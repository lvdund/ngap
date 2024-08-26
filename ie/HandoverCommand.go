package ie

import "ngap/aper"

type HandoverCommand struct {
	MessageType                        MessageType                        `bitstring:"sizeLB:0,sizeUB:150"`
	AmfUeNgapId                        AmfUeNgapId                        `bitstring:"sizeLB:0,sizeUB:150"`
	RanUeNgapId                        RanUeNgapId                        `bitstring:"sizeLB:0,sizeUB:150"`
	HandoverType                       HandoverType                       `bitstring:"sizeLB:0,sizeUB:150"`
	NasSecurityParametersFromNgRan     NasSecurityParametersFromNgRan     `bitstring:"sizeLB:0,sizeUB:150"`
	PduSessionResourceHandoverList     []PduSessionResourceHandoverItem   `bitstring:"sizeLB:0,sizeUB:150"`
	PduSessionResourceToReleaseList    []PduSessionResourceToReleaseItem  `bitstring:"sizeLB:0,sizeUB:150"`
	TargetToSourceTransparentContainer TargetToSourceTransparentContainer `bitstring:"sizeLB:0,sizeUB:150"`
	CriticalityDiagnostics             CriticalityDiagnostics             `bitstring:"sizeLB:0,sizeUB:150"`
}

type PduSessionResourceHandoverItem struct {
	PduSessionId            PduSessionId     `bitstring:"sizeLB:0,sizeUB:150"`
	HandoverCommandTransfer aper.OctetString `octetstring:"sizeLB:0,sizeUB:150"`
}
