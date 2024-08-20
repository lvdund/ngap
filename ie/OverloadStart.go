package ie

type OverloadStart struct {
	MessageType                       *MessageType
	AmfOverloadResponse               *OverloadResponse
	AmfTrafficLoadReductionIndication *TrafficLoadReductionIndication
	OverloadStartNssaiList            *[]OverloadStartNssaiItem
}

type OverloadStartNssaiItem struct {
	SliceOverloadList                   *SliceOverloadList
	SliceOverloadResponse               *OverloadResponse
	SliceTrafficLoadReductionIndication *TrafficLoadReductionIndication
}
