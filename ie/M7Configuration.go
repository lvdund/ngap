package ie

type M7Configuration struct {
M7CollectionPeriod	uint8	//`bitstring:"sizeLB:1,sizeUB:60"`
M7LinksToLog	*[]byte
M7ReportAmount	*[]byte
}
