package ie

import "ngap/aper"

type ExcessPacketDelayThresholdConfiguration struct {
	ExcessPacketDelayThresholdItem ExcessPacketDelayThresholdItem `bitstring:"sizeLB:0,sizeUB:150"`
}

type ExcessPacketDelayThresholdItem struct {
	Ie5Qi                           aper.Integer `Integer:"valueLB:0,valueUB:255"`
	ExcessPacketDelayThresholdValue []byte       `bitstring:"sizeLB:0,sizeUB:150"`
}
