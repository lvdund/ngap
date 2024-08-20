package ie

type NonDynamic5QiDescriptor struct {
Ie5Qi	uint8	//`bitstring:"sizeLB:0,sizeUB:255"`
PriorityLevel	*PriorityLevel
AveragingWindow	*AveragingWindow
MaximumDataBurstVolume	*MaximumDataBurstVolume
CnPacketDelayBudgetDownlink	*ExtendedPacketDelayBudget
CnPacketDelayBudgetUplink	*ExtendedPacketDelayBudget
}
