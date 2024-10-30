package ies

import "github.com/lvdund/ngap/aper"

type EPSTAI struct {
	PLMNIdentity *PLMNIdentity `False,`
	EPSTAC       *EPSTAC       `False,`
	// IEExtensions EPSTAIExtIEs `False,OPTIONAL`
}

func (ie *EPSTAI) Encode(w *aper.AperWriter) (err error) {
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
	if ie.EPSTAC != nil {
		if err = ie.EPSTAC.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *EPSTAI) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.PLMNIdentity = new(PLMNIdentity)
	ie.EPSTAC = new(EPSTAC)
	if err = ie.PLMNIdentity.Decode(r); err != nil {
		return
	}
	if err = ie.EPSTAC.Decode(r); err != nil {
		return
	}
	return
}
