package ie

type AmfConfigurationUpdateAcknowledge struct {
	MessageType                        *MessageType
	AmfTnlAssociationSetupList         *[]AmfTnlAssociationSetupItem
	AmfTnlAssociationFailedToSetupList *TnlAssociationList
	CriticalityDiagnostics             *CriticalityDiagnostics
}

type AmfTnlAssociationSetupItem struct {
	AmfTnlAssociationAddress *CpTransportLayerInformation
}
