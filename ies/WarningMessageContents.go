package ies

import "github.com/lvdund/ngap/aper"

type WarningMessageContents struct {
	Value aper.OctetString `False,1,9600`
}

func (ie *WarningMessageContents) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteOctetString(ie.Value, &aper.Constraint{Lb: 1, Ub: 9600}, false)
	return
}
func (ie *WarningMessageContents) Decode(r *aper.AperReader) (err error) {
	var v aper.OctetString
	if v, err = r.ReadOctetString(&aper.Constraint{Lb: 1, Ub: 9600}, false); err != nil {
		return
	}
	ie.Value = v
	return
}
