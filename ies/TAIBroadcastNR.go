package ies

import "github.com/lvdund/ngap/aper"

type TAIBroadcastNR struct {
	Value []*TAIBroadcastNRItem `False,1,65535`
}

func (ie *TAIBroadcastNR) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*TAIBroadcastNRItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 65535}, false); err != nil {
		return
	}
	return
}
func (ie *TAIBroadcastNR) Decode(r *aper.AperReader) (err error) {
	var newItems []*TAIBroadcastNRItem
	newItems, err = aper.ReadSequenceOfEx[*TAIBroadcastNRItem](func() *TAIBroadcastNRItem { return new(TAIBroadcastNRItem) }, r, &aper.Constraint{Lb: 1, Ub: 65535}, false)
	ie.Value = []*TAIBroadcastNRItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
