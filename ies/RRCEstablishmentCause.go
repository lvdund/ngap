package ies

import "github.com/lvdund/ngap/aper"

const (
	RRCEstablishmentCauseEmergency          aper.Enumerated = 0
	RRCEstablishmentCauseHighpriorityaccess aper.Enumerated = 1
	RRCEstablishmentCauseMtaccess           aper.Enumerated = 2
	RRCEstablishmentCauseMosignalling       aper.Enumerated = 3
	RRCEstablishmentCauseModata             aper.Enumerated = 4
	RRCEstablishmentCauseMovoicecall        aper.Enumerated = 5
	RRCEstablishmentCauseMovideocall        aper.Enumerated = 6
	RRCEstablishmentCauseMosms              aper.Enumerated = 7
	RRCEstablishmentCauseMpspriorityaccess  aper.Enumerated = 8
	RRCEstablishmentCauseMcspriorityaccess  aper.Enumerated = 9
	RRCEstablishmentCauseNotavailable       aper.Enumerated = 10
)

type RRCEstablishmentCause struct {
	Value aper.Enumerated
}

func (ie *RRCEstablishmentCause) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 10}, true)
	return
}
func (ie *RRCEstablishmentCause) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 10}, true)
	ie.Value = aper.Enumerated(v)
	return
}
