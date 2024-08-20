package ie

type PduSessionResourceModifyRequest struct {
MessageType	*MessageType
AmfUeNgapId	*AmfUeNgapId
RanUeNgapId	*RanUeNgapId
RanPagingPriority	*RanPagingPriority
PduSessionResourceModifyRequestList	*[]PduSessionResourceModifyRequestItem
}

type PduSessionResourceModifyRequestItem struct {
PduSessionId	*PduSessionId
NasPdu	*NasPdu
PduSessionResourceModifyRequestTransfer	*[]byte
SNssai	*SNssai
PduSessionExpectedUeActivityBehaviour	*ExpectedUeActivityBehaviour
}
