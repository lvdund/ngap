package ie

type EUtraPagingEdrxInformation struct {
EUtraPagingEdrxCycle	[]byte	//`bitstring:"sizeLB:0,sizeUB:150"`
EUtraPagingTimeWindow	[]byte	//`bitstring:"sizeLB:0,sizeUB:150"`
}
