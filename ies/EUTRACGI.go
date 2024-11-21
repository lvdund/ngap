package ies

import "github.com/lvdund/ngap/aper"

type EUTRACGI struct {
	PLMNIdentity      *PLMNIdentity      `False,`
	EUTRACellIdentity *EUTRACellIdentity `False,`
	// IEExtensions EUTRACGIExtIEs `False,OPTIONAL`
}

func (ie *EUTRACGI) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.PLMNIdentity != nil {
		if err = ie.PLMNIdentity.Encode(w); err != nil {
			return
		}
	}
	if ie.EUTRACellIdentity != nil {
		if err = ie.EUTRACellIdentity.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *EUTRACGI) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.PLMNIdentity = new(PLMNIdentity)
	ie.EUTRACellIdentity = new(EUTRACellIdentity)
	if err = ie.PLMNIdentity.Decode(r); err != nil {
		return
	}
	if err = ie.EUTRACellIdentity.Decode(r); err != nil {
		return
	}
	return
}
