package ie

type BroadcastSessionModificationRequest struct {
MessageType	*MessageType
MbsSessionId	*MbsSessionId
MbsServiceArea	*MbsServiceArea
MbsSessionModificationRequestTransfer	*[]byte
}
