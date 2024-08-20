package ie

type PwsCancelResponse struct {
	MessageType                *MessageType
	MessageIdentifier          *MessageIdentifier
	SerialNumber               *SerialNumber
	BroadcastCancelledAreaList *BroadcastCancelledAreaList
	CriticalityDiagnostics     *CriticalityDiagnostics
}
