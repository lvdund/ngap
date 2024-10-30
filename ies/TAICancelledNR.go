package ies

import "github.com/lvdund/ngap/aper"

type TAICancelledNR struct {
	Value []*TAICancelledNRItem `False,1,65535`
}

func (ie *TAICancelledNR) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*TAICancelledNRItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 65535}, false); err != nil {
		return
	}
	return
}
func (ie *TAICancelledNR) Decode(r *aper.AperReader) (err error) {
	var newItems []*TAICancelledNRItem
	newItems, err = aper.ReadSequenceOfEx[*TAICancelledNRItem](func() *TAICancelledNRItem { return new(TAICancelledNRItem) }, r, &aper.Constraint{Lb: 1, Ub: 65535}, false)
	ie.Value = []*TAICancelledNRItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
