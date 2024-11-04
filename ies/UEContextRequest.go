package ies

import "github.com/lvdund/ngap/aper"

const (
	UEContextRequestPresentRequested aper.Enumerated = 0
)

type UEContextRequest struct {
	Value aper.Enumerated `True,0,0`
}

func (ie *UEContextRequest) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, true)
	return
}
func (ie *UEContextRequest) Decode(r *aper.AperReader) (err error) {
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, true); err != nil {
		return
	}
	ie.Value = aper.Enumerated(v)
	return
}
