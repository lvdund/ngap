package ie

type RrcInactiveTransitionReport struct {
MessageType	MessageType	//`bitstring:"sizeLB:0,sizeUB:150"`
AmfUeNgapId	AmfUeNgapId	//`bitstring:"sizeLB:0,sizeUB:150"`
RanUeNgapId	RanUeNgapId	//`bitstring:"sizeLB:0,sizeUB:150"`
RrcState	RrcState	//`bitstring:"sizeLB:0,sizeUB:150"`
UserLocationInformation	UserLocationInformation	//`bitstring:"sizeLB:0,sizeUB:150"`
}
