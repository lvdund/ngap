package ies

import "github.com/lvdund/ngap/aper"

type UnavailableGUAMIList struct {
	Value []*UnavailableGUAMIItem `False,1,256`
}

func (ie *UnavailableGUAMIList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*UnavailableGUAMIItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return
	}
	return
}
func (ie *UnavailableGUAMIList) Decode(r *aper.AperReader) (err error) {
	var newItems []*UnavailableGUAMIItem
	newItems, err = aper.ReadSequenceOfEx[*UnavailableGUAMIItem](func() *UnavailableGUAMIItem { return new(UnavailableGUAMIItem) }, r, &aper.Constraint{Lb: 1, Ub: 256}, false)
	ie.Value = []*UnavailableGUAMIItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
