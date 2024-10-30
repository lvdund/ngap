package ies

import "github.com/lvdund/ngap/aper"

type NumberOfBroadcasts struct {
	Value aper.Integer `False,0,65535`
}

func (ie *NumberOfBroadcasts) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 0, Ub: 65535}, false)
	return
}
func (ie *NumberOfBroadcasts) Decode(r *aper.AperReader) (err error) {
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
		return
	}
	ie.Value = aper.Integer(v)
	return
}
