package ies

import "github.com/lvdund/ngap/aper"

const (
	CauseMiscControlprocessingoverload             aper.Enumerated = 0
	CauseMiscNotenoughuserplaneprocessingresources aper.Enumerated = 1
	CauseMiscHardwarefailure                       aper.Enumerated = 2
	CauseMiscOmintervention                        aper.Enumerated = 3
	CauseMiscUnknownplmn                           aper.Enumerated = 4
	CauseMiscUnspecified                           aper.Enumerated = 5
)

type CauseMisc struct {
	Value aper.Enumerated `True,0,5`
}

func (ie *CauseMisc) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 5}, true)
	return
}
func (ie *CauseMisc) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 5}, true)
	ie.Value = aper.Enumerated(v)
	return
}
