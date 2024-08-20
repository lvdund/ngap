package ie

type Dynamic5QiDescriptor struct {
	PriorityLevel               *PriorityLevel
	PacketDelayBudget           *PacketDelayBudget
	PacketErrorRate             *PacketErrorRate
	Ie5Qi                       uint8 //`bitstring:"sizeLB:0,sizeUB:255"`
	DelayCritical               *[]byte
	AveragingWindow             *AveragingWindow
	MaximumDataBurstVolume      *MaximumDataBurstVolume
	ExtendedPacketDelayBudget   *ExtendedPacketDelayBudget
	CnPacketDelayBudgetDownlink *ExtendedPacketDelayBudget
	CnPacketDelayBudgetUplink   *ExtendedPacketDelayBudget
}
