package ie

type ExtendedRanNodeName struct {
	RanNodeNameVisible []byte `bitstring:"sizeLB:1,sizeUB:150"`
	RanNodeNameUtf8    []byte `bitstring:"sizeLB:1,sizeUB:150"`
}
