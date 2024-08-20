package ie

type PduSessionResourceReleaseResponse struct {
MessageType	*MessageType
AmfUeNgapId	*AmfUeNgapId
RanUeNgapId	*RanUeNgapId
PduSessionResourceReleasedList	*[]PduSessionResourceReleasedItem
UserLocationInformation	*UserLocationInformation
CriticalityDiagnostics	*CriticalityDiagnostics
}
