package ie

type UeContextResumeResponse struct {
	MessageType                          *MessageType
	AmfUeNgapId                          *AmfUeNgapId
	RanUeNgapId                          *RanUeNgapId
	PduSessionResourceResumeList         *[]PduSessionResourceResumeItem
	PduSessionResourceFailedToResumeList *[]PduSessionResourceFailedToResumeItem
	SecurityContext                      *SecurityContext
	SuspendResponseIndication            *SuspendResponseIndication
	ExtendedConnectedTime                *ExtendedConnectedTime
	CriticalityDiagnostics               *CriticalityDiagnostics
}

type PduSessionResourceResumeItem struct {
	PduSessionId                    *PduSessionId
	UeContextResumeResponseTransfer *[]byte
}

type PduSessionResourceFailedToResumeItem struct {
	PduSessionId *PduSessionId
	Cause        *Cause
}
