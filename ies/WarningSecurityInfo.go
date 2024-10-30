package ies

import "github.com/lvdund/ngap/aper"

type WarningSecurityInfo struct {
	Value aper.OctetString `False,50,50`
}

func (ie *WarningSecurityInfo) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteOctetString(ie.Value, &aper.Constraint{Lb: 50, Ub: 50}, false)
	return
}
func (ie *WarningSecurityInfo) Decode(r *aper.AperReader) (err error) {
	var v aper.OctetString
	if v, err = r.ReadOctetString(&aper.Constraint{Lb: 50, Ub: 50}, false); err != nil {
		return
	}
	ie.Value = v
	return
}
