package ie

import "ngap/aper"

type CellTrafficTrace struct {
MessageType	MessageType	//`bitstring:"sizeLB:0,sizeUB:150"`
AmfUeNgapId	AmfUeNgapId	//`bitstring:"sizeLB:0,sizeUB:150"`
RanUeNgapId	RanUeNgapId	//`bitstring:"sizeLB:0,sizeUB:150"`
NgRanTraceId	aper.OctetString	//`octetstring:"sizeLB:8,sizeUB:8"`
NgRanCgi	NgRanCgi	//`bitstring:"sizeLB:0,sizeUB:150"`
TraceCollectionEntityIpAddress	TransportLayerAddress	//`bitstring:"sizeLB:0,sizeUB:150"`
PrivacyIndicator	[]byte	//`bitstring:"sizeLB:0,sizeUB:150"`
TraceCollectionEntityUri	Uri	//`bitstring:"sizeLB:0,sizeUB:150"`
}
