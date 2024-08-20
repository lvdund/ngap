package ie

type DownlinkUeAssociatedNrppaTransport struct {
	MessageType *MessageType
	AmfUeNgapId *AmfUeNgapId
	RanUeNgapId *RanUeNgapId
	RoutingId   *RoutingId
	NrppaPdu    *NrppaPdu
}
