package ie

type AlternativeQosParametersSetList struct {
	AlternativeQosParametersSetItem *AlternativeQosParametersSetItem
}

type AlternativeQosParametersSetItem struct {
	AlternativeQosParametersSetIndex *AlternativeQosParametersSetIndex
	GuaranteedFlowBitRateDownlink    *BitRate
	GuaranteedFlowBitRateUplink      *BitRate
	PacketDelayBudget                *PacketDelayBudget
	PacketErrorRate                  *PacketErrorRate
}
