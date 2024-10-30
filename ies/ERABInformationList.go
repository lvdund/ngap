package ies

import "github.com/lvdund/ngap/aper"

type ERABInformationList struct {
	Value []*ERABInformationItem `False,1,256`
}

func (ie *ERABInformationList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*ERABInformationItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return
	}
	return
}
func (ie *ERABInformationList) Decode(r *aper.AperReader) (err error) {
	var newItems []*ERABInformationItem
	newItems, err = aper.ReadSequenceOfEx[*ERABInformationItem](func() *ERABInformationItem { return new(ERABInformationItem) }, r, &aper.Constraint{Lb: 1, Ub: 256}, false)
	ie.Value = []*ERABInformationItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
