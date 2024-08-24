package ie

type Cause struct {
	ChoiceCauseGroup ChoiceCauseGroup //`bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceCauseGroup struct {
	RadioNetworkLayer RadioNetworkLayer //`bitstring:"sizeLB:0,sizeUB:150"`
	TransportLayer    TransportLayer    //`bitstring:"sizeLB:0,sizeUB:150"`
	Nas               Nas               //`bitstring:"sizeLB:0,sizeUB:150"`
	Protocol          Protocol          //`bitstring:"sizeLB:0,sizeUB:150"`
	Miscellaneous     Miscellaneous     //`bitstring:"sizeLB:0,sizeUB:150"`
}

type RadioNetworkLayer struct {
	RadioNetworkLayerCause []byte //`bitstring:"sizeLB:0,sizeUB:150"`
}

type TransportLayer struct {
	TransportLayerCause []byte //`bitstring:"sizeLB:0,sizeUB:150"`
}

type Nas struct {
	NasCause []byte //`bitstring:"sizeLB:0,sizeUB:150"`
}

type Protocol struct {
	ProtocolCause []byte //`bitstring:"sizeLB:0,sizeUB:150"`
}

type Miscellaneous struct {
	MiscellaneousCause []byte //`bitstring:"sizeLB:0,sizeUB:150"`
}
