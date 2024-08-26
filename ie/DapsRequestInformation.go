package ie

type DapsRequestInformation struct {
	DapsIndicator []byte `bitstring:"sizeLB:0,sizeUB:150"`
}
