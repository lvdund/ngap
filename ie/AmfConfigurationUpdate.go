package ie

type AmfConfigurationUpdate struct {
MessageType	*MessageType
AmfName	*AmfName
ServedGuamiList	*[]ServedGuamiItem
RelativeAmfCapacity	*RelativeAmfCapacity
PlmnSupportList	*[]PlmnSupportItem
AmfTnlAssociationToAddList	*[]AmfTnlAssociationToAddItem
AmfTnlAssociationToRemoveList	*[]AmfTnlAssociationToRemoveItem
AmfTnlAssociationToUpdateList	*[]AmfTnlAssociationToUpdateItem
ExtendedAmfName	*ExtendedAmfName
}

type ServedGuamiItem struct {
Guami	*Guami
BackupAmfName	*AmfName
GuamiType	*[]byte
}

type PlmnSupportItem struct {
PlmnIdentity	*PlmnIdentity
SliceSupportList	*SliceSupportList
NpnSupport	*NpnSupport
ExtendedSliceSupportList	*ExtendedSliceSupportList
OnboardingSupport	*[]byte
}

type AmfTnlAssociationToAddItem struct {
AmfTnlAssociationAddress	*CpTransportLayerInformation
TnlAssociationUsage	*TnlAssociationUsage
TnlAddressWeightFactor	*TnlAddressWeightFactor
}

type AmfTnlAssociationToRemoveItem struct {
AmfTnlAssociationAddress	*CpTransportLayerInformation
TnlAssociationTransportLayerAddressNgRan	*CpTransportLayerInformation
}

type AmfTnlAssociationToUpdateItem struct {
AmfTnlAssociationAddress	*CpTransportLayerInformation
TnlAssociationUsage	*TnlAssociationUsage
TnlAddressWeightFactor	*TnlAddressWeightFactor
}
