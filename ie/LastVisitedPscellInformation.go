package ie

import "ngap/aper"

type LastVisitedPscellInformation struct {
PscellId	NgRanCgi	//`bitstring:"sizeLB:0,sizeUB:150"`
TimeStay	aper.Integer	//`Integer:"valueLB:0,valueUB:40950"`
}
