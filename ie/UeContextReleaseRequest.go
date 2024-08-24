package ie

type UeContextReleaseRequest struct {
MessageType	MessageType	//`bitstring:"sizeLB:0,sizeUB:150"`
AmfUeNgapId	AmfUeNgapId	//`bitstring:"sizeLB:0,sizeUB:150"`
RanUeNgapId	RanUeNgapId	//`bitstring:"sizeLB:0,sizeUB:150"`
PduSessionResourceList	[]PduSessionResourceItem	//`bitstring:"sizeLB:0,sizeUB:150"`
Cause	Cause	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type PduSessionResourceItem struct {
PduSessionId	PduSessionId	//`bitstring:"sizeLB:0,sizeUB:150"`
}
