package ies

import "github.com/lvdund/ngap/aper"

type LocationReportingReferenceID struct {
	Value aper.Integer `True,1,64`
}

func (ie *LocationReportingReferenceID) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 1, Ub: 64}, true)
	return
}
func (ie *LocationReportingReferenceID) Decode(r *aper.AperReader) (err error) {
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 64}, true); err != nil {
		return
	}
	ie.Value = aper.Integer(v)
	return
}
