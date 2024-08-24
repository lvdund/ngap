package ie

import "ngap/aper"

type EUtranRadioResourceStatus struct {
DlGbrPrbUsage	aper.Integer	//`Integer:"valueLB:0,valueUB:100"`
UlGbrPrbUsage	aper.Integer	//`Integer:"valueLB:0,valueUB:100"`
DlNonGbrPrbUsage	aper.Integer	//`Integer:"valueLB:0,valueUB:100"`
UlNonGbrPrbUsage	aper.Integer	//`Integer:"valueLB:0,valueUB:100"`
DlTotalPrbUsage	aper.Integer	//`Integer:"valueLB:0,valueUB:100"`
UlTotalPrbUsage	aper.Integer	//`Integer:"valueLB:0,valueUB:100"`
DlSchedulingPdcchCceUsage	aper.Integer	//`Integer:"valueLB:0,valueUB:100"`
UlSchedulingPdcchCceUsage	aper.Integer	//`Integer:"valueLB:0,valueUB:100"`
}
