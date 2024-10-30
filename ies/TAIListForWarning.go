package ies

import "github.com/lvdund/ngap/aper"

type TAIListForWarning struct {
	Value []*TAI `False,1,65535`
}

func (ie *TAIListForWarning) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*TAI](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 65535}, false); err != nil {
		return
	}
	return
}
func (ie *TAIListForWarning) Decode(r *aper.AperReader) (err error) {
	var newItems []*TAI
	newItems, err = aper.ReadSequenceOfEx[*TAI](func() *TAI { return new(TAI) }, r, &aper.Constraint{Lb: 1, Ub: 65535}, false)
	ie.Value = []*TAI{}
	ie.Value = append(ie.Value, newItems...)
	return
}
