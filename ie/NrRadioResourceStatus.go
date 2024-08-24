package ie

import "ngap/aper"

type NrRadioResourceStatus struct {
	DlGbrPrbUsageForMimo    aper.Integer //`Integer:"valueLB:0,valueUB:100"`
	UlGbrPrbUsageForMimo    aper.Integer //`Integer:"valueLB:0,valueUB:100"`
	DlNonGbrPrbUsageForMimo aper.Integer //`Integer:"valueLB:0,valueUB:100"`
	UlNonGbrPrbUsageForMimo aper.Integer //`Integer:"valueLB:0,valueUB:100"`
	DlTotalPrbUsageForMimo  aper.Integer //`Integer:"valueLB:0,valueUB:100"`
	UlTotalPrbUsageForMimo  aper.Integer //`Integer:"valueLB:0,valueUB:100"`
}
