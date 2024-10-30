package ies

import "github.com/lvdund/ngap/aper"

type WarningAreaCoordinates struct {
	Value aper.OctetString `False,1,1024`
}

func (ie *WarningAreaCoordinates) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteOctetString(ie.Value, &aper.Constraint{Lb: 1, Ub: 1024}, false)
	return
}
func (ie *WarningAreaCoordinates) Decode(r *aper.AperReader) (err error) {
	var v aper.OctetString
	if v, err = r.ReadOctetString(&aper.Constraint{Lb: 1, Ub: 1024}, false); err != nil {
		return
	}
	ie.Value = v
	return
}
