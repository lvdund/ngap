package ie

import "ngap/aper"

type TnlAddressWeightFactor struct {
	TnlAddressWeightFactor aper.Integer `Integer:"valueLB:0,valueUB:255"`
}
