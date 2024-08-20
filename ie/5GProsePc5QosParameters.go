package ie

type Ie5GProsePc5QosParameters struct {
Ie5GProsePc5QosFlowList	*[]Ie5GProsePc5QosFlowItem
Ie5GProsePc5LinkAggregateBitRates	*BitRate
}

type Ie5GProsePc5QosFlowItem struct {
Pqi	uint8	//`bitstring:"sizeLB:0,sizeUB:255"`
Ie5GProsePc5FlowBitRates	*Ie5GProsePc5FlowBitRates
Range	*[]byte
}

type Ie5GProsePc5FlowBitRates struct {
GuaranteedFlowBitRate	*BitRate
MaximumFlowBitRate	*BitRate
}
