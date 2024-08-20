package ie

type CellTrafficTrace struct {
MessageType	*MessageType
AmfUeNgapId	*AmfUeNgapId
RanUeNgapId	*RanUeNgapId
NgRanTraceId	[]byte	//`bitstring:"sizeLB:8,sizeUB:8"`
NgRanCgi	*NgRanCgi
TraceCollectionEntityIpAddress	*TransportLayerAddress
PrivacyIndicator	*[]byte
TraceCollectionEntityUri	*Uri
}
