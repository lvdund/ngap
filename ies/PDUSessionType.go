package ies

import "github.com/lvdund/ngap/aper"

const (
	PDUSessionTypeIpv4         aper.Enumerated = 0
	PDUSessionTypeIpv6         aper.Enumerated = 1
	PDUSessionTypeIpv4V6       aper.Enumerated = 2
	PDUSessionTypeEthernet     aper.Enumerated = 3
	PDUSessionTypeUnstructured aper.Enumerated = 4
)

type PDUSessionType struct {
	Value aper.Enumerated
}

func (ie *PDUSessionType) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 4}, true)
	return
}
func (ie *PDUSessionType) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 4}, true)
	ie.Value = aper.Enumerated(v)
	return
}
