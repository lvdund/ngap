package ie

type UplinkNonUeAssociatedNrppaTransport struct {
	MessageType MessageType `bitstring:"sizeLB:0,sizeUB:150"`
	RoutingId   RoutingId   `bitstring:"sizeLB:0,sizeUB:150"`
	NrppaPdu    NrppaPdu    `bitstring:"sizeLB:0,sizeUB:150"`
}
