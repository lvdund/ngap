package ies

import "github.com/lvdund/ngap/aper"

type TAI struct {
	PLMNIdentity *PLMNIdentity `False,`
	TAC          *TAC          `False,`
	// IEExtensions TAIExtIEs `False,OPTIONAL`
}

func (ie *TAI) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.PLMNIdentity != nil {
		if err = ie.PLMNIdentity.Encode(w); err != nil {
			return
		}
	}
	if ie.TAC != nil {
		if err = ie.TAC.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *TAI) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.PLMNIdentity = new(PLMNIdentity)
	ie.TAC = new(TAC)
	if err = ie.PLMNIdentity.Decode(r); err != nil {
		return
	}
	if err = ie.TAC.Decode(r); err != nil {
		return
	}
	return
}
