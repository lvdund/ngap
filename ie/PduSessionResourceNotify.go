package ie

type PduSessionResourceNotify struct {
	MessageType                    *MessageType
	AmfUeNgapId                    *AmfUeNgapId
	RanUeNgapId                    *RanUeNgapId
	PduSessionResourceNotifyList   *[]PduSessionResourceNotifyItem
	PduSessionResourceReleasedList *[]PduSessionResourceReleasedItem
	UserLocationInformation        *UserLocationInformation
}

type PduSessionResourceNotifyItem struct {
	PduSessionId                     *PduSessionId
	PduSessionResourceNotifyTransfer *[]byte
}
