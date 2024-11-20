package ies

import "github.com/lvdund/ngap/aper"

type ExpectedUEActivityBehaviour struct {
	ExpectedActivityPeriod                 *ExpectedActivityPeriod                 `False,OPTIONAL`
	ExpectedIdlePeriod                     *ExpectedIdlePeriod                     `False,OPTIONAL`
	SourceOfUEActivityBehaviourInformation *SourceOfUEActivityBehaviourInformation `False,OPTIONAL`
	// IEExtensions ExpectedUEActivityBehaviourExtIEs `False,OPTIONAL`
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
		if err = ie.ExpectedActivityPeriod.Encode(w); err != nil {
			return
		}
	}
	if ie.ExpectedIdlePeriod != nil {
		if err = ie.ExpectedIdlePeriod.Encode(w); err != nil {
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
	ie.ExpectedActivityPeriod = new(ExpectedActivityPeriod)
	ie.ExpectedIdlePeriod = new(ExpectedIdlePeriod)
	ie.SourceOfUEActivityBehaviourInformation = new(SourceOfUEActivityBehaviourInformation)
	if aper.IsBitSet(optionals, 1) {
		if err = ie.ExpectedActivityPeriod.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 2) {
		if err = ie.ExpectedIdlePeriod.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 3) {
		if err = ie.SourceOfUEActivityBehaviourInformation.Decode(r); err != nil {
			return
		}
	}
	return
}
