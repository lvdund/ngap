package ie

import "ngap/aper"

type RelativeAmfCapacity struct {
	RelativeAmfCapacity aper.Integer `Integer:"valueLB:0,valueUB:255"`
}
