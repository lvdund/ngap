package ie

import "ngap/aper"

type DistributionReleaseRequest struct {
	MessageType                           MessageType      `bitstring:"sizeLB:0,sizeUB:150"`
	MbsSessionId                          MbsSessionId     `bitstring:"sizeLB:0,sizeUB:150"`
	MbsAreaSessionId                      MbsAreaSessionId `bitstring:"sizeLB:0,sizeUB:150"`
	MbsDistributionReleaseRequestTransfer aper.OctetString `octetstring:"sizeLB:0,sizeUB:150"`
	Cause                                 Cause            `bitstring:"sizeLB:0,sizeUB:150"`
}
