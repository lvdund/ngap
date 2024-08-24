package ie

type M4Configuration struct {
	M4CollectionPeriod []byte //`bitstring:"sizeLB:0,sizeUB:150"`
	M4LinksToLog       []byte //`bitstring:"sizeLB:0,sizeUB:150"`
	M4ReportAmount     []byte //`bitstring:"sizeLB:0,sizeUB:150"`
}
