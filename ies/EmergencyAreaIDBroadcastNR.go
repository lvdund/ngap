package ies

import "github.com/lvdund/ngap/aper"

type EmergencyAreaIDBroadcastNR struct {
	Value []*EmergencyAreaIDBroadcastNRItem `False,1,65535`
}

func (ie *EmergencyAreaIDBroadcastNR) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*EmergencyAreaIDBroadcastNRItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 65535}, false); err != nil {
		return
	}
	return
}
func (ie *EmergencyAreaIDBroadcastNR) Decode(r *aper.AperReader) (err error) {
	var newItems []*EmergencyAreaIDBroadcastNRItem
	newItems, err = aper.ReadSequenceOfEx[*EmergencyAreaIDBroadcastNRItem](func() *EmergencyAreaIDBroadcastNRItem { return new(EmergencyAreaIDBroadcastNRItem) }, r, &aper.Constraint{Lb: 1, Ub: 65535}, false)
	ie.Value = []*EmergencyAreaIDBroadcastNRItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
