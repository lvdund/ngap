package ie

type OverloadStart struct {
MessageType	MessageType	//`bitstring:"sizeLB:0,sizeUB:150"`
AmfOverloadResponse	OverloadResponse	//`bitstring:"sizeLB:0,sizeUB:150"`
AmfTrafficLoadReductionIndication	TrafficLoadReductionIndication	//`bitstring:"sizeLB:0,sizeUB:150"`
OverloadStartNssaiList	[]OverloadStartNssaiItem	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type OverloadStartNssaiItem struct {
SliceOverloadList	SliceOverloadList	//`bitstring:"sizeLB:0,sizeUB:150"`
SliceOverloadResponse	OverloadResponse	//`bitstring:"sizeLB:0,sizeUB:150"`
SliceTrafficLoadReductionIndication	TrafficLoadReductionIndication	//`bitstring:"sizeLB:0,sizeUB:150"`
}
