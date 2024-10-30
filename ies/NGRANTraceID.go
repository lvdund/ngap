package ies

import "github.com/lvdund/ngap/aper"

type NGRANTraceID struct {
	Value aper.OctetString `False,8,8`
}

func (ie *NGRANTraceID) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteOctetString(ie.Value, &aper.Constraint{Lb: 8, Ub: 8}, false)
	return
}
func (ie *NGRANTraceID) Decode(r *aper.AperReader) (err error) {
	var v aper.OctetString
	if v, err = r.ReadOctetString(&aper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
		return
	}
	ie.Value = v
	return
}
