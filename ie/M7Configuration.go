package ie

import "ngap/aper"

type M7Configuration struct {
	M7CollectionPeriod aper.Integer //`Integer:"valueLB:1,valueUB:60"`
	M7LinksToLog       []byte       //`bitstring:"sizeLB:0,sizeUB:150"`
	M7ReportAmount     []byte       //`bitstring:"sizeLB:0,sizeUB:150"`
}
