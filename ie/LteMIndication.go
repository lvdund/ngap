package ie

type LteMIndication struct {
	LteMIndication []byte `bitstring:"sizeLB:0,sizeUB:150"`
}
