package ie

type PduSessionResourceModifyIndication struct {
	MessageType                            *MessageType
	AmfUeNgapId                            *AmfUeNgapId
	RanUeNgapId                            *RanUeNgapId
	PduSessionResourceModifyIndicationList *[]PduSessionResourceModifyIndicationItem
	UserLocationInformation                *UserLocationInformation
}

type PduSessionResourceModifyIndicationItem struct {
	PduSessionId                               *PduSessionId
	PduSessionResourceModifyIndicationTransfer *[]byte
}
