package ies

import "github.com/lvdund/ngap/aper"

const (
	EventTypeDirect                          aper.Enumerated = 0
	EventTypeChangeofservecell               aper.Enumerated = 1
	EventTypeUepresenceinareaofinterest      aper.Enumerated = 2
	EventTypeStopchangeofservecell           aper.Enumerated = 3
	EventTypeStopuepresenceinareaofinterest  aper.Enumerated = 4
	EventTypeCancellocationreportingfortheue aper.Enumerated = 5
)

type EventType struct {
	Value aper.Enumerated `True,0,5`
}

func (ie *EventType) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 5}, true)
	return
}
func (ie *EventType) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 5}, true)
	ie.Value = aper.Enumerated(v)
	return
}
