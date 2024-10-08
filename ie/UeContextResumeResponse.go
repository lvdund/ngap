package ie

import "ngap/aper"

type UeContextResumeResponse struct {
	MessageType                          MessageType                            `bitstring:"sizeLB:0,sizeUB:150"`
	AmfUeNgapId                          AmfUeNgapId                            `bitstring:"sizeLB:0,sizeUB:150"`
	RanUeNgapId                          RanUeNgapId                            `bitstring:"sizeLB:0,sizeUB:150"`
	PduSessionResourceResumeList         []PduSessionResourceResumeItem         `bitstring:"sizeLB:0,sizeUB:150"`
	PduSessionResourceFailedToResumeList []PduSessionResourceFailedToResumeItem `bitstring:"sizeLB:0,sizeUB:150"`
	SecurityContext                      SecurityContext                        `bitstring:"sizeLB:0,sizeUB:150"`
	SuspendResponseIndication            SuspendResponseIndication              `bitstring:"sizeLB:0,sizeUB:150"`
	ExtendedConnectedTime                ExtendedConnectedTime                  `bitstring:"sizeLB:0,sizeUB:150"`
	CriticalityDiagnostics               CriticalityDiagnostics                 `bitstring:"sizeLB:0,sizeUB:150"`
}

type PduSessionResourceResumeItem struct {
	PduSessionId                    PduSessionId     `bitstring:"sizeLB:0,sizeUB:150"`
	UeContextResumeResponseTransfer aper.OctetString `octetstring:"sizeLB:0,sizeUB:150"`
}

type PduSessionResourceFailedToResumeItem struct {
	PduSessionId PduSessionId `bitstring:"sizeLB:0,sizeUB:150"`
	Cause        Cause        `bitstring:"sizeLB:0,sizeUB:150"`
}
