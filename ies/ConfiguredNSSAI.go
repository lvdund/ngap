package ies

import "github.com/lvdund/ngap/aper"

type ConfiguredNSSAI struct {
	Value aper.OctetString `False,128,128`
}

func (ie *ConfiguredNSSAI) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteOctetString(ie.Value, &aper.Constraint{Lb: 128, Ub: 128}, false)
	return
}
func (ie *ConfiguredNSSAI) Decode(r *aper.AperReader) (err error) {
	var v aper.OctetString
	if v, err = r.ReadOctetString(&aper.Constraint{Lb: 128, Ub: 128}, false); err != nil {
		return
	}
	ie.Value = v
	return
}
