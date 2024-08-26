package ie

import "ngap/aper"

type CountValueForPdcpSnLength18 struct {
	PdcpSnLength18       aper.Integer `Integer:"valueLB:0,valueUB:262143"`
	HfnForPdcpSnLength18 aper.Integer `Integer:"valueLB:0,valueUB:16383"`
}
