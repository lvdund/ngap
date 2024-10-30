package ies

import "github.com/lvdund/ngap/aper"

type RejectedNSSAIinPLMN struct {
	Value aper.OctetString `False,32,32`
}

func (ie *RejectedNSSAIinPLMN) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteOctetString(ie.Value, &aper.Constraint{Lb: 32, Ub: 32}, false)
	return
}
func (ie *RejectedNSSAIinPLMN) Decode(r *aper.AperReader) (err error) {
	var v aper.OctetString
	if v, err = r.ReadOctetString(&aper.Constraint{Lb: 32, Ub: 32}, false); err != nil {
		return
	}
	ie.Value = v
	return
}
