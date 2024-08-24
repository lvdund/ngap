package ie

import "ngap/aper"

type PacketLossRate struct {
	PacketLossRate aper.Integer //`Integer:"valueLB:0,valueUB:1000"`
}
