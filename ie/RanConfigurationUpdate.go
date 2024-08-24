package ie

type RanConfigurationUpdate struct {
	MessageType                     MessageType                       //`bitstring:"sizeLB:0,sizeUB:150"`
	RanNodeName                     []byte                            //`bitstring:"sizeLB:1,sizeUB:150"`
	SupportedTaList                 []SupportedTaItem                 //`bitstring:"sizeLB:0,sizeUB:150"`
	DefaultPagingDrx                PagingDrx                         //`bitstring:"sizeLB:0,sizeUB:150"`
	GlobalRanNodeId                 GlobalRanNodeId                   //`bitstring:"sizeLB:0,sizeUB:150"`
	NgRanTnlAssociationToRemoveList []NgRanTnlAssociationToRemoveItem //`bitstring:"sizeLB:0,sizeUB:150"`
	NbIotDefaultPagingDrx           NbIotDefaultPagingDrx             //`bitstring:"sizeLB:0,sizeUB:150"`
	ExtendedRanNodeName             ExtendedRanNodeName               //`bitstring:"sizeLB:0,sizeUB:150"`
}

type NgRanTnlAssociationToRemoveItem struct {
	TnlAssociationTransportLayerAddress      CpTransportLayerInformation //`bitstring:"sizeLB:0,sizeUB:150"`
	TnlAssociationTransportLayerAddressAtAmf CpTransportLayerInformation //`bitstring:"sizeLB:0,sizeUB:150"`
}
