package ies

import "github.com/lvdund/ngap/aper"

type EUTRACGIList struct {
	Value []*EUTRACGI `False,1,256`
}

func (ie *EUTRACGIList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*EUTRACGI](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return
	}
	return
}
func (ie *EUTRACGIList) Decode(r *aper.AperReader) (err error) {
	var newItems []*EUTRACGI
	newItems, err = aper.ReadSequenceOfEx[*EUTRACGI](func() *EUTRACGI { return new(EUTRACGI) }, r, &aper.Constraint{Lb: 1, Ub: 256}, false)
	ie.Value = []*EUTRACGI{}
	ie.Value = append(ie.Value, newItems...)
	return
}
