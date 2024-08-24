package ie

import "ngap/aper"

type NrFrequencyInfo struct {
	NrArfcn             aper.Integer          //`Integer:"valueLB:0,valueUB:150"`
	NrFrequencyBandList []NrFrequencyBandItem //`bitstring:"sizeLB:0,sizeUB:150"`
}

type NrFrequencyBandItem struct {
	NrFrequencyBand aper.Integer //`Integer:"valueLB:0,valueUB:150"`
}
