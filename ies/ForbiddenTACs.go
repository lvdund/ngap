package ies

import "github.com/lvdund/ngap/aper"

type ForbiddenTACs struct {
	Value []*TAC `False,1,4096`
}

func (ie *ForbiddenTACs) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*TAC](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 4096}, false); err != nil {
		return
	}
	return
}
func (ie *ForbiddenTACs) Decode(r *aper.AperReader) (err error) {
	var newItems []*TAC
	newItems, err = aper.ReadSequenceOfEx[*TAC](func() *TAC { return new(TAC) }, r, &aper.Constraint{Lb: 1, Ub: 4096}, false)
	ie.Value = []*TAC{}
	ie.Value = append(ie.Value, newItems...)
	return
}
