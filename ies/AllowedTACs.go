package ies

import "github.com/lvdund/ngap/aper"

type AllowedTACs struct {
	Value []*TAC `False,1,16`
}

func (ie *AllowedTACs) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*TAC](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return
	}
	return
}
func (ie *AllowedTACs) Decode(r *aper.AperReader) (err error) {
	var newItems []*TAC
	newItems, err = aper.ReadSequenceOfEx[*TAC](func() *TAC { return new(TAC) }, r, &aper.Constraint{Lb: 1, Ub: 16}, false)
	ie.Value = []*TAC{}
	ie.Value = append(ie.Value, newItems...)
	return
}
