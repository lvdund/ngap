package ies

import "github.com/lvdund/ngap/aper"

type EmergencyAreaIDBroadcastEUTRA struct {
	Value []*EmergencyAreaIDBroadcastEUTRAItem `False,1,65535`
}

func (ie *EmergencyAreaIDBroadcastEUTRA) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*EmergencyAreaIDBroadcastEUTRAItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 65535}, false); err != nil {
		return
	}
	return
}
func (ie *EmergencyAreaIDBroadcastEUTRA) Decode(r *aper.AperReader) (err error) {
	var newItems []*EmergencyAreaIDBroadcastEUTRAItem
	newItems, err = aper.ReadSequenceOfEx[*EmergencyAreaIDBroadcastEUTRAItem](func() *EmergencyAreaIDBroadcastEUTRAItem { return new(EmergencyAreaIDBroadcastEUTRAItem) }, r, &aper.Constraint{Lb: 1, Ub: 65535}, false)
	ie.Value = []*EmergencyAreaIDBroadcastEUTRAItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
