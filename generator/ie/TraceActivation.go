package ie

type TraceActivation struct {
	NGRANTraceID                   NGRANTraceID          `False,`
	InterfacesToTrace              InterfacesToTrace     `False,`
	TraceDepth                     TraceDepth            `False,`
	TraceCollectionEntityIPAddress TransportLayerAddress `False,`
	IEExtensions                   TraceActivationExtIEs `False,OPTIONAL`
}
