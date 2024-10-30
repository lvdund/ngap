package ies

import "github.com/lvdund/ngap/aper"

type CellIDBroadcastEUTRA struct {
	Value []*CellIDBroadcastEUTRAItem `False,1,65535`
}

func (ie *CellIDBroadcastEUTRA) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*CellIDBroadcastEUTRAItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 65535}, false); err != nil {
		return
	}
	return
}
func (ie *CellIDBroadcastEUTRA) Decode(r *aper.AperReader) (err error) {
	var newItems []*CellIDBroadcastEUTRAItem
	newItems, err = aper.ReadSequenceOfEx[*CellIDBroadcastEUTRAItem](func() *CellIDBroadcastEUTRAItem { return new(CellIDBroadcastEUTRAItem) }, r, &aper.Constraint{Lb: 1, Ub: 65535}, false)
	ie.Value = []*CellIDBroadcastEUTRAItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
