package ie

import "gen/aper"

type Criticality struct {
	Value aper.Enumerated `aper:"valueLB:0,valueUB:2"`
}
type ProcedureCode struct {
	Value aper.Integer `aper:"valueLB:0,valueUB:255"`
}
type TriggeringMessage struct {
	Value aper.Enumerated `aper:"valueLB:0,valueUB:2"`
}
type ProtocolIEID struct {
	Value aper.Integer `aper:"valueLB:0,valueUB:65535"`
}
