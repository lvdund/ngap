package ies

import "github.com/lvdund/ngap/aper"

type EUTRACGIListForWarning struct {
	Value []*EUTRACGI `False,1,65535`
}

func (ie *EUTRACGIListForWarning) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*EUTRACGI](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 65535}, false); err != nil {
		return
	}
	return
}
func (ie *EUTRACGIListForWarning) Decode(r *aper.AperReader) (err error) {
	var newItems []*EUTRACGI
	newItems, err = aper.ReadSequenceOfEx[*EUTRACGI](func() *EUTRACGI { return new(EUTRACGI) }, r, &aper.Constraint{Lb: 1, Ub: 65535}, false)
	ie.Value = []*EUTRACGI{}
	ie.Value = append(ie.Value, newItems...)
	return
}
