package ies

import "github.com/lvdund/ngap/aper"

type ServiceAreaInformationItem struct {
	PLMNIdentity   *PLMNIdentity   `False,`
	AllowedTACs    *AllowedTACs    `False,OPTIONAL`
	NotAllowedTACs *NotAllowedTACs `False,OPTIONAL`
	// IEExtensions ServiceAreaInformationItemExtIEs `False,OPTIONAL`
}

func (ie *ServiceAreaInformationItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.AllowedTACs != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.NotAllowedTACs != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 3)
	if ie.PLMNIdentity != nil {
		if err = ie.PLMNIdentity.Encode(w); err != nil {
			return
		}
	}
	if ie.AllowedTACs != nil {
		if err = ie.AllowedTACs.Encode(w); err != nil {
			return
		}
	}
	if ie.NotAllowedTACs != nil {
		if err = ie.NotAllowedTACs.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *ServiceAreaInformationItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
		return
	}
	ie.PLMNIdentity = new(PLMNIdentity)
	ie.AllowedTACs = new(AllowedTACs)
	ie.NotAllowedTACs = new(NotAllowedTACs)
	if err = ie.PLMNIdentity.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		if err = ie.AllowedTACs.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 2) {
		if err = ie.NotAllowedTACs.Decode(r); err != nil {
			return
		}
	}
	return
}
