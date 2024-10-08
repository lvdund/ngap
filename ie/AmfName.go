package ie

type AmfName struct {
	AmfName []byte `bitstring:"sizeLB:1,sizeUB:150"`
}
