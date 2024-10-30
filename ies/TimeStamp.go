package ies

import "github.com/lvdund/ngap/aper"

type TimeStamp struct {
	Value aper.OctetString `False,4,4`
}

func (ie *TimeStamp) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteOctetString(ie.Value, &aper.Constraint{Lb: 4, Ub: 4}, false)
	return
}
func (ie *TimeStamp) Decode(r *aper.AperReader) (err error) {
	var v aper.OctetString
	if v, err = r.ReadOctetString(&aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
		return
	}
	ie.Value = v
	return
}
