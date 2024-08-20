package ie

type NgSetupResponse struct {
	MessageType            *MessageType
	AmfName                *AmfName
	ServedGuamiList        *[]ServedGuamiItem
	RelativeAmfCapacity    *RelativeAmfCapacity
	PlmnSupportList        *[]PlmnSupportItem
	CriticalityDiagnostics *CriticalityDiagnostics
	UeRetentionInformation *UeRetentionInformation
	IabSupported           *[]byte
	ExtendedAmfName        *ExtendedAmfName
}
