package ies

import "github.com/lvdund/ngap/aper"

type SST struct {
	Value aper.OctetString `False,1,1`
}

func (ie *SST) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteOctetString(ie.Value, &aper.Constraint{Lb: 1, Ub: 1}, false)
	return
}
func (ie *SST) Decode(r *aper.AperReader) (err error) {
	var v aper.OctetString
	if v, err = r.ReadOctetString(&aper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return
	}
	ie.Value = v
	return
}
