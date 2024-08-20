package ie

type TraceFailureIndication struct {
	MessageType  *MessageType
	AmfUeNgapId  *AmfUeNgapId
	RanUeNgapId  *RanUeNgapId
	NgRanTraceId []byte //`bitstring:"sizeLB:8,sizeUB:8"`
	Cause        *Cause
}
