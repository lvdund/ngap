package ie

import "ngap/aper"

type NumberOfBroadcasts struct {
	NumberOfBroadcasts aper.Integer `Integer:"valueLB:0,valueUB:65535"`
}
