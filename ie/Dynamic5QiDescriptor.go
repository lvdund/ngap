package ie

import "ngap/aper"

type Dynamic5QiDescriptor struct {
	PriorityLevel               PriorityLevel             `bitstring:"sizeLB:0,sizeUB:150"`
	PacketDelayBudget           PacketDelayBudget         `bitstring:"sizeLB:0,sizeUB:150"`
	PacketErrorRate             PacketErrorRate           `bitstring:"sizeLB:0,sizeUB:150"`
	Ie5Qi                       aper.Integer              `Integer:"valueLB:0,valueUB:255"`
	DelayCritical               []byte                    `bitstring:"sizeLB:0,sizeUB:150"`
	AveragingWindow             AveragingWindow           `bitstring:"sizeLB:0,sizeUB:150"`
	MaximumDataBurstVolume      MaximumDataBurstVolume    `bitstring:"sizeLB:0,sizeUB:150"`
	ExtendedPacketDelayBudget   ExtendedPacketDelayBudget `bitstring:"sizeLB:0,sizeUB:150"`
	CnPacketDelayBudgetDownlink ExtendedPacketDelayBudget `bitstring:"sizeLB:0,sizeUB:150"`
	CnPacketDelayBudgetUplink   ExtendedPacketDelayBudget `bitstring:"sizeLB:0,sizeUB:150"`
}
