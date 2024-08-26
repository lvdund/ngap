package ie

import "ngap/aper"

type TraceFailureIndication struct {
	MessageType  MessageType      `bitstring:"sizeLB:0,sizeUB:150"`
	AmfUeNgapId  AmfUeNgapId      `bitstring:"sizeLB:0,sizeUB:150"`
	RanUeNgapId  RanUeNgapId      `bitstring:"sizeLB:0,sizeUB:150"`
	NgRanTraceId aper.OctetString `octetstring:"sizeLB:8,sizeUB:8"`
	Cause        Cause            `bitstring:"sizeLB:0,sizeUB:150"`
}
