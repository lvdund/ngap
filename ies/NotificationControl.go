package ies

import "github.com/lvdund/ngap/aper"

const (
	NotificationControlNotificationrequested aper.Enumerated = 0
)

type NotificationControl struct {
	Value aper.Enumerated `True,0,0`
}

func (ie *NotificationControl) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, true)
	return
}
func (ie *NotificationControl) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, true)
	ie.Value = aper.Enumerated(v)
	return
}
