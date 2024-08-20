package ie

type PduSessionResourceReleaseCommand struct {
MessageType	*MessageType
AmfUeNgapId	*AmfUeNgapId
RanUeNgapId	*RanUeNgapId
RanPagingPriority	*RanPagingPriority
NasPdu	*NasPdu
PduSessionResourceToReleaseList	*[]PduSessionResourceToReleaseItem
}

type PduSessionResourceToReleaseItem struct {
PduSessionId	*PduSessionId
PduSessionResourceReleaseCommandTransfer	*[]byte
}
