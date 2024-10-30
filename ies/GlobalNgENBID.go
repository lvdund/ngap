package ies

import "github.com/lvdund/ngap/aper"

type GlobalNgENBID struct {
	PLMNIdentity *PLMNIdentity `False,`
	NgENBID      *NgENBID      `False,`
	// IEExtensions GlobalNgENBIDExtIEs `False,OPTIONAL`
}

func (ie *GlobalNgENBID) Encode(w *aper.AperWriter) (err error) {
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
	if ie.NgENBID != nil {
		if err = ie.NgENBID.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *GlobalNgENBID) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.PLMNIdentity = new(PLMNIdentity)
	ie.NgENBID = new(NgENBID)
	if err = ie.PLMNIdentity.Decode(r); err != nil {
		return
	}
	if err = ie.NgENBID.Decode(r); err != nil {
		return
	}
	return
}
