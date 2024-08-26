package ie

type EndIndication struct {
	EndIndication []byte `bitstring:"sizeLB:0,sizeUB:150"`
}
