package ie

import "ngap/aper"

type AlternativeQosParametersSetIndex struct {
	AlternativeQosParametersSetIndex aper.Integer `Integer:"valueLB:1,valueUB:8"`
}
