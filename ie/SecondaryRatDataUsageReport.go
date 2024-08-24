package ie

import "ngap/aper"

type SecondaryRatDataUsageReport struct {
MessageType	MessageType	//`bitstring:"sizeLB:0,sizeUB:150"`
AmfUeNgapId	AmfUeNgapId	//`bitstring:"sizeLB:0,sizeUB:150"`
RanUeNgapId	RanUeNgapId	//`bitstring:"sizeLB:0,sizeUB:150"`
PduSessionResourceSecondaryRatUsageList	[]PduSessionResourceSecondaryRatUsageItem	//`bitstring:"sizeLB:0,sizeUB:150"`
HandoverFlag	[]byte	//`bitstring:"sizeLB:0,sizeUB:150"`
UserLocationInformation	UserLocationInformation	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type PduSessionResourceSecondaryRatUsageItem struct {
PduSessionId	PduSessionId	//`bitstring:"sizeLB:0,sizeUB:150"`
SecondaryRatDataUsageReportTransfer	aper.OctetString	//`octetstring:"sizeLB:0,sizeUB:150"`
}
