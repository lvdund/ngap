package ies

import (
	"github.com/lvdund/ngap/aper"
)

type GlobalGNBID struct {
	PLMNIdentity *PLMNIdentity `False,`
	GNBID        *GNBID        `False,`
	// IEExtensions GlobalGNBIDExtIEs `False,OPTIONAL`
}

func (ie *GlobalGNBID) Encode(w *aper.AperWriter) (err error) {
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
	if ie.GNBID != nil {
		if err = ie.GNBID.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *GlobalGNBID) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.PLMNIdentity = new(PLMNIdentity)
	ie.GNBID = new(GNBID)
	if err = ie.PLMNIdentity.Decode(r); err != nil {
		return
	}
	if err = ie.GNBID.Decode(r); err != nil {
		return
	}
	return
}
