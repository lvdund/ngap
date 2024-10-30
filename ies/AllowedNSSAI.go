package ies

import "github.com/lvdund/ngap/aper"

type AllowedNSSAI struct {
	Value []*AllowedNSSAIItem `False,1,8`
}

func (ie *AllowedNSSAI) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*AllowedNSSAIItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return
	}
	return
}
func (ie *AllowedNSSAI) Decode(r *aper.AperReader) (err error) {
	var newItems []*AllowedNSSAIItem
	newItems, err = aper.ReadSequenceOfEx[*AllowedNSSAIItem](func() *AllowedNSSAIItem { return new(AllowedNSSAIItem) }, r, &aper.Constraint{Lb: 1, Ub: 8}, false)
	ie.Value = []*AllowedNSSAIItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
