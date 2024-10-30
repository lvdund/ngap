package ies

import "github.com/lvdund/ngap/aper"

type NGAPMessage struct {
	Value aper.OctetString
}

func (ie *NGAPMessage) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteOctetString(ie.Value, &aper.Constraint{Lb: 8, Ub: 8}, false)
	return
}
func (ie *NGAPMessage) Decode(r *aper.AperReader) (err error) {
	var v aper.OctetString
	if v, err = r.ReadOctetString(&aper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
		return
	}
	ie.Value = v
	return
}
