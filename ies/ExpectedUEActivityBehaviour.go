package ies

import "github.com/lvdund/ngap/aper"

type ExpectedUEActivityBehaviour struct {
	ExpectedActivityPeriod                 *int64
	ExpectedIdlePeriod                     *int64
	SourceOfUEActivityBehaviourInformation *SourceOfUEActivityBehaviourInformation
	// IEExtensions  *ExpectedUEActivityBehaviourExtIEs
}

func (ie *ExpectedUEActivityBehaviour) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.ExpectedActivityPeriod != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.ExpectedIdlePeriod != nil {
		aper.SetBit(optionals, 2)
	}
	if ie.SourceOfUEActivityBehaviourInformation != nil {
		aper.SetBit(optionals, 3)
	}
	w.WriteBits(optionals, 4)
	if ie.ExpectedActivityPeriod != nil {
		tmp_ExpectedActivityPeriod := NewINTEGER(*ie.ExpectedActivityPeriod, aper.Constraint{Lb: 1, Ub: 181}, true)
		if err = tmp_ExpectedActivityPeriod.Encode(w); err != nil {
			return
		}
	}
	if ie.ExpectedIdlePeriod != nil {
		tmp_ExpectedIdlePeriod := NewINTEGER(*ie.ExpectedIdlePeriod, aper.Constraint{Lb: 1, Ub: 181}, true)
		if err = tmp_ExpectedIdlePeriod.Encode(w); err != nil {
			return
		}
	}
	if ie.SourceOfUEActivityBehaviourInformation != nil {
		if err = ie.SourceOfUEActivityBehaviourInformation.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *ExpectedUEActivityBehaviour) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(4); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp_ExpectedActivityPeriod := INTEGER{
			c:   aper.Constraint{Lb: 1, Ub: 181},
			ext: false,
		}
		if err = tmp_ExpectedActivityPeriod.Decode(r); err != nil {
			return
		}
		*ie.ExpectedActivityPeriod = int64(tmp_ExpectedActivityPeriod.Value)
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_ExpectedIdlePeriod := INTEGER{
			c:   aper.Constraint{Lb: 1, Ub: 181},
			ext: false,
		}
		if err = tmp_ExpectedIdlePeriod.Decode(r); err != nil {
			return
		}
		*ie.ExpectedIdlePeriod = int64(tmp_ExpectedIdlePeriod.Value)
	}
	if aper.IsBitSet(optionals, 3) {
		if err = ie.SourceOfUEActivityBehaviourInformation.Decode(r); err != nil {
			return
		}
	}
	return
}
