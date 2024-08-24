package ie

import "ngap/aper"

type ExpectedUeActivityBehaviour struct {
	ExpectedActivityPeriod                 aper.Integer //`Integer:"valueLB:1,valueUB:30"`
	ExpectedIdlePeriod                     aper.Integer //`Integer:"valueLB:1,valueUB:30"`
	SourceOfUeActivityBehaviourInformation []byte       //`bitstring:"sizeLB:0,sizeUB:150"`
}
