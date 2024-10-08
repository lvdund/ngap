package ie

import "ngap/aper"

type Ie5GProsePc5QosParameters struct {
	Ie5GProsePc5QosFlowList           []Ie5GProsePc5QosFlowItem `bitstring:"sizeLB:0,sizeUB:150"`
	Ie5GProsePc5LinkAggregateBitRates BitRate                   `bitstring:"sizeLB:0,sizeUB:150"`
}

type Ie5GProsePc5QosFlowItem struct {
	Pqi                      aper.Integer             `Integer:"valueLB:0,valueUB:255"`
	Ie5GProsePc5FlowBitRates Ie5GProsePc5FlowBitRates `bitstring:"sizeLB:0,sizeUB:150"`
	Range                    []byte                   `bitstring:"sizeLB:0,sizeUB:150"`
}

type Ie5GProsePc5FlowBitRates struct {
	GuaranteedFlowBitRate BitRate `bitstring:"sizeLB:0,sizeUB:150"`
	MaximumFlowBitRate    BitRate `bitstring:"sizeLB:0,sizeUB:150"`
}
