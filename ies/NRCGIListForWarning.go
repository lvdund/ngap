package ies

import "github.com/lvdund/ngap/aper"

type NRCGIListForWarning struct {
	Value []*NRCGI `False,1,65535`
}

func (ie *NRCGIListForWarning) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*NRCGI](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 65535}, false); err != nil {
		return
	}
	return
}
func (ie *NRCGIListForWarning) Decode(r *aper.AperReader) (err error) {
	var newItems []*NRCGI
	newItems, err = aper.ReadSequenceOfEx[*NRCGI](func() *NRCGI { return new(NRCGI) }, r, &aper.Constraint{Lb: 1, Ub: 65535}, false)
	ie.Value = []*NRCGI{}
	ie.Value = append(ie.Value, newItems...)
	return
}
