package ie

import "ngap/aper"

type InterSystemCellActivationReply struct {
	ActivatedCellsList ActivatedCellsList //`bitstring:"sizeLB:0,sizeUB:150"`
	ActivationId       aper.Integer       //`Integer:"valueLB:0,valueUB:16384"`
}

type ActivatedCellsList struct {
	NgRanCgi NgRanCgi //`bitstring:"sizeLB:0,sizeUB:150"`
}
