package ie

type HandoverNotify struct {
	MessageType             MessageType             `bitstring:"sizeLB:0,sizeUB:150"`
	AmfUeNgapId             AmfUeNgapId             `bitstring:"sizeLB:0,sizeUB:150"`
	RanUeNgapId             RanUeNgapId             `bitstring:"sizeLB:0,sizeUB:150"`
	UserLocationInformation UserLocationInformation `bitstring:"sizeLB:0,sizeUB:150"`
	NotifySourceNgRanNode   []byte                  `bitstring:"sizeLB:0,sizeUB:150"`
}
