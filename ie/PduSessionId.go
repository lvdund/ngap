package ie

import "ngap/aper"

type PduSessionId struct {
	PduSessionId aper.Integer //`Integer:"valueLB:0,valueUB:255"`
}
