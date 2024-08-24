package ie

type TnlAssociationList struct {
TnlAssociationItem	TnlAssociationItem	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type TnlAssociationItem struct {
TnlAssociationAddress	CpTransportLayerInformation	//`bitstring:"sizeLB:0,sizeUB:150"`
Cause	Cause	//`bitstring:"sizeLB:0,sizeUB:150"`
}
