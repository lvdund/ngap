package ie

type AlternativeQosParametersSetList struct {
	AlternativeQosParametersSetItem AlternativeQosParametersSetItem `bitstring:"sizeLB:0,sizeUB:150"`
}

type AlternativeQosParametersSetItem struct {
	AlternativeQosParametersSetIndex AlternativeQosParametersSetIndex `bitstring:"sizeLB:0,sizeUB:150"`
	GuaranteedFlowBitRateDownlink    BitRate                          `bitstring:"sizeLB:0,sizeUB:150"`
	GuaranteedFlowBitRateUplink      BitRate                          `bitstring:"sizeLB:0,sizeUB:150"`
	PacketDelayBudget                PacketDelayBudget                `bitstring:"sizeLB:0,sizeUB:150"`
	PacketErrorRate                  PacketErrorRate                  `bitstring:"sizeLB:0,sizeUB:150"`
}
