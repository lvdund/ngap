package ies

import "github.com/lvdund/ngap/aper"

const (
	CauseMiscPresentControlProcessingOverload             aper.Enumerated = 0
	CauseMiscPresentNotEnoughUserPlaneProcessingResources aper.Enumerated = 1
	CauseMiscPresentHardwareFailure                       aper.Enumerated = 2
	CauseMiscPresentOmIntervention                        aper.Enumerated = 3
	CauseMiscPresentUnknownPLMN                           aper.Enumerated = 4
	CauseMiscPresentUnspecified                           aper.Enumerated = 5
)


type CauseMisc struct {
	Value aper.Enumerated `True,0,6`
}

func (ie *CauseMisc) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 6}, true)
	return
}
func (ie *CauseMisc) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 6}, true)
	ie.Value = aper.Enumerated(v)
	return
}
