package ie

import "ngap/aper"

type Periodicity struct {
	Periodicity aper.Integer `Integer:"valueLB:0,valueUB:640000"`
}
