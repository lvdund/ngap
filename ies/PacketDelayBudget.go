package ies

import "github.com/lvdund/ngap/aper"

type PacketDelayBudget struct {
	Value aper.Integer `True,0,1023`
}

func (ie *PacketDelayBudget) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 0, Ub: 1023}, true)
	return
}
func (ie *PacketDelayBudget) Decode(r *aper.AperReader) (err error) {
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 1023}, true); err != nil {
		return
	}
	ie.Value = aper.Integer(v)
	return
}
