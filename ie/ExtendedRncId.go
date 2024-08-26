package ie

import "ngap/aper"

type ExtendedRncId struct {
	ExtendedRncId aper.Integer `Integer:"valueLB:4096,valueUB:65535"`
}
