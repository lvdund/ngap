package ie

type DataForwardingAccepted struct {
	DataForwardingAccepted []byte `bitstring:"sizeLB:0,sizeUB:150"`
}
