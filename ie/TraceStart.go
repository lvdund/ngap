package ie

type TraceStart struct {
	MessageType     MessageType     `bitstring:"sizeLB:0,sizeUB:150"`
	AmfUeNgapId     AmfUeNgapId     `bitstring:"sizeLB:0,sizeUB:150"`
	RanUeNgapId     RanUeNgapId     `bitstring:"sizeLB:0,sizeUB:150"`
	TraceActivation TraceActivation `bitstring:"sizeLB:0,sizeUB:150"`
}
