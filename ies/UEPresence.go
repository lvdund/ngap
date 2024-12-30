package ies

import "github.com/lvdund/ngap/aper"

const (
	UEPresenceIn      aper.Enumerated = 0
	UEPresenceOut     aper.Enumerated = 1
	UEPresenceUnknown aper.Enumerated = 2
)

type UEPresence struct {
	Value aper.Enumerated
}

func (ie *UEPresence) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, true)
	return
}
func (ie *UEPresence) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, true)
	ie.Value = aper.Enumerated(v)
	return
}
