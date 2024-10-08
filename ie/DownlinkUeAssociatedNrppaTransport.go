package ie

type DownlinkUeAssociatedNrppaTransport struct {
	MessageType MessageType `bitstring:"sizeLB:0,sizeUB:150"`
	AmfUeNgapId AmfUeNgapId `bitstring:"sizeLB:0,sizeUB:150"`
	RanUeNgapId RanUeNgapId `bitstring:"sizeLB:0,sizeUB:150"`
	RoutingId   RoutingId   `bitstring:"sizeLB:0,sizeUB:150"`
	NrppaPdu    NrppaPdu    `bitstring:"sizeLB:0,sizeUB:150"`
}
