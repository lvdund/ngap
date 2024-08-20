package ie

type Pc5QosParameters struct {
	Pc5QosFlowList           *[]Pc5QosFlowItem
	Pc5LinkAggregateBitRates *BitRate
}

type Pc5QosFlowItem struct {
	Pqi             uint8 //`bitstring:"sizeLB:0,sizeUB:255"`
	Pc5FlowBitRates *Pc5FlowBitRates
	Range           *[]byte
}

type Pc5FlowBitRates struct {
	GuaranteedFlowBitRate *BitRate
	MaximumFlowBitRate    *BitRate
}
