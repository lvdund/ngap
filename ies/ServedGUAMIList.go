package ies

import "github.com/lvdund/ngap/aper"

type ServedGUAMIList struct {
	Value []*ServedGUAMIItem `False,1,256`
}

func (ie *ServedGUAMIList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*ServedGUAMIItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return
	}
	return
}
func (ie *ServedGUAMIList) Decode(r *aper.AperReader) (err error) {
	var newItems []*ServedGUAMIItem
	newItems, err = aper.ReadSequenceOfEx[*ServedGUAMIItem](func() *ServedGUAMIItem { return new(ServedGUAMIItem) }, r, &aper.Constraint{Lb: 1, Ub: 256}, false)
	ie.Value = []*ServedGUAMIItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
