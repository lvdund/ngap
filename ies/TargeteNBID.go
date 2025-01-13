package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type TargeteNBID struct {
	GlobalENBID    GlobalNgENBID `madatory`
	SelectedEPSTAI EPSTAI        `madatory`
	// IEExtensions *TargeteNBIDExtIEs `optional`
}

func (ie *TargeteNBID) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.GlobalENBID.Encode(w); err != nil {
		err = utils.WrapError("Encode GlobalENBID", err)
		return
	}
	if err = ie.SelectedEPSTAI.Encode(w); err != nil {
		err = utils.WrapError("Encode SelectedEPSTAI", err)
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
		err = utils.WrapError("Read GlobalENBID", err)
		return
	}
	if err = ie.SelectedEPSTAI.Decode(r); err != nil {
		err = utils.WrapError("Read SelectedEPSTAI", err)
		return
	}
	return
}
