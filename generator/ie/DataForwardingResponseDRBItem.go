package ie

type DataForwardingResponseDRBItem struct {
	DRBID                        DRBID                               `False,`
	DLForwardingUPTNLInformation UPTransportLayerInformation         `False,OPTIONAL`
	ULForwardingUPTNLInformation UPTransportLayerInformation         `False,OPTIONAL`
	IEExtensions                 DataForwardingResponseDRBItemExtIEs `False,OPTIONAL`
}
