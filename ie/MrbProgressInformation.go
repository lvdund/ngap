package ie

type MrbProgressInformation struct {
	ChoicePdcpSnStatus ChoicePdcpSnStatus //`bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoicePdcpSnStatus struct {
	Ie12Bits Ie12Bits //`bitstring:"sizeLB:0,sizeUB:150"`
	Ie18Bits Ie18Bits //`bitstring:"sizeLB:0,sizeUB:150"`
}
