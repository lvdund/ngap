package ie

type NgSetupRequest struct {
	MessageType            MessageType            //`bitstring:"sizeLB:0,sizeUB:150"`
	GlobalRanNodeId        GlobalRanNodeId        //`bitstring:"sizeLB:0,sizeUB:150"`
	RanNodeName            []byte                 //`bitstring:"sizeLB:1,sizeUB:150"`
	SupportedTaList        []SupportedTaItem      //`bitstring:"sizeLB:0,sizeUB:150"`
	DefaultPagingDrx       PagingDrx              //`bitstring:"sizeLB:0,sizeUB:150"`
	UeRetentionInformation UeRetentionInformation //`bitstring:"sizeLB:0,sizeUB:150"`
	NbIotDefaultPagingDrx  NbIotDefaultPagingDrx  //`bitstring:"sizeLB:0,sizeUB:150"`
	ExtendedRanNodeName    ExtendedRanNodeName    //`bitstring:"sizeLB:0,sizeUB:150"`
}

type SupportedTaItem struct {
	Tac                     Tac                     //`bitstring:"sizeLB:0,sizeUB:150"`
	BroadcastPlmnList       []BroadcastPlmnItem     //`bitstring:"sizeLB:0,sizeUB:150"`
	ConfiguredTacIndication ConfiguredTacIndication //`bitstring:"sizeLB:0,sizeUB:150"`
	RatInformation          RatInformation          //`bitstring:"sizeLB:0,sizeUB:150"`
}

type BroadcastPlmnItem struct {
	PlmnIdentity                PlmnIdentity             //`bitstring:"sizeLB:0,sizeUB:150"`
	TaiSliceSupportList         SliceSupportList         //`bitstring:"sizeLB:0,sizeUB:150"`
	NpnSupport                  NpnSupport               //`bitstring:"sizeLB:0,sizeUB:150"`
	ExtendedTaiSliceSupportList ExtendedSliceSupportList //`bitstring:"sizeLB:0,sizeUB:150"`
	TaiNsagSupportList          TaiNsagSupportList       //`bitstring:"sizeLB:0,sizeUB:150"`
}
