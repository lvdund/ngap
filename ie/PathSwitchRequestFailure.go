package ie

type PathSwitchRequestFailure struct {
MessageType	*MessageType
AmfUeNgapId	*AmfUeNgapId
RanUeNgapId	*RanUeNgapId
PduSessionResourceReleasedList	*[]PduSessionResourceReleasedItem
CriticalityDiagnostics	*CriticalityDiagnostics
}

type PduSessionResourceReleasedItem struct {
PduSessionId	*PduSessionId
PathSwitchRequestUnsuccessfulTransfer	*[]byte
}
