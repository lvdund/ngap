package ie

import "ngap/aper"

type MbsAreaSessionId struct {
	MbsAreaSessionId aper.Integer `Integer:"valueLB:0,valueUB:65535"`
}
