package ie

type Cause struct {
	ChoiceCauseGroup *ChoiceCauseGroup
}

type ChoiceCauseGroup struct {
	RadioNetworkLayer *RadioNetworkLayer
	TransportLayer    *TransportLayer
	Nas               *Nas
	Protocol          *Protocol
	Miscellaneous     *Miscellaneous
}

type RadioNetworkLayer struct {
	RadioNetworkLayerCause *[]byte
}

type TransportLayer struct {
	TransportLayerCause *[]byte
}

type Nas struct {
	NasCause *[]byte
}

type Protocol struct {
	ProtocolCause *[]byte
}

type Miscellaneous struct {
	MiscellaneousCause *[]byte
}
