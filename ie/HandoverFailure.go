package ie

type HandoverFailure struct {
	MessageType                               *MessageType
	AmfUeNgapId                               *AmfUeNgapId
	Cause                                     *Cause
	CriticalityDiagnostics                    *CriticalityDiagnostics
	TargetToSourceFailureTransparentContainer *TargetToSourceFailureTransparentContainer
}
