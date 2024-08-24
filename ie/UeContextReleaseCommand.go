package ie

type UeContextReleaseCommand struct {
	MessageType     MessageType     //`bitstring:"sizeLB:0,sizeUB:150"`
	ChoiceUeNgapIds ChoiceUeNgapIds //`bitstring:"sizeLB:0,sizeUB:150"`
	Cause           Cause           //`bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceUeNgapIds struct {
	UeNgapIdPair UeNgapIdPair //`bitstring:"sizeLB:0,sizeUB:150"`
	AmfUeNgapId  AmfUeNgapId  //`bitstring:"sizeLB:0,sizeUB:150"`
}

type UeNgapIdPair struct {
	AmfUeNgapId AmfUeNgapId //`bitstring:"sizeLB:0,sizeUB:150"`
	RanUeNgapId RanUeNgapId //`bitstring:"sizeLB:0,sizeUB:150"`
}
