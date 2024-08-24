package ie

type ExtendedAmfName struct {
	AmfNameVisible []byte //`bitstring:"sizeLB:1,sizeUB:150"`
	AmfNameUtf8    []byte //`bitstring:"sizeLB:1,sizeUB:150"`
}
