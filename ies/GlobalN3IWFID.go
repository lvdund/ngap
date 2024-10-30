package ies

import "github.com/lvdund/ngap/aper"

type GlobalN3IWFID struct {
	PLMNIdentity *PLMNIdentity `False,`
	N3IWFID      *N3IWFID      `False,`
	// IEExtensions GlobalN3IWFIDExtIEs `False,OPTIONAL`
}

func (ie *GlobalN3IWFID) Encode(w *aper.AperWriter) (err error) {
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
	if ie.N3IWFID != nil {
		if err = ie.N3IWFID.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *GlobalN3IWFID) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.PLMNIdentity = new(PLMNIdentity)
	ie.N3IWFID = new(N3IWFID)
	if err = ie.PLMNIdentity.Decode(r); err != nil {
		return
	}
	if err = ie.N3IWFID.Decode(r); err != nil {
		return
	}
	return
}
