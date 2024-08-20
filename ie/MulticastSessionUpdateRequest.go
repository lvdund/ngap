package ie

type MulticastSessionUpdateRequest struct {
MessageType	*MessageType
MbsSessionId	*MbsSessionId
MbsAreaSessionId	*MbsAreaSessionId
MulticastSessionUpdateRequestTransfer	*[]byte
}
