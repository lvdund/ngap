package ies

import "github.com/lvdund/ngap/aper"

type NRPPaPDU struct {
	Value aper.OctetString `False,0,0`
}

func (ie *NRPPaPDU) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteOctetString(ie.Value, &aper.Constraint{Lb: 0, Ub: 0}, false)
	return
}
func (ie *NRPPaPDU) Decode(r *aper.AperReader) (err error) {
	var v aper.OctetString
	if v, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return
	}
	ie.Value = v
	return
}
