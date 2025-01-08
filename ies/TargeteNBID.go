package ies

import "github.com/lvdund/ngap/aper"

type TargeteNBID struct {
	GlobalENBID    GlobalNgENBID
	SelectedEPSTAI EPSTAI
	// IEExtensions *TargeteNBIDExtIEs `optional`
}

func (ie *TargeteNBID) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.GlobalENBID.Encode(w); err != nil {
		return
	}
	if err = ie.SelectedEPSTAI.Encode(w); err != nil {
		return
	}
	return
}
func (ie *TargeteNBID) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.GlobalENBID.Decode(r); err != nil {
		return
	}
	if err = ie.SelectedEPSTAI.Decode(r); err != nil {
		return
	}
	return
}
