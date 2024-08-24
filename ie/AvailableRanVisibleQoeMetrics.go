package ie

type AvailableRanVisibleQoeMetrics struct {
	ApplicationLayerBufferLevelList []byte //`bitstring:"sizeLB:0,sizeUB:150"`
	PlayoutDelayForMediaStartup     []byte //`bitstring:"sizeLB:0,sizeUB:150"`
}
