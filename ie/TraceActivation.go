package ie

type TraceActivation struct {
	NgRanTraceId                   []byte //`bitstring:"sizeLB:8,sizeUB:8"`
	InterfacesToTrace              []byte //`bitstring:"sizeLB:8,sizeUB:8"`
	TraceDepth                     *[]byte
	TraceCollectionEntityIpAddress *TransportLayerAddress
	MdtConfiguration               *MdtConfiguration
	TraceCollectionEntityUri       *Uri
}
