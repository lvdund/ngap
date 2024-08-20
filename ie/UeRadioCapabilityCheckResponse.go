package ie

type UeRadioCapabilityCheckResponse struct {
	MessageType              *MessageType
	AmfUeNgapId              *AmfUeNgapId
	RanUeNgapId              *RanUeNgapId
	ImsVoiceSupportIndicator *ImsVoiceSupportIndicator
	CriticalityDiagnostics   *CriticalityDiagnostics
}
