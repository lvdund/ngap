package ie

type GbrQosFlowInformation struct {
MaximumFlowBitRateDownlink	BitRate	//`bitstring:"sizeLB:0,sizeUB:150"`
MaximumFlowBitRateUplink	BitRate	//`bitstring:"sizeLB:0,sizeUB:150"`
GuaranteedFlowBitRateDownlink	BitRate	//`bitstring:"sizeLB:0,sizeUB:150"`
GuaranteedFlowBitRateUplink	BitRate	//`bitstring:"sizeLB:0,sizeUB:150"`
NotificationControl	[]byte	//`bitstring:"sizeLB:0,sizeUB:150"`
MaximumPacketLossRateDownlink	PacketLossRate	//`bitstring:"sizeLB:0,sizeUB:150"`
MaximumPacketLossRateUplink	PacketLossRate	//`bitstring:"sizeLB:0,sizeUB:150"`
AlternativeQosParametersSetList	AlternativeQosParametersSetList	//`bitstring:"sizeLB:0,sizeUB:150"`
}
