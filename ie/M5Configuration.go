package ie

type M5Configuration struct {
	M5CollectionPeriod []byte //`bitstring:"sizeLB:0,sizeUB:150"`
	M5LinksToLog       []byte //`bitstring:"sizeLB:0,sizeUB:150"`
	M5ReportAmount     []byte //`bitstring:"sizeLB:0,sizeUB:150"`
}
