package ies

import "github.com/lvdund/ngap/aper"

type NRCGI struct {
	PLMNIdentity   *PLMNIdentity   `False,`
	NRCellIdentity *NRCellIdentity `False,`
	// IEExtensions NRCGIExtIEs `False,OPTIONAL`
}

func (ie *NRCGI) Encode(w *aper.AperWriter) (err error) {
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
	if ie.NRCellIdentity != nil {
		if err = ie.NRCellIdentity.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *NRCGI) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.PLMNIdentity = new(PLMNIdentity)
	ie.NRCellIdentity = new(NRCellIdentity)
	if err = ie.PLMNIdentity.Decode(r); err != nil {
		return
	}
	if err = ie.NRCellIdentity.Decode(r); err != nil {
		return
	}
	return
}
