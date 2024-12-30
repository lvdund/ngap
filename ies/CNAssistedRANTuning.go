package ies

import "github.com/lvdund/ngap/aper"

type CNAssistedRANTuning struct {
	ExpectedUEBehaviour *ExpectedUEBehaviour
	// IEExtensions  *CNAssistedRANTuningExtIEs
}

func (ie *CNAssistedRANTuning) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.ExpectedUEBehaviour != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if ie.ExpectedUEBehaviour != nil {
		if err = ie.ExpectedUEBehaviour.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *CNAssistedRANTuning) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		if err = ie.ExpectedUEBehaviour.Decode(r); err != nil {
			return
		}
	}
	return
}
