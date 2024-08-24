package ie

type AmfConfigurationUpdateAcknowledge struct {
MessageType	MessageType	//`bitstring:"sizeLB:0,sizeUB:150"`
AmfTnlAssociationSetupList	[]AmfTnlAssociationSetupItem	//`bitstring:"sizeLB:0,sizeUB:150"`
AmfTnlAssociationFailedToSetupList	TnlAssociationList	//`bitstring:"sizeLB:0,sizeUB:150"`
CriticalityDiagnostics	CriticalityDiagnostics	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type AmfTnlAssociationSetupItem struct {
AmfTnlAssociationAddress	CpTransportLayerInformation	//`bitstring:"sizeLB:0,sizeUB:150"`
}
