package ie

type BroadcastSessionSetupRequest struct {
	MessageType                    *MessageType
	MbsSessionId                   *MbsSessionId
	SNssai                         *SNssai
	MbsServiceArea                 *MbsServiceArea
	MbsSessionSetupRequestTransfer *[]byte
}
