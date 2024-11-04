package ies

import "github.com/lvdund/ngap/aper"

const (
	OverloadActionRejectnonemergencymodt                                    aper.Enumerated = 0
	OverloadActionRejectrrccrsignalling                                     aper.Enumerated = 1
	OverloadActionPermitemergencysessionsandmobileterminatedservicesonly    aper.Enumerated = 2
	OverloadActionPermithighprioritysessionsandmobileterminatedservicesonly aper.Enumerated = 3
)

type OverloadAction struct {
	Value aper.Enumerated `True,0,3`
}

func (ie *OverloadAction) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, true)
	return
}
func (ie *OverloadAction) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, true)
	ie.Value = aper.Enumerated(v)
	return
}
