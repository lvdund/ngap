package ie

type UeContextModificationResponse struct {
	MessageType             *MessageType
	AmfUeNgapId             *AmfUeNgapId
	RanUeNgapId             *RanUeNgapId
	RrcState                *RrcState
	UserLocationInformation *UserLocationInformation
	CriticalityDiagnostics  *CriticalityDiagnostics
}
