package ie

type NgSetupRequest struct {
	MessageType            *MessageType
	GlobalRanNodeId        *GlobalRanNodeId
	RanNodeName            []byte //`bitstring:"sizeLB:1,sizeUB:150"`
	SupportedTaList        *[]SupportedTaItem
	DefaultPagingDrx       *PagingDrx
	UeRetentionInformation *UeRetentionInformation
	NbIotDefaultPagingDrx  *NbIotDefaultPagingDrx
	ExtendedRanNodeName    *ExtendedRanNodeName
}

type SupportedTaItem struct {
	Tac                     *Tac
	BroadcastPlmnList       *[]BroadcastPlmnItem
	ConfiguredTacIndication *ConfiguredTacIndication
	RatInformation          *RatInformation
}

type BroadcastPlmnItem struct {
	PlmnIdentity                *PlmnIdentity
	TaiSliceSupportList         *SliceSupportList
	NpnSupport                  *NpnSupport
	ExtendedTaiSliceSupportList *ExtendedSliceSupportList
	TaiNsagSupportList          *TaiNsagSupportList
}
