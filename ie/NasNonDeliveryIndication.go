package ie

type NasNonDeliveryIndication struct {
	MessageType *MessageType
	AmfUeNgapId *AmfUeNgapId
	RanUeNgapId *RanUeNgapId
	NasPdu      *NasPdu
	Cause       *Cause
}
