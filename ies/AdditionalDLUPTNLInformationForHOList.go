package ies

import "github.com/lvdund/ngap/aper"

type AdditionalDLUPTNLInformationForHOList struct {
	Value []*AdditionalDLUPTNLInformationForHOItem `False,1,3`
}

func (ie *AdditionalDLUPTNLInformationForHOList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*AdditionalDLUPTNLInformationForHOItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 3}, false); err != nil {
		return
	}
	return
}
func (ie *AdditionalDLUPTNLInformationForHOList) Decode(r *aper.AperReader) (err error) {
	var newItems []*AdditionalDLUPTNLInformationForHOItem
	newItems, err = aper.ReadSequenceOfEx[*AdditionalDLUPTNLInformationForHOItem](func() *AdditionalDLUPTNLInformationForHOItem { return new(AdditionalDLUPTNLInformationForHOItem) }, r, &aper.Constraint{Lb: 1, Ub: 3}, false)
	ie.Value = []*AdditionalDLUPTNLInformationForHOItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
