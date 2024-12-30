package ies

import "github.com/lvdund/ngap/aper"

const (
	DataForwardingAcceptedDataforwardingaccepted aper.Enumerated = 0
)

type DataForwardingAccepted struct {
	Value aper.Enumerated
}

func (ie *DataForwardingAccepted) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, true)
	return
}
func (ie *DataForwardingAccepted) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, true)
	ie.Value = aper.Enumerated(v)
	return
}
