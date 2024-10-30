package ies

import "github.com/lvdund/ngap/aper"

type SupportedTAList struct {
	Value []*SupportedTAItem `False,1,256`
}

func (ie *SupportedTAList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*SupportedTAItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return
	}
	return
}
func (ie *SupportedTAList) Decode(r *aper.AperReader) (err error) {
	var newItems []*SupportedTAItem
	newItems, err = aper.ReadSequenceOfEx[*SupportedTAItem](func() *SupportedTAItem { return new(SupportedTAItem) }, r, &aper.Constraint{Lb: 1, Ub: 256}, false)
	ie.Value = []*SupportedTAItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
