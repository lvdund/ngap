package ie

import "ngap/aper"

type TrafficLoadReductionIndication struct {
	TrafficLoadReductionIndication aper.Integer `Integer:"valueLB:1,valueUB:99"`
}
