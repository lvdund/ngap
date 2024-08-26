package ie

import "ngap/aper"

type Pc5QosParameters struct {
	Pc5QosFlowList           []Pc5QosFlowItem `bitstring:"sizeLB:0,sizeUB:150"`
	Pc5LinkAggregateBitRates BitRate          `bitstring:"sizeLB:0,sizeUB:150"`
}

type Pc5QosFlowItem struct {
	Pqi             aper.Integer    `Integer:"valueLB:0,valueUB:255"`
	Pc5FlowBitRates Pc5FlowBitRates `bitstring:"sizeLB:0,sizeUB:150"`
	Range           []byte          `bitstring:"sizeLB:0,sizeUB:150"`
}

type Pc5FlowBitRates struct {
	GuaranteedFlowBitRate BitRate `bitstring:"sizeLB:0,sizeUB:150"`
	MaximumFlowBitRate    BitRate `bitstring:"sizeLB:0,sizeUB:150"`
}
