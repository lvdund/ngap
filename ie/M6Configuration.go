package ie

type M6Configuration struct {
M6ReportInterval	[]byte	//`bitstring:"sizeLB:0,sizeUB:150"`
M6LinksToLog	[]byte	//`bitstring:"sizeLB:0,sizeUB:150"`
M6ReportAmount	[]byte	//`bitstring:"sizeLB:0,sizeUB:150"`
ExcessPacketDelayThresholdConfiguration	ExcessPacketDelayThresholdConfiguration	//`bitstring:"sizeLB:0,sizeUB:150"`
}
