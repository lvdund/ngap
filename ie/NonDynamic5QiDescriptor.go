package ie

import "ngap/aper"

type NonDynamic5QiDescriptor struct {
Ie5Qi	aper.Integer	//`Integer:"valueLB:0,valueUB:255"`
PriorityLevel	PriorityLevel	//`bitstring:"sizeLB:0,sizeUB:150"`
AveragingWindow	AveragingWindow	//`bitstring:"sizeLB:0,sizeUB:150"`
MaximumDataBurstVolume	MaximumDataBurstVolume	//`bitstring:"sizeLB:0,sizeUB:150"`
CnPacketDelayBudgetDownlink	ExtendedPacketDelayBudget	//`bitstring:"sizeLB:0,sizeUB:150"`
CnPacketDelayBudgetUplink	ExtendedPacketDelayBudget	//`bitstring:"sizeLB:0,sizeUB:150"`
}
