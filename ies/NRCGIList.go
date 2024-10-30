package ies

import "github.com/lvdund/ngap/aper"

type NRCGIList struct {
	Value []*NRCGI `False,1,16384`
}

func (ie *NRCGIList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*NRCGI](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 16384}, false); err != nil {
		return
	}
	return
}
func (ie *NRCGIList) Decode(r *aper.AperReader) (err error) {
	var newItems []*NRCGI
	newItems, err = aper.ReadSequenceOfEx[*NRCGI](func() *NRCGI { return new(NRCGI) }, r, &aper.Constraint{Lb: 1, Ub: 16384}, false)
	ie.Value = []*NRCGI{}
	ie.Value = append(ie.Value, newItems...)
	return
}
