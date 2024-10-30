package ies

import "github.com/lvdund/ngap/aper"

type BroadcastPLMNList struct {
	Value []*BroadcastPLMNItem `False,1,12`
}

func (ie *BroadcastPLMNList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*BroadcastPLMNItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 12}, false); err != nil {
		return
	}
	return
}
func (ie *BroadcastPLMNList) Decode(r *aper.AperReader) (err error) {
	var newItems []*BroadcastPLMNItem
	newItems, err = aper.ReadSequenceOfEx[*BroadcastPLMNItem](func() *BroadcastPLMNItem { return new(BroadcastPLMNItem) }, r, &aper.Constraint{Lb: 1, Ub: 12}, false)
	ie.Value = []*BroadcastPLMNItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
