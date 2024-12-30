package ies

import "github.com/lvdund/ngap/aper"

const (
	CauseTransportTransportresourceunavailable aper.Enumerated = 0
	CauseTransportUnspecified                  aper.Enumerated = 1
)

type CauseTransport struct {
	Value aper.Enumerated
}

func (ie *CauseTransport) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, true)
	return
}
func (ie *CauseTransport) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true)
	ie.Value = aper.Enumerated(v)
	return
}
