package ie

type MulticastSessionDeactivationRequest struct {
	MessageType                                 *MessageType
	MbsSessionId                                *MbsSessionId
	MulticastSessionDeactivationRequestTransfer *[]byte
}
