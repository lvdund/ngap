package ie

type HandoverRequired struct {
	MessageType                        MessageType                        `bitstring:"sizeLB:0,sizeUB:150"`
	AmfUeNgapId                        AmfUeNgapId                        `bitstring:"sizeLB:0,sizeUB:150"`
	RanUeNgapId                        RanUeNgapId                        `bitstring:"sizeLB:0,sizeUB:150"`
	HandoverType                       HandoverType                       `bitstring:"sizeLB:0,sizeUB:150"`
	Cause                              Cause                              `bitstring:"sizeLB:0,sizeUB:150"`
	TargetId                           TargetId                           `bitstring:"sizeLB:0,sizeUB:150"`
	DirectForwardingPathAvailability   DirectForwardingPathAvailability   `bitstring:"sizeLB:0,sizeUB:150"`
	PduSessionResourceList             []PduSessionResourceItem           `bitstring:"sizeLB:0,sizeUB:150"`
	SourceToTargetTransparentContainer SourceToTargetTransparentContainer `bitstring:"sizeLB:0,sizeUB:150"`
}
