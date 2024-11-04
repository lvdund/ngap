package ies

import "github.com/lvdund/ngap/aper"

const (
	CauseNasNormalrelease         aper.Enumerated = 0
	CauseNasAuthenticationfailure aper.Enumerated = 1
	CauseNasDeregister            aper.Enumerated = 2
	CauseNasUnspecified           aper.Enumerated = 3
)

type CauseNas struct {
	Value aper.Enumerated `True,0,3`
}

func (ie *CauseNas) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, true)
	return
}
func (ie *CauseNas) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, true)
	ie.Value = aper.Enumerated(v)
	return
}
