package ies

import "github.com/lvdund/ngap/aper"

type TAIBroadcastEUTRA struct {
	Value []*TAIBroadcastEUTRAItem `False,1,65535`
}

func (ie *TAIBroadcastEUTRA) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*TAIBroadcastEUTRAItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 65535}, false); err != nil {
		return
	}
	return
}
func (ie *TAIBroadcastEUTRA) Decode(r *aper.AperReader) (err error) {
	var newItems []*TAIBroadcastEUTRAItem
	newItems, err = aper.ReadSequenceOfEx[*TAIBroadcastEUTRAItem](func() *TAIBroadcastEUTRAItem { return new(TAIBroadcastEUTRAItem) }, r, &aper.Constraint{Lb: 1, Ub: 65535}, false)
	ie.Value = []*TAIBroadcastEUTRAItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
