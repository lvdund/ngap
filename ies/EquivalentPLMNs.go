package ies

import "github.com/lvdund/ngap/aper"

type EquivalentPLMNs struct {
	Value []*PLMNIdentity `False,1,15`
}

func (ie *EquivalentPLMNs) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*PLMNIdentity](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
		return
	}
	return
}
func (ie *EquivalentPLMNs) Decode(r *aper.AperReader) (err error) {
	var newItems []*PLMNIdentity
	newItems, err = aper.ReadSequenceOfEx[*PLMNIdentity](func() *PLMNIdentity { return new(PLMNIdentity) }, r, &aper.Constraint{Lb: 1, Ub: 15}, false)
	ie.Value = []*PLMNIdentity{}
	ie.Value = append(ie.Value, newItems...)
	return
}
