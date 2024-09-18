package ie

import "ngap/aper"

// Need to import "github.com/free5gc/aper" if it uses "aper"

const (
	CriticalityPresentReject aper.Enumerated = 0
	CriticalityPresentIgnore aper.Enumerated = 1
	CriticalityPresentNotify aper.Enumerated = 2
)

type Criticality struct {
	Value aper.Enumerated `aper:"valueLB:0,valueUB:2"`
}

func (ie *Criticality) Decode(r aper.AperReader) error {

	return nil
}

func (ie *Criticality) Encode(r aper.AperWriter) (err error) {
	if err = r.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return err
	}
	return nil
}
