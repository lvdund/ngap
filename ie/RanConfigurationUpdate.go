package ie

type RanConfigurationUpdate struct {
	MessageType                     *MessageType
	RanNodeName                     []byte //`bitstring:"sizeLB:1,sizeUB:150"`
	SupportedTaList                 *[]SupportedTaItem
	DefaultPagingDrx                *PagingDrx
	GlobalRanNodeId                 *GlobalRanNodeId
	NgRanTnlAssociationToRemoveList *[]NgRanTnlAssociationToRemoveItem
	NbIotDefaultPagingDrx           *NbIotDefaultPagingDrx
	ExtendedRanNodeName             *ExtendedRanNodeName
}

type NgRanTnlAssociationToRemoveItem struct {
	TnlAssociationTransportLayerAddress      *CpTransportLayerInformation
	TnlAssociationTransportLayerAddressAtAmf *CpTransportLayerInformation
}
