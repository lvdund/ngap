package ies

import "github.com/lvdund/ngap/aper"

type TargeteNBID struct {
	GlobalENBID    *GlobalNgENBID `True,`
	SelectedEPSTAI *EPSTAI        `True,`
	// IEExtensions TargeteNBIDExtIEs `False,OPTIONAL`
}

func (ie *TargeteNBID) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.GlobalENBID != nil {
		if err = ie.GlobalENBID.Encode(w); err != nil {
			return
		}
	}
	if ie.SelectedEPSTAI != nil {
		if err = ie.SelectedEPSTAI.Encode(w); err != nil {
			return
		}
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
	ie.GlobalENBID = new(GlobalNgENBID)
	ie.SelectedEPSTAI = new(EPSTAI)
	if err = ie.GlobalENBID.Decode(r); err != nil {
		return
	}
	if err = ie.SelectedEPSTAI.Decode(r); err != nil {
		return
	}
	return
}
