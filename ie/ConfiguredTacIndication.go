package ie

type ConfiguredTacIndication struct {
	ConfiguredTacIndication []byte `bitstring:"sizeLB:0,sizeUB:150"`
}
