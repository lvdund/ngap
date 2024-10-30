package ies

import "github.com/lvdund/ngap/aper"

type AMFUENGAPID struct {
	Value aper.Integer `False,0,1099511627775`
}

func (ie *AMFUENGAPID) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 0, Ub: 1099511627775}, false)
	return
}
func (ie *AMFUENGAPID) Decode(r *aper.AperReader) (err error) {
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 1099511627775}, false); err != nil {
		return
	}
	ie.Value = aper.Integer(v)
	return
}
