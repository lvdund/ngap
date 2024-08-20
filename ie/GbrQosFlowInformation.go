package ie

type GbrQosFlowInformation struct {
	MaximumFlowBitRateDownlink      *BitRate
	MaximumFlowBitRateUplink        *BitRate
	GuaranteedFlowBitRateDownlink   *BitRate
	GuaranteedFlowBitRateUplink     *BitRate
	NotificationControl             *[]byte
	MaximumPacketLossRateDownlink   *PacketLossRate
	MaximumPacketLossRateUplink     *PacketLossRate
	AlternativeQosParametersSetList *AlternativeQosParametersSetList
}
