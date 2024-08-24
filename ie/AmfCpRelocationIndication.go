package ie

type AmfCpRelocationIndication struct {
	MessageType  MessageType  //`bitstring:"sizeLB:0,sizeUB:150"`
	AmfUeNgapId  AmfUeNgapId  //`bitstring:"sizeLB:0,sizeUB:150"`
	RanUeNgapId  RanUeNgapId  //`bitstring:"sizeLB:0,sizeUB:150"`
	SNssai       SNssai       //`bitstring:"sizeLB:0,sizeUB:150"`
	AllowedNssai AllowedNssai //`bitstring:"sizeLB:0,sizeUB:150"`
}
