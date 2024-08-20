package ie

type ExcessPacketDelayThresholdConfiguration struct {
ExcessPacketDelayThresholdItem	*ExcessPacketDelayThresholdItem
}

type ExcessPacketDelayThresholdItem struct {
Ie5Qi	uint8	//`bitstring:"sizeLB:0,sizeUB:255"`
ExcessPacketDelayThresholdValue	*[]byte
}
