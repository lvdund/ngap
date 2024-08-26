package ie

import "ngap/aper"

type NetworkInstance struct {
	NetworkInstance aper.Integer `Integer:"valueLB:1,valueUB:256"`
}
