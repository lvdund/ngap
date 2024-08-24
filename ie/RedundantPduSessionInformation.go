package ie

import "ngap/aper"

type RedundantPduSessionInformation struct {
	Rsn              []byte       //`bitstring:"sizeLB:0,sizeUB:150"`
	PduSessionPairId aper.Integer //`Integer:"valueLB:0,valueUB:150"`
}
