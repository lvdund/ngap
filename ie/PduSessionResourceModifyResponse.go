package ie

type PduSessionResourceModifyResponse struct {
MessageType	*MessageType
AmfUeNgapId	*AmfUeNgapId
RanUeNgapId	*RanUeNgapId
PduSessionResourceModifyResponseList	*[]PduSessionResourceModifyResponseItem
PduSessionResourceFailedToModifyList	*[]PduSessionResourceFailedToModifyItem
UserLocationInformation	*UserLocationInformation
CriticalityDiagnostics	*CriticalityDiagnostics
}

type PduSessionResourceModifyResponseItem struct {
PduSessionId	*PduSessionId
PduSessionResourceModifyResponseTransfer	*[]byte
}

type PduSessionResourceFailedToModifyItem struct {
PduSessionId	*PduSessionId
PduSessionResourceModifyUnsuccessfulTransfer	*[]byte
}
