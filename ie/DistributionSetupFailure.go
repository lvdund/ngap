package ie

import "ngap/aper"

type DistributionSetupFailure struct {
	MessageType                              MessageType            `bitstring:"sizeLB:0,sizeUB:150"`
	MbsSessionId                             MbsSessionId           `bitstring:"sizeLB:0,sizeUB:150"`
	MbsAreaSessionId                         MbsAreaSessionId       `bitstring:"sizeLB:0,sizeUB:150"`
	MbsDistributionSetupUnsuccessfulTransfer aper.OctetString       `octetstring:"sizeLB:0,sizeUB:150"`
	Cause                                    Cause                  `bitstring:"sizeLB:0,sizeUB:150"`
	CriticalityDiagnostics                   CriticalityDiagnostics `bitstring:"sizeLB:0,sizeUB:150"`
}
