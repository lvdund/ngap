package ies

import "github.com/lvdund/ngap/aper"

type CellIDBroadcastNR struct {
	Value []*CellIDBroadcastNRItem `False,1,65535`
}

func (ie *CellIDBroadcastNR) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*CellIDBroadcastNRItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 65535}, false); err != nil {
		return
	}
	return
}
func (ie *CellIDBroadcastNR) Decode(r *aper.AperReader) (err error) {
	var newItems []*CellIDBroadcastNRItem
	newItems, err = aper.ReadSequenceOfEx[*CellIDBroadcastNRItem](func() *CellIDBroadcastNRItem { return new(CellIDBroadcastNRItem) }, r, &aper.Constraint{Lb: 1, Ub: 65535}, false)
	ie.Value = []*CellIDBroadcastNRItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
