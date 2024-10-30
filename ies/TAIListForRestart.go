package ies

import "github.com/lvdund/ngap/aper"

type TAIListForRestart struct {
	Value []*TAI `False,1,2048`
}

func (ie *TAIListForRestart) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*TAI](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 2048}, false); err != nil {
		return
	}
	return
}
func (ie *TAIListForRestart) Decode(r *aper.AperReader) (err error) {
	var newItems []*TAI
	newItems, err = aper.ReadSequenceOfEx[*TAI](func() *TAI { return new(TAI) }, r, &aper.Constraint{Lb: 1, Ub: 2048}, false)
	ie.Value = []*TAI{}
	ie.Value = append(ie.Value, newItems...)
	return
}
