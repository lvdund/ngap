package ies

import "github.com/lvdund/ngap/aper"

type TAICancelledEUTRA struct {
	Value []*TAICancelledEUTRAItem `False,1,65535`
}

func (ie *TAICancelledEUTRA) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*TAICancelledEUTRAItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 65535}, false); err != nil {
		return
	}
	return
}
func (ie *TAICancelledEUTRA) Decode(r *aper.AperReader) (err error) {
	var newItems []*TAICancelledEUTRAItem
	newItems, err = aper.ReadSequenceOfEx[*TAICancelledEUTRAItem](func() *TAICancelledEUTRAItem { return new(TAICancelledEUTRAItem) }, r, &aper.Constraint{Lb: 1, Ub: 65535}, false)
	ie.Value = []*TAICancelledEUTRAItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
