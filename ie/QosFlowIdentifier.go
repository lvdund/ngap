package ie

import "ngap/aper"

type QosFlowIdentifier struct {
	QosFlowIdentifier aper.Integer `Integer:"valueLB:0,valueUB:63"`
}
