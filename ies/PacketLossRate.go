package ies

import "github.com/lvdund/ngap/aper"

type PacketLossRate struct {
	Value aper.Integer `True,0,1000`
}

func (ie *PacketLossRate) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 0, Ub: 1000}, true)
	return
}
func (ie *PacketLossRate) Decode(r *aper.AperReader) (err error) {
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 1000}, true); err != nil {
		return
	}
	ie.Value = aper.Integer(v)
	return
}
