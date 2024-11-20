package ies

import "github.com/lvdund/ngap/aper"

type ForbiddenAreaInformationItem struct {
	PLMNIdentity  *PLMNIdentity  `False,`
	ForbiddenTACs *ForbiddenTACs `False,`
	// IEExtensions ForbiddenAreaInformationItemExtIEs `False,OPTIONAL`
}

func (ie *ForbiddenAreaInformationItem) Encode(w *aper.AperWriter) (err error) {
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
	if ie.ForbiddenTACs != nil {
		if err = ie.ForbiddenTACs.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *ForbiddenAreaInformationItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.PLMNIdentity = new(PLMNIdentity)
	ie.ForbiddenTACs = new(ForbiddenTACs)
	if err = ie.PLMNIdentity.Decode(r); err != nil {
		return
	}
	if err = ie.ForbiddenTACs.Decode(r); err != nil {
		return
	}
	return
}
