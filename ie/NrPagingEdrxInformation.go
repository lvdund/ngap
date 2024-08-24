package ie

type NrPagingEdrxInformation struct {
NrPagingEdrxCycle	[]byte	//`bitstring:"sizeLB:0,sizeUB:150"`
NrPagingTimeWindow	[]byte	//`bitstring:"sizeLB:0,sizeUB:150"`
}
