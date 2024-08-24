package ie

type AmfConfigurationUpdate struct {
	MessageType                   MessageType                     //`bitstring:"sizeLB:0,sizeUB:150"`
	AmfName                       AmfName                         //`bitstring:"sizeLB:0,sizeUB:150"`
	ServedGuamiList               []ServedGuamiItem               //`bitstring:"sizeLB:0,sizeUB:150"`
	RelativeAmfCapacity           RelativeAmfCapacity             //`bitstring:"sizeLB:0,sizeUB:150"`
	PlmnSupportList               []PlmnSupportItem               //`bitstring:"sizeLB:0,sizeUB:150"`
	AmfTnlAssociationToAddList    []AmfTnlAssociationToAddItem    //`bitstring:"sizeLB:0,sizeUB:150"`
	AmfTnlAssociationToRemoveList []AmfTnlAssociationToRemoveItem //`bitstring:"sizeLB:0,sizeUB:150"`
	AmfTnlAssociationToUpdateList []AmfTnlAssociationToUpdateItem //`bitstring:"sizeLB:0,sizeUB:150"`
	ExtendedAmfName               ExtendedAmfName                 //`bitstring:"sizeLB:0,sizeUB:150"`
}

type ServedGuamiItem struct {
	Guami         Guami   //`bitstring:"sizeLB:0,sizeUB:150"`
	BackupAmfName AmfName //`bitstring:"sizeLB:0,sizeUB:150"`
	GuamiType     []byte  //`bitstring:"sizeLB:0,sizeUB:150"`
}

type PlmnSupportItem struct {
	PlmnIdentity             PlmnIdentity             //`bitstring:"sizeLB:0,sizeUB:150"`
	SliceSupportList         SliceSupportList         //`bitstring:"sizeLB:0,sizeUB:150"`
	NpnSupport               NpnSupport               //`bitstring:"sizeLB:0,sizeUB:150"`
	ExtendedSliceSupportList ExtendedSliceSupportList //`bitstring:"sizeLB:0,sizeUB:150"`
	OnboardingSupport        []byte                   //`bitstring:"sizeLB:0,sizeUB:150"`
}

type AmfTnlAssociationToAddItem struct {
	AmfTnlAssociationAddress CpTransportLayerInformation //`bitstring:"sizeLB:0,sizeUB:150"`
	TnlAssociationUsage      TnlAssociationUsage         //`bitstring:"sizeLB:0,sizeUB:150"`
	TnlAddressWeightFactor   TnlAddressWeightFactor      //`bitstring:"sizeLB:0,sizeUB:150"`
}

type AmfTnlAssociationToRemoveItem struct {
	AmfTnlAssociationAddress                 CpTransportLayerInformation //`bitstring:"sizeLB:0,sizeUB:150"`
	TnlAssociationTransportLayerAddressNgRan CpTransportLayerInformation //`bitstring:"sizeLB:0,sizeUB:150"`
}

type AmfTnlAssociationToUpdateItem struct {
	AmfTnlAssociationAddress CpTransportLayerInformation //`bitstring:"sizeLB:0,sizeUB:150"`
	TnlAssociationUsage      TnlAssociationUsage         //`bitstring:"sizeLB:0,sizeUB:150"`
	TnlAddressWeightFactor   TnlAddressWeightFactor      //`bitstring:"sizeLB:0,sizeUB:150"`
}
