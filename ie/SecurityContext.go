package ie

import "ngap/aper"

type SecurityContext struct {
NextHopChainingCount	aper.Integer	//`Integer:"valueLB:0,valueUB:7"`
NextHopNh	SecurityKey	//`bitstring:"sizeLB:0,sizeUB:150"`
}
